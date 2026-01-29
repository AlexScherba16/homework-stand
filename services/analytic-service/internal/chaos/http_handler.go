package chaos

import (
	"net/http"

	"analytic-service/internal/chaos/mode"
)

func ModeSetHandler(store *mode.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := r.URL.Query().Get("mode")
		switch mode.Mode(m) {
		case mode.OK, mode.Slow, mode.Error, mode.Flaky, mode.RareError:
			store.Set(mode.Mode(m))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("mode set to " + m))
		default:
			http.Error(w, "unknown mode", http.StatusBadRequest)
		}
	}
}
