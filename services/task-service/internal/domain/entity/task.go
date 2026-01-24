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

type TaskResult map[string]any // for protobuf.Struct / JSONB

type Task struct {
	ID         uuid.UUID
	UserID     int64
	CategoryID string

	Comment       string
	ExecutionTime time.Duration

	Price decimal.Decimal

	Status TaskStatus
	Result TaskResult // optional

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTask(
	userID int64,
	categoryID string,
	comment string,
	executionTime time.Duration,
) *Task {
	id, _ := uuid.NewV7()
	now := time.Now()
	return &Task{
		ID:            id,
		UserID:        userID,
		CategoryID:    categoryID,
		Comment:       comment,
		ExecutionTime: executionTime,

		Status: TaskStatusCreated,

		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *Task) SetPrice(price decimal.Decimal) {
	t.Price = price
}
