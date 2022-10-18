package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/wellingtonlira/simulator/infra/kafka"
)

func Producer(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := NewRoute

}
