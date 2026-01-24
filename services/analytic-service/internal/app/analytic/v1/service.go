package v1

import "analytic-service/internal/applicaton/service"

type Implementation struct {
	services *service.Registry
}

func NewAnalyticService(services *service.Registry) *Implementation {
	return &Implementation{
		services: services,
	}
}
