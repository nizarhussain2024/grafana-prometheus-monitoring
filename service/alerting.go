package main

import (
	"sync"
	"time"
)

type AlertRule struct {
	Name        string
	Condition   func() bool
	Severity    string
	Threshold   float64
	Duration    time.Duration
	Firing      bool
	LastFired   time.Time
}

type AlertManager struct {
	mu    sync.RWMutex
	rules []*AlertRule
	alerts map[string]*Alert
}

type Alert struct {
	Rule      *AlertRule
	FiredAt   time.Time
	ResolvedAt *time.Time
	Message   string
}

var alertManager = &AlertManager{
	rules:  make([]*AlertRule, 0),
	alerts: make(map[string]*Alert),
}

func (am *AlertManager) AddRule(rule *AlertRule) {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.rules = append(am.rules, rule)
}

func (am *AlertManager) Evaluate() {
	am.mu.Lock()
	defer am.mu.Unlock()

	for _, rule := range am.rules {
		if rule.Condition() {
			if !rule.Firing {
				rule.Firing = true
				rule.LastFired = time.Now()
				am.alerts[rule.Name] = &Alert{
					Rule:    rule,
					FiredAt: time.Now(),
					Message: rule.Name + " alert fired",
				}
			}
		} else {
			if rule.Firing {
				rule.Firing = false
				if alert, exists := am.alerts[rule.Name]; exists {
					now := time.Now()
					alert.ResolvedAt = &now
				}
			}
		}
	}
}

func (am *AlertManager) GetActiveAlerts() []*Alert {
	am.mu.RLock()
	defer am.mu.RUnlock()

	active := make([]*Alert, 0)
	for _, alert := range am.alerts {
		if alert.Rule.Firing {
			active = append(active, alert)
		}
	}
	return active
}

func startAlertEvaluation() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			alertManager.Evaluate()
		}
	}()
}

