package messagebus

import (
	"task-service/internal/infrastructure/messagebus/producer"
	"task-service/internal/pkg/closer"
)

type Registry struct {
	Producers producer.Producers
}

func NewRegistry() *Registry {
	registry := &Registry{
		Producers: producer.NewProducers(),
	}

	closer.Add(registry.Producers.TaskEvents.Close)
	return registry
}
