package main

import (
	"context"

	"profile-service/internal"
)

func main() {
	ctx := context.Background()
	internal.New(ctx).Run(ctx)
}
