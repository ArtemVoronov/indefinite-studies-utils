package watcher

import (
	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/log"
	kafkaService "github.com/ArtemVoronov/indefinite-studies-utils/pkg/services/kafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OnGetKafkaMessage func(msg *kafka.Message)
type OnGetKafkaError func(err error)

type WatcherService struct {
	kafkaConsumer     *kafkaService.KafkaConsumerService
	quit              chan struct{}
	kafkaMessagesChan chan *kafka.Message
	kafkaErrorsChan   chan error
}

func CreateWatcherService(kafkaHostname string, kafkaGroupId string, topic string, pollPeriodInMillis int, onGetKafkaMessage OnGetKafkaMessage, onGetKafkaError OnGetKafkaError) *WatcherService {
	kafkaConsumer, err := kafkaService.CreateKafkaConsumerService(kafkaHostname, kafkaGroupId)
	if err != nil {
		log.Fatalf("unable to create kafka consumer: %s", err)
		return nil
	}

	quit := make(chan struct{})

	kafkaMessagesChan, kafkaErrorsChan, err := kafkaConsumer.PollTopics(quit, topic, pollPeriodInMillis)
	if err != nil {
		defer close(quit)
		log.Fatalf("unable to pull topics: %s", err)
		return nil
	}

	go func() {
		for e := range kafkaMessagesChan {
			onGetKafkaMessage(e)
		}
	}()
	go func() {
		for e := range kafkaErrorsChan {
			onGetKafkaError(e)
		}
	}()

	return &WatcherService{
		kafkaConsumer:     kafkaConsumer,
		quit:              quit,
		kafkaMessagesChan: kafkaMessagesChan,
		kafkaErrorsChan:   kafkaErrorsChan,
	}
}

func (s *WatcherService) Shutdown() error {
	defer close(s.quit)
	defer close(s.kafkaMessagesChan)
	defer close(s.kafkaErrorsChan)
	return s.kafkaConsumer.Shutdown()
}
