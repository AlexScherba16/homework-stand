package chaos

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	"analytic-service/internal/chaos/mode"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var requests int32 // количество одновременных запросов

func ModeInterceptor(store *mode.Store) grpc.UnaryServerInterceptor {
	rand.Seed(time.Now().UnixNano())

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		atomic.AddInt32(&requests, 1)

		md := store.Get()
		switch md {
		case mode.Error:
			return nil, status.Error(codes.Unavailable, "analytic in error mode")
		case mode.RareError:
			// 5% вероятности ошибки
			if rand.Intn(100) < 5 {
				return nil, status.Error(codes.Unavailable, "analytic flaky error")
			}
			return handler(ctx, req)
		case mode.Slow:
			delay := 200*time.Millisecond + time.Duration(atomic.LoadInt32(&requests))*100*time.Millisecond

			select {
			case <-time.After(delay):
				return handler(ctx, req)
			case <-ctx.Done():
				return nil, status.Error(codes.DeadlineExceeded, "request deadline exceeded")
			}
		case mode.Flaky:
			// 40% вероятности ошибки
			if rand.Intn(100) < 40 {
				return nil, status.Error(codes.Unavailable, "analytic flaky error")
			}
			return handler(ctx, req)
		default: // ModeOK
			atomic.StoreInt32(&requests, 0)
			return handler(ctx, req)
		}
	}
}
