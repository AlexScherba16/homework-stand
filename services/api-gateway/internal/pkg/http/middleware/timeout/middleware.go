package timeout

import (
	"context"
	"net/http"

	"api-gateway/config"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		timeout := config.Instance().ExternalCalls.Timeout

		if timeout > 0 {
			ctx, cancel := context.WithTimeout(request.Context(), timeout)
			defer cancel()

			next.ServeHTTP(writer, request.WithContext(ctx))
			return
		}

		next.ServeHTTP(writer, request)
	})
}
