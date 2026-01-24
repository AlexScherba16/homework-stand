package task_created

import (
	"context"
	"fmt"
	"log/slog"

	"analytic-service/internal/applicaton/service/task/accept_task"
	"analytic-service/internal/infrastructure/messagebus/subscriber/scheme/task_events/task_created/event"

	"github.com/IBM/sarama"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type TaskCreator interface {
	Create(ctx context.Context, request accept_task.CreateTaskRequest) error
}

type MessageHandler struct {
	creator TaskCreator
}

func NewMessageHandler(creator TaskCreator) *MessageHandler {
	return &MessageHandler{creator: creator}
}

func (h *MessageHandler) Handle(ctx context.Context, _ sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) error {
	// десереализуем сообщение
	deserialized, err := event.Deserialize(message)
	if err != nil {
		slog.Error(fmt.Sprintf("Ошибка десереализации сообщения: %s", err.Error()))
		return err
	}

	amount, err := decimal.NewFromString(deserialized.Price)
	if err != nil {
		return err
	}

	return h.creator.Create(ctx, accept_task.NewCreateTaskRequest(
		uuid.FromStringOrNil(deserialized.TaskID),
		deserialized.UserID,
		deserialized.CategoryID,
		deserialized.Status,
		deserialized.Comment,
		deserialized.ExecutionTime,
		deserialized.CreatedAt,
		amount,
	))
}
