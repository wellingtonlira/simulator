package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type kafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(MsgChan chan *ckafka.Message) *kafkaConsumer {
	return &kafkaConsumer{
		MsgChan: MsgChan,
	}
}

func (k *kafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("Error consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaConsumer")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("KafkaConsumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
