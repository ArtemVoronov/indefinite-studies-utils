package kafka

import (
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

func CreateKafkaProducerService(hostname string) (*KafkaProducerService, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": hostname})
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

func (s *KafkaProducerService) Shutdown() error {
	s.producer.Close()
	return nil
}

func (s *KafkaConsumerService) Shutdown() error {
	return s.consumer.Close()
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

func (s *KafkaConsumerService) SubscribeTopics(quit <-chan struct{}, topics []string, pollPeriod time.Duration) (chan *kafka.Message, chan error, error) {
	out := make(chan *kafka.Message)
	outErr := make(chan error)

	err := s.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		defer close(out)
		defer close(outErr)
		return nil, nil, err
	}

	go func() {
		defer close(out)
		defer close(outErr)
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

	return out, outErr, nil
}

func (s *KafkaConsumerService) Unsubscribe() error {
	return s.consumer.Unsubscribe()
}

func (s *KafkaConsumerService) PollTopics(quit <-chan struct{}, topic string, pollPeriodInMillis int) (chan *kafka.Message, chan error, error) {
	out := make(chan *kafka.Message)
	outErr := make(chan error)

	err := s.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		defer close(out)
		defer close(outErr)
		return nil, nil, err
	}

	go func() {
		defer close(out)
		defer close(outErr)
		for {
			select {
			case <-quit:
				log.Debug("kafka consumer quit")
				return

			default:
				ev := s.consumer.Poll(pollPeriodInMillis)
				if ev == nil {
					continue
				}

				switch e := ev.(type) {
				case *kafka.Message:
					out <- e
				case kafka.Error:
					outErr <- fmt.Errorf("consumer error: %s", e)
				default:
					// TODO: check if need other event types
					// fmt.Printf("Ignored %v\n", e)
				}
			}
		}
	}()

	return out, outErr, nil
}
