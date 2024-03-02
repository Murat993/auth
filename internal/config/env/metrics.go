package env

import (
	"context"
	"github.com/Murat993/auth/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "my_space"
	appName   = "my_app"
)

type metricsConfig struct {
	requestCounter        prometheus.Counter
	responseCounter       *prometheus.CounterVec
	histogramResponseTime *prometheus.HistogramVec
}

func NewMetrics(_ context.Context) config.MetricsConfig {
	return &metricsConfig{
		requestCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_requests_total",
				Help:      "Количество запросов к серверу",
			},
		),
		responseCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_responses_total",
				Help:      "Количество ответов от сервера",
			},
			[]string{"status", "method"},
		),
		histogramResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      appName + "_histogram_response_time_seconds",
				Help:      "Время ответа от сервера",
				Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
			},
			[]string{"status"},
		),
	}
}

func (m *metricsConfig) IncRequestCounter() {
	m.requestCounter.Inc()
}

func (m *metricsConfig) IncResponseCounter(status string, method string) {
	m.responseCounter.WithLabelValues(status, method).Inc()
}

func (m *metricsConfig) HistogramResponseTimeObserve(status string, time float64) {
	m.histogramResponseTime.WithLabelValues(status).Observe(time)
}
