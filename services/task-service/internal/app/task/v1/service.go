package v1

import service "task-service/internal/application/service/task"

type Implementation struct {
	services *service.Registry
}

func NewTaskService(services *service.Registry) *Implementation {
	return &Implementation{
		services: services,
	}
}
