package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestSummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_summary_seconds",
			Help: "HTTP request summary in seconds",
			Objectives: map[float64]float64{
				0.5:  0.05,
				0.9:  0.01,
				0.95: 0.01,
				0.99: 0.001,
			},
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestSummary)
}




