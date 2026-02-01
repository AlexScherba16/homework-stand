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

	// Подключение к gRPC серверу
	conn, err := grpc.Dial(profileAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := profileV1.NewProfileServiceClient(conn)

	// Список пользователей
	userIDs := []int64{1, 2, 3}
	userIndex := 0

	// Канал для логирования
	ch := make(chan string, 100)
	defer close(ch)

	// Запускаем отдельный горутин для логирования
	go func() {
		for msg := range ch {
			slog.Info(msg)
		}
	}()

	// Тикер для генерации запросов каждые 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		idx := userIndex
		userIndex = (userIndex + 1) % len(userIDs)

		go func(userID int64) {
			start := time.Now()

			// Контекст с таймаутом на один запрос 3с
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			resp, err := client.GetProfile(ctx, &profileV1.GetProfileRequest{
				UserId: userID,
			})

			elapsed := time.Since(start)

			if err != nil {
				// Преобразуем ошибку в gRPC статус, если возможно
				if st, ok := status.FromError(err); ok {
					ch <- fmt.Sprintf("[ERROR] user=%d grpc_status=%v err=%v took=%v",
						userID, st.Code(), err, elapsed)
				} else {
					ch <- fmt.Sprintf("[ERROR] user=%d err=%v took=%v", userID, err, elapsed)
				}
				return
			}

			// Успешный результат
			ch <- fmt.Sprintf("[RESULT] user=%d took=%v result=%v", userID, elapsed, resp.String())
		}(userIDs[idx])
	}
}
