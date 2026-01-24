package v1

import (
	taskV1 "api-gateway/internal/pkg/pb/api-gateway/task/v1"

	externalV1 "api-gateway/internal/pkg/pb/external/task-service/task/v1"
)

type Implementation struct {
	taskV1.UnimplementedTaskServiceServer
	external externalV1.TaskServiceClient
}

func NewTaskService(external externalV1.TaskServiceClient) *Implementation {
	return &Implementation{external: external}
}
