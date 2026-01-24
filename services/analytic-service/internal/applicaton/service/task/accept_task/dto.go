package accept_task

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type CreateTaskRequest struct {
	taskID        uuid.UUID
	userID        int64
	categoryID    string
	status        string
	comment       string
	executionTime time.Duration
	createdAt     time.Time
	price         decimal.Decimal
}

func NewCreateTaskRequest(
	taskID uuid.UUID,
	userID int64,
	categoryID string,
	status string,
	comment string,
	executionTime time.Duration,
	createdAt time.Time,
	price decimal.Decimal,
) CreateTaskRequest {
	return CreateTaskRequest{
		taskID:        taskID,
		userID:        userID,
		categoryID:    categoryID,
		status:        status,
		comment:       comment,
		executionTime: executionTime,
		createdAt:     createdAt,
		price:         price,
	}
}
