package producer

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Retry.Backoff = 5 * time.Second

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{Producer: producer}, nil
}

func (p *KafkaProducer) SendMessage(topic string, key []byte, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	_, _, err := p.Producer.SendMessage(msg)
	if err != nil {
		log.Printf("Error sending message to topic %s: %v", topic, err)
		return err
	}

	log.Printf("Message sent to topic %s", topic)
	return nil
}

func (p *KafkaProducer) Close() error {
	return p.Producer.Close()
}
