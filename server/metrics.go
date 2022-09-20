package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestCounterMetrics = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "server_request_total",
		Help: "How many request hit server",
	}, []string{"path", "status", "debug"})

	RequestDurationMetrics = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "server_request_duration_ms",
		Help:    "How long it takes to process request",
		Buckets: []float64{1, 3, 5, 10, 20, 50, 100, 200, 500, 1000},
	}, []string{"path", "status", "debug"})
)
