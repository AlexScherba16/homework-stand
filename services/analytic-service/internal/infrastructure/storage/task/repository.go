package task

import (
	"context"
	"errors"

	"analytic-service/internal/domain/entity"
	"analytic-service/internal/pkg/transaction/wrapper"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	wrapper wrapper.Database
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{wrapper: wrapper.NewDatabase(pool)}
}

func (s *Storage) InsertTask(ctx context.Context, task *entity.Task) error {
	const sql = `
		insert into task_analytic(
			task_id,
			user_id,
			category_id,
			comment,
			status,
			price_amount,
			execution_time,
			created_at,
			updated_at
		) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		on conflict (task_id) do update set 
			status         = excluded.status,
			comment        = excluded.comment,
			price_amount   = excluded.price_amount,
			execution_time = excluded.execution_time,
			updated_at     = now();`

	_, err := s.wrapper.Pool(ctx).Exec(ctx, sql,
		task.TaskID,
		task.UserID,
		task.CategoryID,
		task.Comment,
		task.Status,
		task.Price,
		task.ExecutionTime,
		task.CreatedAt,
		task.UpdatedAt,
	)
	return err
}

func (s *Storage) CountTasks(ctx context.Context, userID int64) (uint64, error) {
	const sql = `select count(*) as task_count
				 from task_analytic
				 where user_id = $1;`

	var count int
	err := pgxscan.Get(ctx, s.wrapper.Pool(ctx), &count, sql, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return uint64(count), nil
}
