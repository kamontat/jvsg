package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "request_total",
		Help: "How many job execute",
	})
	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "raw_request_ms",
		Help:    "How long raw request take",
		Buckets: []float64{50, 70, 90, 120, 150, 200, 300, 500},
	})
	requestWithJsonDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "json_request_ms",
		Help:    "How long json request take",
		Buckets: []float64{50, 70, 90, 120, 150, 200, 300, 500},
	})
)

func NewRequestMetric() {
	requestTotal.Inc()
}

func RequestDurationMetric(start time.Time) {
	duration := float64(time.Since(start).Milliseconds())
	requestDuration.Observe(duration)
}

func RequestWithJsonDurationMetric(start time.Time) {
	duration := float64(time.Since(start).Milliseconds())
	requestWithJsonDuration.Observe(duration)
}
