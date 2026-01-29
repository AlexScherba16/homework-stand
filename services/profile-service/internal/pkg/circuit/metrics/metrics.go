package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "circuit"

var (
	once sync.Once

	example = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "example",
		Help:      "example",
	}, []string{"label_1", "label_2"})
)

func init() {
	once.Do(func() {
		prometheus.MustRegister(example)
	})
}

func IncExample(a, b string) {
	example.WithLabelValues(a, b).Inc()
}
