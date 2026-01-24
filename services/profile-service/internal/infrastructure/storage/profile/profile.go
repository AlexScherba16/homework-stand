package profile

import (
	"context"
	"errors"

	"profile-service/internal/domain/entity"
	"profile-service/internal/infrastructure/storage/profile/dao"
	"profile-service/internal/pkg/perror"
	"profile-service/internal/pkg/transaction/wrapper"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	wrapper wrapper.Database
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{wrapper: wrapper.NewDatabase(pool)}
}

func (s *Storage) CreateProfile(ctx context.Context, p *entity.Profile) error {
	const sql = `
		insert into profile (user_id, name, email, is_task_allowed)
		values ($1, $2, $3, $4);`

	_, err := s.wrapper.Pool(ctx).Exec(ctx, sql,
		p.UserID,
		p.Name,
		p.Email,
		p.IsTaskAllowed,
	)
	if err != nil {
		// конфликт по user_id или email
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return perror.NewConflictErr("профиль уже существует")
		}
		return err
	}

	return nil
}

func (s *Storage) GetProfile(ctx context.Context, userID int64) (*entity.Profile, error) {
	var (
		sql = `select user_id,name,email,is_task_allowed,created_at from profile
			   where user_id = $1
			   limit 1;`

		profile dao.Profile
	)

	err := pgxscan.Get(ctx, s.wrapper.Pool(ctx), &profile, sql, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, perror.NewNotFoundErr("профиль не найден", err)
		}
		return nil, err
	}

	return profile.ConvertTo(), nil
}

func (s *Storage) GetProfiles(ctx context.Context, userIDs []int64) ([]*entity.Profile, error) {
	if len(userIDs) == 0 {
		return []*entity.Profile{}, nil
	}

	var (
		sql = `select user_id, name, email, is_task_allowed, created_at
			   from profile
			   where user_id = any($1);`

		profiles []dao.Profile
	)

	err := pgxscan.Select(ctx, s.wrapper.Pool(ctx), &profiles, sql, userIDs)
	if err != nil {
		return nil, err
	}

	result := make([]*entity.Profile, 0, len(profiles))
	for _, p := range profiles {
		result = append(result, p.ConvertTo())
	}

	return result, nil
}
