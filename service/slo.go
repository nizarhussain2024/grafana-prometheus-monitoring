package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Service Level Indicators
	requestSuccessRate = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sli_request_success_rate",
			Help: "Request success rate (SLI)",
		},
		[]string{"service"},
	)

	requestLatencyP99 = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sli_request_latency_p99_seconds",
			Help: "99th percentile request latency (SLI)",
		},
		[]string{"service"},
	)

	// Service Level Objectives
	sloAvailability = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "slo_availability_target",
			Help: "Availability SLO target (e.g., 0.99 for 99%)",
		},
		[]string{"service"},
	)

	sloLatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "slo_latency_target_seconds",
			Help: "Latency SLO target in seconds",
		},
		[]string{"service"},
	)
)

func init() {
	prometheus.MustRegister(requestSuccessRate)
	prometheus.MustRegister(requestLatencyP99)
	prometheus.MustRegister(sloAvailability)
	prometheus.MustRegister(sloLatency)
}

func recordSLI(service string, successRate float64, latencyP99 float64) {
	requestSuccessRate.WithLabelValues(service).Set(successRate)
	requestLatencyP99.WithLabelValues(service).Set(latencyP99)
}

func setSLO(service string, availabilityTarget float64, latencyTarget float64) {
	sloAvailability.WithLabelValues(service).Set(availabilityTarget)
	sloLatency.WithLabelValues(service).Set(latencyTarget)
}



