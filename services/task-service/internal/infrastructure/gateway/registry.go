package gateway

import (
	"task-service/config"
	"task-service/internal/infrastructure/gateway/profile"

	"google.golang.org/grpc"
)

type Registry struct {
	Profile *profile.Client
}

func NewRegistry(conn map[string]grpc.ClientConnInterface) *Registry {
	return &Registry{
		Profile: profile.NewClient(profile.NewExternalClient(conn[config.ProfileService])),
	}
}
