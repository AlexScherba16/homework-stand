package service

import (
	"analytic-service/internal/applicaton/service/task/accept_task"
	"analytic-service/internal/applicaton/service/task/get_task_count"
	"analytic-service/internal/infrastructure/storage"
)

type Registry struct {
	AcceptTask   *accept_task.Service
	GetTaskCount *get_task_count.Service
}

func NewRegistry(storage *storage.Registry) *Registry {
	return &Registry{
		AcceptTask:   accept_task.NewService(storage.TaskStorage),
		GetTaskCount: get_task_count.NewService(storage.TaskStorage),
	}
}
