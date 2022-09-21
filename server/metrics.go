package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "server_request_count",
		Help: "How many request hit server",
	}, []string{"path", "status", "debug"})

	requestDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "server_request_ms",
		Help: "How long it takes to process request",
	}, []string{"path", "status", "debug"})
	requestDurationBucket = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "server_request_bucket_ms",
		Help:    "How long it takes to process request",
		Buckets: []float64{10, 20, 50, 100, 200, 500, 1000},
	}, []string{"path", "status", "debug"})
)

func NewRequestMetric(path, status, debug string) {
	requestCount.WithLabelValues(path, status, debug).Inc()
}

func RequestDurationMetric(path, status, debug string, start time.Time) {
	var duration = float64(time.Since(start).Milliseconds())

	requestDuration.WithLabelValues(path, status, debug).Set(duration)
	requestDurationBucket.WithLabelValues(path, status, debug).Observe(duration)
}
