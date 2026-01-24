package producer

import (
	"task-service/internal/pkg/msgbus/producer"
)

type Producers struct {
	TaskEvents *producer.MessageProducer
}

func NewProducers() Producers {
	return Producers{
		TaskEvents: producer.NewMessageProducer(),
	}
}
