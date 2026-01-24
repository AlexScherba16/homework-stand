package payload

import (
	"time"

	"profile-service/internal/domain/entity"
)

//go:generate easyjson -all payload.go

//easyjson:json
type Profile struct {
	UserID        int64     `json:"user_id,omitempty"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	IsTaskAllowed bool      `json:"is_task_allowed"`
	CreatedAt     time.Time `json:"created_at"`
	CachedAt      time.Time `json:"cached_at"`
}

func (p *Profile) Expired(ttl time.Duration) bool {
	return time.Since(p.CachedAt) >= ttl
}

func ConvertFrom(profile *entity.Profile) Profile {
	return Profile{
		UserID:        profile.UserID,
		Email:         profile.Email,
		Name:          profile.Name,
		IsTaskAllowed: profile.IsTaskAllowed,
		CreatedAt:     profile.CreatedAt,
		CachedAt:      time.Now(), // для небольшой сигнатуры оставлю тут
	}
}

func (p *Profile) ConvertTo() *entity.Profile {
	return &entity.Profile{
		UserID:        p.UserID,
		Name:          p.Name,
		Email:         p.Email,
		IsTaskAllowed: p.IsTaskAllowed,
		CreatedAt:     p.CreatedAt,
	}
}
