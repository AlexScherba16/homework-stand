package task

import (
	"context"
	"errors"

	"task-service/internal/domain/entity"
	"task-service/internal/infrastructure/storage/task/dao"
	"task-service/internal/pkg/terror"
	"task-service/internal/pkg/transaction/wrapper"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/lo"
)

type Storage struct {
	wrapper wrapper.Database
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{wrapper: wrapper.NewDatabase(pool)}
}

func (s *Storage) CreateTask(ctx context.Context, task *entity.Task) error {
	sql := `insert into task(
			id,
			user_id,
			category_id,
			comment,
			execution_time,
			price_amount,
			status,
			created_at,
			updated_at
		) values ($1,$2,$3,$4,$5,$6,$7,$8,$9);`

	_, err := s.wrapper.Pool(ctx).Exec(ctx, sql,
		task.ID,              // uuid
		task.UserID,          // bigint
		task.CategoryID,      // text
		task.Comment,         // text / null
		task.ExecutionTime,   // interval
		task.Price.IntPart(), // bigint
		task.Status,          // varchar
		task.CreatedAt,       // timestamp
		task.UpdatedAt,       // timestamp
	)
	return err
}

func (s *Storage) GetTaskByID(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	var (
		sql = `select
			id,
			user_id,
			category_id,
			comment,
			execution_time,
			price_amount,
			status,
			result,
			created_at,
			updated_at
		from task
		where id = $1
		limit 1;`

		taskDAO dao.Task
	)

	err := pgxscan.Get(ctx, s.wrapper.Pool(ctx), &taskDAO, sql, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, terror.NewNotFoundErr("задаче не найдена", err)
		}
		return nil, err
	}

	return taskDAO.ConvertTo(), nil
}

func (s *Storage) GetTasksByUserID(ctx context.Context, userID int64) ([]*entity.Task, error) {
	var (
		sql = `select
			id,
			user_id,
			category_id,
			comment,
			execution_time,
			price_amount,
			status,
			result,
			created_at,
			updated_at
		from task
		where user_id = $1
		order by created_at desc;`

		tasks = make([]*dao.Task, 0)
	)

	err := pgxscan.Select(ctx, s.wrapper.Pool(ctx), &tasks, sql, userID)
	if err != nil {
		return nil, err
	}

	return lo.Map(tasks, func(task *dao.Task, _ int) *entity.Task {
		return task.ConvertTo()
	}), nil
}
