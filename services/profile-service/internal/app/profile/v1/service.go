package v1

import (
	"profile-service/internal/application/service"
	profileV1 "profile-service/internal/pkg/pb/profile-service/profile/v1"
)

type Implementation struct {
	profileV1.UnimplementedProfileServiceServer
	services *service.Registry
}

func NewProfileService(services *service.Registry) *Implementation {
	return &Implementation{services: services}
}
