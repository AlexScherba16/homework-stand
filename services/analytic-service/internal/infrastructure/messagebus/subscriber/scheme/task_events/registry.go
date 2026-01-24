package task_events

import (
	"analytic-service/internal/applicaton/service"
	"analytic-service/internal/infrastructure/messagebus/subscriber/scheme/task_events/task_created"
)

type Registry struct {
	TaskCreated *task_created.MessageHandler
}

func NewRegistry(services *service.Registry) *Registry {
	return &Registry{
		TaskCreated: task_created.NewMessageHandler(services.AcceptTask),
	}
}
