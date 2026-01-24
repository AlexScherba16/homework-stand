package gateway

import (
	"profile-service/config"
	"profile-service/internal/infrastructure/gateway/analytic"

	"google.golang.org/grpc"
)

type Registry struct {
	Analytic *analytic.Client
}

func NewRegistry(conn map[string]grpc.ClientConnInterface) *Registry {
	return &Registry{
		Analytic: analytic.NewClient(analytic.NewExternalClient(conn[config.AnalyticService])),
	}
}
