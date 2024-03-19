package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducerService struct {
	producer *kafka.Producer
}

type KafkaConsumerService struct {
	consumer *kafka.Consumer
}

type KafkaAdminService struct {
	QueryTimeout time.Duration
	admin        *kafka.AdminClient
}

func CreateKafkaProducerService(hostname string) (*KafkaProducerService, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": hostname,
	})
	if err != nil {
		return nil, err
	}
	return &KafkaProducerService{producer: p}, nil
}

func CreateKafkaConsumerService(hostname string, groupId string) (*KafkaConsumerService, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": hostname,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	return &KafkaConsumerService{consumer: c}, nil
}

func CreateKafkaAdminService(hostname string, queryTimeout time.Duration) (*KafkaAdminService, error) {
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": hostname,
	})
	if err != nil {
		return nil, err
	}
	return &KafkaAdminService{admin: a, QueryTimeout: queryTimeout}, nil
}

func (s *KafkaProducerService) Shutdown() error {
	s.producer.Close()
	return nil
}

func (s *KafkaConsumerService) Shutdown() error {
	return s.consumer.Close()
}

func (s *KafkaAdminService) Shutdown() error {
	s.admin.Close()
	return nil
}

func (s *KafkaProducerService) CreateMessage(topic string, message string) error {
	return s.CreateMessageWithinPartition(topic, message, kafka.PartitionAny)
}

func (s *KafkaProducerService) CreateMessageWithinPartition(topic string, message string, partition int32) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err := s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Error("Delivery failed", m.TopicPartition.Error.Error())
		return m.TopicPartition.Error
	}

	log.Debug(fmt.Sprintf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset))

	return nil
}

func (s *KafkaConsumerService) StartReadingMessages(quit <-chan struct{}, out chan *kafka.Message, outErr chan error, topics []string, pollPeriod time.Duration) {
	go func() {
		for {
			select {
			case <-quit:
				log.Debug("kafka consumer quit")
				return
			default:
				msg, err := s.consumer.ReadMessage(pollPeriod)
				if err == nil {
					out <- msg
				} else if err.(kafka.Error).Code() != kafka.ErrTimedOut {
					outErr <- fmt.Errorf("consumer error: %s", err)
				}
			}
		}
	}()
}

func (s *KafkaConsumerService) SubscribeTopics(topics []string) error {
	return s.consumer.SubscribeTopics(topics, nil)
}

func (s *KafkaConsumerService) Unsubscribe() error {
	return s.consumer.Unsubscribe()
}

func (s *KafkaAdminService) CreateTopics(topics []string, numPartitions int) error {
	specs := make([]kafka.TopicSpecification, 0, len(topics))

	for _, topic := range topics {
		spec := kafka.TopicSpecification{
			Topic:         topic,
			NumPartitions: numPartitions,
		}
		specs = append(specs, spec)
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.QueryTimeout)
	defer cancel()

	_, err := s.admin.CreateTopics(ctx, specs)
	if err != nil {
		return err
	}
	return nil
}
