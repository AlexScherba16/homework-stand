package v1

import (
	profileV1 "api-gateway/internal/pkg/pb/api-gateway/profile/v1"
	externalV1 "api-gateway/internal/pkg/pb/external/profile-service/profile/v1"
)

type Implementation struct {
	profileV1.UnimplementedProfileServiceServer
	external externalV1.ProfileServiceClient
}

func NewProfileService(external externalV1.ProfileServiceClient) *Implementation {
	return &Implementation{external: external}
}
