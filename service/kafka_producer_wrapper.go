package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducerWrapper struct {
	KafkaConfig KafkaConfiguration
	Writer      *kafka.Writer
}

func NewKafkaProducerWrapper(topic string, partition int, bootstrapServers []string) *KafkaProducerWrapper {
	producer := &KafkaProducerWrapper{
		KafkaConfig: KafkaConfiguration{
			Topic: topic, Partition: partition, BootstrapServers: bootstrapServers},
	}
	producer.Writer = &kafka.Writer{
		Addr:                   kafka.TCP(bootstrapServers...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	return producer
}

func (producer *KafkaProducerWrapper) Work(message string) {
	messages := []kafka.Message{
		{
			Key:   []byte("KafkaProducerWrapper"),
			Value: []byte(message),
		},
	}
	var err error
	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = producer.Writer.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Fatalf("unexpected error %v", err)
		}
	}
}
