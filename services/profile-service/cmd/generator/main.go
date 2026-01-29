package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	profileV1 "profile-service/internal/pkg/pb/profile-service/profile/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	profileAddr := "localhost:7002" // адрес profile сервиса
	conn, err := grpc.Dial(profileAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := profileV1.NewProfileServiceClient(conn)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	userIDs := []int64{1, 2, 3} // список пользователей

	userIndex := 0

	for range ticker.C {
		start := time.Now()

		userID := userIDs[userIndex]
		userIndex = (userIndex + 1) % len(userIDs) // циклично проходим по списку

		// создаём контекст с timeout на одну операцию 3 секунды
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		resp, err := client.GetProfile(ctx, &profileV1.GetProfileRequest{
			UserId: userID,
		})
		cancel()

		elapsed := time.Since(start)

		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				slog.Info(fmt.Sprintf("[ERROR] gRPC status=%v  err=%v (took %v)\n",
					st.Code(), err, elapsed))
			} else {
				slog.Info(fmt.Sprintf("[ERROR] %v (took %v)\n", err, elapsed))
			}
			continue
		}

		// результат и время
		slog.Info(fmt.Sprintf("[RESULT] took=%v result=%v\n", elapsed, resp.String()))
	}
}
