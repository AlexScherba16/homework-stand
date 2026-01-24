package event

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
)

type TaskCreated struct {
	TaskID        string        `json:"task_id"`
	UserID        int64         `json:"user_id"`
	CategoryID    string        `json:"category_id"`
	Comment       string        `json:"comment"`
	Status        string        `json:"status"`
	Price         string        `json:"price"`
	ExecutionTime time.Duration `json:"execution_time"`
	CreatedAt     time.Time     `json:"created_at"`
}

// Deserialize переводит сообщение из одного представления в другое и валидирует
func Deserialize(message *sarama.ConsumerMessage) (TaskCreated, error) {
	var taskCreated TaskCreated
	err := json.Unmarshal(message.Value, &taskCreated)
	if err != nil {
		return TaskCreated{}, err
	}

	return taskCreated, nil
}
