package consumer

import (
	"log"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	Consumer sarama.Consumer
}

func NewKafkaConsumer(brokers []string) (*KafkaConsumer, error) {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{Consumer: consumer}, nil
}

func (c *KafkaConsumer) ConsumeMessages(topic string, partition int32, offset int64) {
	partitionConsumer, err := c.Consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("Consumed message from topic %s: %s", topic, string(msg.Value))
	}
}

func (c *KafkaConsumer) Close() error {
	return c.Consumer.Close()
}
