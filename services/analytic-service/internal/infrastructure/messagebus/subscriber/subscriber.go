package subscriber

import (
	"analytic-service/config"
	"analytic-service/internal/pkg/msgbus/subscriber"
)

type Subscribers struct {
	TaskEvents *subscriber.MessageSubscriber
}

func NewSubscribers() Subscribers {
	return Subscribers{
		TaskEvents: subscriber.NewMessageSubscriber(config.TaskEventsTopic),
	}
}
