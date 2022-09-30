package kafka

import (
	"fmt"
	"time"

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
	s.consumer.Close()
	return nil
}

func (s KafkaProducerService) CreateMessage(topic string, message string) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err := s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		// TODO: clean
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		return m.TopicPartition.Error
	}

	// TODO: clean
	fmt.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)

	return nil
}

func (s KafkaConsumerService) SubscribeTopics(quit <-chan struct{}, topic string, regexp string, pollPeriod time.Duration) (chan *kafka.Message, chan error) {
	out := make(chan *kafka.Message)
	outErr := make(chan error)

	s.consumer.SubscribeTopics([]string{topic, regexp}, nil)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Printf("kafka consumer quit\n")
				return

			default:
				msg, err := s.consumer.ReadMessage(pollPeriod)
				if err == nil {
					out <- msg
				} else {
					outErr <- fmt.Errorf("consumer error: %s", err)
				}

			}
		}
	}()

	return out, outErr
}

func (s *KafkaConsumerService) Unsubscribe() {
	s.consumer.Unsubscribe()
}

func (s *KafkaConsumerService) PollTopics(quit <-chan struct{}, topic string, pollPeriodInMillis int) (chan *kafka.Message, chan error) {
	out := make(chan *kafka.Message)
	outErr := make(chan error)

	s.consumer.SubscribeTopics([]string{topic}, nil)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Printf("kafka consumer quit\n")
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

	return out, outErr
}
