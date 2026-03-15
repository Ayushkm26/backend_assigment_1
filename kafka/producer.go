package kafka

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

var Producer sarama.SyncProducer

func InitProducer() {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(
		[]string{"localhost:9092"},
		config,
	)

	if err != nil {
		log.Fatal("Kafka producer error:", err)
	}

	Producer = producer
}
func PublishOrderCreated(event interface{}) error {

	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "order-created",
		Value: sarama.StringEncoder(jsonData),
	}

	_, _, err = Producer.SendMessage(msg)

	return err
}
