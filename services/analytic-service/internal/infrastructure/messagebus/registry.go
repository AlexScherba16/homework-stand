package messagebus

import (
	"context"

	"analytic-service/internal/applicaton/service"
	"analytic-service/internal/infrastructure/messagebus/subscriber"
	"analytic-service/internal/infrastructure/messagebus/subscriber/scheme/task_events"
	"analytic-service/internal/pkg/closer"
)

type handlers struct {
	TaskEvents *task_events.Registry
}

type Registry struct {
	handlers    handlers
	subscribers subscriber.Subscribers
}

func NewRegistry(services *service.Registry) *Registry {
	registry := &Registry{
		subscribers: subscriber.NewSubscribers(),
		handlers: handlers{
			TaskEvents: task_events.NewRegistry(services),
		},
	}

	closer.Add(registry.subscribers.TaskEvents.Close)
	return registry
}

func (r *Registry) Run(ctx context.Context) {
	// Можно добавить полноценный multiplexer для событий этого топика, но пока сосредоточимся на одном
	go r.subscribers.TaskEvents.Subscribe(ctx, r.handlers.TaskEvents.TaskCreated.Handle)
}
