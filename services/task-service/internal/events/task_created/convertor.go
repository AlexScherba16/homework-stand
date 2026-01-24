package task_created

import (
	"context"
	"encoding/json"
	"time"

	"task-service/config"
	"task-service/internal/domain/entity"

	"task-service/internal/pkg/event"

	"task-service/internal/pkg/pipe"
)

type TaskCreatedEvent struct {
	TaskID        string        `json:"task_id"`
	UserID        int64         `json:"user_id"`
	CategoryID    string        `json:"category_id"`
	Status        string        `json:"status"`
	Comment       string        `json:"comment"`
	ExecutionTime time.Duration `json:"execution_time"`
	CreatedAt     time.Time     `json:"created_at"`
	Price         string        `json:"price"`
}

func New(task *entity.Task) pipe.Func[event.Events] {
	return func(ctx context.Context, events event.Events) (event.Events, error) {
		// формируем событие
		body, err := json.Marshal(&TaskCreatedEvent{
			TaskID:        task.ID.String(),
			UserID:        task.UserID,
			CategoryID:    task.CategoryID,
			Status:        string(task.Status),
			Comment:       task.Comment,
			ExecutionTime: task.ExecutionTime,
			CreatedAt:     task.CreatedAt,
			Price:         task.Price.String(),
		})
		if err != nil {
			return nil, err
		}

		// загорловки в сообщении любые
		headers := map[string]string{
			"x-app-name":   "task-service",
			"x-event-type": "task-created",
		}

		// это выносится в общие функции
		headersRaw, err := json.Marshal(headers)
		if err != nil {
			return nil, err
		}

		return append(events, event.Event{
			Key:     event.Raw(task.ID.String()),
			Body:    body,
			Headers: headersRaw,
			Schema:  config.TaskEventsTopic,
		}), nil
	}
}
