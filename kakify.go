package kafkify

import (
	"log"

	"github.com/DecodeWorms/kakify/config"
	"github.com/DecodeWorms/kakify/consumer"
	"github.com/DecodeWorms/kakify/producer"

	"github.com/IBM/sarama"
)

func Initialize() {
	conf := config.LoadConfig()

	prod, err := producer.NewKafkaProducer(conf.Brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer prod.Close()

	cons, err := consumer.NewKafkaConsumer(conf.Brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer cons.Close()

	// Example usage
	err = prod.SendMessage("example-topic", []byte("key"), []byte("value"))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	cons.ConsumeMessages("example-topic", 0, sarama.OffsetNewest)
}
