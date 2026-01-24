package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type TaskStatus string

const (
	TaskStatusCreated   TaskStatus = "created"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

type Task struct {
	TaskID        uuid.UUID
	UserID        int64
	CategoryID    string
	Comment       string
	Status        string
	Price         decimal.Decimal
	ExecutionTime time.Duration
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewTask(
	taskID uuid.UUID,
	userID int64,
	categoryID string,
	status string,
	comment string,
	executionTime time.Duration,
	createdAt time.Time,
	price decimal.Decimal,
) *Task {
	return &Task{
		TaskID:        taskID,
		UserID:        userID,
		CategoryID:    categoryID,
		Status:        status,
		Comment:       comment,
		ExecutionTime: executionTime,
		CreatedAt:     createdAt,
		Price:         price,
		UpdatedAt:     time.Now(),
	}
}
