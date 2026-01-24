package dao

import (
	"profile-service/internal/domain/entity"
	"profile-service/internal/pkg/xo"
)

type Profile struct {
	xo.Profile
}

func (p Profile) ConvertTo() *entity.Profile {
	return &entity.Profile{
		UserID:        p.UserID,
		Name:          p.Name,
		Email:         p.Email,
		IsTaskAllowed: p.IsTaskAllowed.Bool,
		CreatedAt:     p.CreatedAt,
	}
}
