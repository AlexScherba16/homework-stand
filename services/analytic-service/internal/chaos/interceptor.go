package chaos

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"sync/atomic"
	"time"

	"analytic-service/internal/chaos/mode"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var slowSince atomic.Int64 // unix nano; 0 => не в Slow

func ModeInterceptor(store *mode.Store) grpc.UnaryServerInterceptor {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		slog.Info(fmt.Sprintf("method '%s' was invoked", info.FullMethod))

		md := store.Get()

		switch md {

		// ---------------- SLOW ----------------

		case mode.Slow:
			now := time.Now().UnixNano()

			// Если только что вошли в Slow — запоминаем время входа
			if slowSince.Load() == 0 {
				slowSince.Store(now)
			}

			elapsed := time.Duration(now - slowSince.Load())

			// Каждые 500ms деградация усиливается
			steps := elapsed / (500 * time.Millisecond)

			delay := 200*time.Millisecond +
				steps*100*time.Millisecond
			select {
			case <-time.After(delay):
				return handler(ctx, req)
			case <-ctx.Done():
				return nil, status.Error(codes.DeadlineExceeded, "request deadline exceeded")
			}

		// ---------------- ERROR ----------------

		case mode.Error:
			resetSlow()
			return nil, status.Error(codes.Unavailable, "analytic in error mode")

		// ---------------- RARE ERROR ----------------

		case mode.RareError:
			resetSlow()
			if r.Intn(100) < 5 {
				return nil, status.Error(codes.Unavailable, "analytic rare error")
			}
			return handler(ctx, req)

		// ---------------- FLAKY ----------------

		case mode.Flaky:
			resetSlow()
			if r.Intn(100) < 40 {
				return nil, status.Error(codes.Unavailable, "analytic flaky error")
			}
			return handler(ctx, req)

		// ---------------- OK ----------------

		default: // ModeOK
			resetSlow()
			return handler(ctx, req)
		}
	}
}

func resetSlow() {
	slowSince.Store(0)
}
