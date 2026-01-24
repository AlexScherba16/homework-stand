package analytic

import (
	"context"

	analyticV1 "profile-service/internal/pkg/pb/external/analytic-service/analytic/v1"

	"google.golang.org/grpc"
)

type ExternalClient analyticV1.AnalyticServiceClient

func NewExternalClient(conn grpc.ClientConnInterface) ExternalClient {
	return analyticV1.NewAnalyticServiceClient(conn)
}

type Client struct {
	external ExternalClient
}

func NewClient(external ExternalClient) *Client {
	return &Client{
		external: external,
	}
}

func (c *Client) GetUserTaskCount(ctx context.Context, userID int64) (uint64, error) {
	response, err := c.external.GetUserTaskCount(ctx, &analyticV1.GetUserTaskCountRequest{
		UserId: userID,
	})
	if err != nil {
		return 0, err
	}

	return response.GetTaskCount(), nil
}
