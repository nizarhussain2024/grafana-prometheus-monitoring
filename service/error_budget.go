package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	errorBudgetRemaining = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "slo_error_budget_remaining",
			Help: "Remaining error budget percentage",
		},
		[]string{"service", "slo_name"},
	)

	errorBudgetConsumed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "slo_error_budget_consumed",
			Help: "Consumed error budget percentage",
		},
		[]string{"service", "slo_name"},
	)

	errorBudgetBurnRate = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "slo_error_budget_burn_rate",
			Help: "Error budget burn rate (errors per hour)",
		},
		[]string{"service", "slo_name"},
	)
)

func init() {
	prometheus.MustRegister(errorBudgetRemaining)
	prometheus.MustRegister(errorBudgetConsumed)
	prometheus.MustRegister(errorBudgetBurnRate)
}

func calculateErrorBudget(service, sloName string, availabilityTarget, currentAvailability float64) {
	// Error budget = 1 - availability target
	errorBudget := 1.0 - availabilityTarget
	
	// Consumed = (target - current) / error budget
	consumed := (availabilityTarget - currentAvailability) / errorBudget
	if consumed < 0 {
		consumed = 0
	}
	
	remaining := 1.0 - consumed
	if remaining < 0 {
		remaining = 0
	}

	errorBudgetRemaining.WithLabelValues(service, sloName).Set(remaining)
	errorBudgetConsumed.WithLabelValues(service, sloName).Set(consumed)
}

func recordBurnRate(service, sloName string, rate float64) {
	errorBudgetBurnRate.WithLabelValues(service, sloName).Set(rate)
}


