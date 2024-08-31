package config

import (
	"os"
)

type KafkaConfig struct {
	Brokers []string
}

func LoadConfig() KafkaConfig {
	brokers := os.Getenv("KAFKA_BROKERS")
	return KafkaConfig{
		Brokers: []string{brokers},
	}
}
