package dao

import (
	"encoding/json"
	"time"

	"task-service/internal/domain/entity"
	"task-service/internal/pkg/xo"

	"github.com/shopspring/decimal"
)

type Task struct {
	xo.Task
}

func (t Task) ConvertTo() *entity.Task {
	var result map[string]any

	if len(t.Result) > 0 {
		_ = json.Unmarshal(t.Result, &result)
	}

	var comment string
	if t.Comment.Valid {
		comment = t.Comment.String
	}

	var execTime time.Duration
	if t.ExecutionTime != nil {
		execTime = *t.ExecutionTime
	}

	return &entity.Task{
		ID:            t.ID,
		UserID:        t.UserID,
		CategoryID:    t.CategoryID,
		Comment:       comment,
		ExecutionTime: execTime,
		Price:         decimal.NewFromInt(t.PriceAmount),
		Status:        entity.TaskStatus(t.Status),
		Result:        result,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}
