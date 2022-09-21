package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "request_count",
		Help: "How many job execute",
	})
	requestDuration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "raw_request_ms",
		Help: "How long raw request take",
	})
	requestDurationBucket = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "raw_request_bucket_ms",
		Help:    "How long raw request take",
		Buckets: []float64{50, 70, 90, 120, 150, 200, 300, 500},
	})
	requestWithJsonDuration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "raw_request_ms",
		Help: "How long raw request take",
	})
	requestWithJsonDurationBucket = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "json_request_bucket_ms",
		Help:    "How long json request take",
		Buckets: []float64{50, 70, 90, 120, 150, 200, 300, 500},
	})
)

func NewRequestMetric() {
	requestCount.Inc()
}

func RequestDurationMetric(start time.Time) {
	duration := float64(time.Since(start).Milliseconds())

	requestDuration.Set(duration)
	requestDurationBucket.Observe(duration)
}

func RequestWithJsonDurationMetric(start time.Time) {
	duration := float64(time.Since(start).Milliseconds())

	requestWithJsonDuration.Set(duration)
	requestWithJsonDurationBucket.Observe(duration)
}
