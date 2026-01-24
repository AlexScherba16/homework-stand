package profile

import (
	"context"

	profileV1 "task-service/internal/pkg/pb/external/profile-service/profile/v1"

	"google.golang.org/grpc"
)

type ExternalClient profileV1.ProfileServiceClient

func NewExternalClient(conn grpc.ClientConnInterface) ExternalClient {
	return profileV1.NewProfileServiceClient(conn)
}

type Client struct {
	external ExternalClient
}

func NewClient(external ExternalClient) *Client {
	return &Client{
		external: external,
	}
}

func (c *Client) CheckPermission(ctx context.Context, userID int64) (bool, error) {
	response, err := c.external.CheckPermission(ctx, &profileV1.CheckPermissionRequest{
		UserId: userID,
	})
	if err != nil {
		return false, err
	}

	return response.GetAllowed(), nil
}
