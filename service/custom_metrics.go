package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	businessMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "business_operations_total",
			Help: "Total number of business operations",
		},
		[]string{"operation_type", "status"},
	)

	processingTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "processing_time_seconds",
			Help:    "Time spent processing operations",
			Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
		},
		[]string{"operation_type"},
	)
)

func init() {
	prometheus.MustRegister(businessMetrics)
	prometheus.MustRegister(processingTime)
}

func recordBusinessOperation(opType, status string) {
	businessMetrics.WithLabelValues(opType, status).Inc()
}

func recordProcessingTime(opType string, duration float64) {
	processingTime.WithLabelValues(opType).Observe(duration)
}



