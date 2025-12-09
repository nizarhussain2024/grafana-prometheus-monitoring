package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "HTTP request duration in seconds",
		},
		[]string{"method", "endpoint"},
	)
)

var (
	activeConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Number of active connections",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(activeConnections)
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","service":"prometheus-monitoring","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
		
		duration := time.Since(start).Seconds()
		httpRequestDuration.WithLabelValues(r.Method, "/health").Observe(duration)
		httpRequestsTotal.WithLabelValues(r.Method, "/health", "200").Inc()
		activeConnections.Inc()
		defer activeConnections.Dec()
	})

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"Data endpoint","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
		
		duration := time.Since(start).Seconds()
		httpRequestDuration.WithLabelValues(r.Method, "/api/data").Observe(duration)
		httpRequestsTotal.WithLabelValues(r.Method, "/api/data", "200").Inc()
	})

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Grafana Prometheus Monitoring service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
