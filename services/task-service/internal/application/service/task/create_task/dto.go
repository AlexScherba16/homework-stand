package create_task

import (
	"time"
)

type CreateTaskDTO struct {
	UserID        int64
	CategoryID    string
	Comment       string
	ExecutionTime time.Duration
}
