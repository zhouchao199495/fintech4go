package service

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumerWrapper struct {
	Topic            string
	Partition        int
	BootstrapServers []string
	KafkaConfig      KafkaConfiguration

	Reader   *kafka.Reader
	Consumer Consumer
}

func NewKafkaConsumerWrapper(topic string, partition int, bootstrapServers []string) *KafkaConsumerWrapper {
	consumer := &KafkaConsumerWrapper{
		KafkaConfig: KafkaConfiguration{
			Topic: topic, Partition: partition, BootstrapServers: bootstrapServers},
	}
	consumer.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   bootstrapServers,
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB

	})
	ff := consumer.Reader.Offset()
	fmt.Println("ff:", ff)
	return consumer
}

func NewConsumerGroupWrapper(topic string, groupId string, partition int, bootstrapServers []string) *KafkaConsumerWrapper {
	consumer := &KafkaConsumerWrapper{Topic: topic, Partition: partition, BootstrapServers: bootstrapServers}
	consumer.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  bootstrapServers,
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB

	})

	return consumer
}

func (consumer *KafkaConsumerWrapper) WorkFromBeginning() {
	consumer.Reader.SetOffset(0)
	consumer.Work()
}

func (consumer *KafkaConsumerWrapper) WorkFromSpecialOffset(offset int64) {
	consumer.Reader.SetOffset(offset)
	consumer.Work()
}

func (consumer *KafkaConsumerWrapper) Work() {
	for {
		m, err := consumer.Reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s \n", m.Topic, m.Partition, m.Offset, string(m.Key))
		consumer.Consumer.Consume(string(m.Value))
		consumer.Reader.CommitMessages(context.Background(), m)
	}
}
