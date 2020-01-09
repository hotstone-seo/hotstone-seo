package repository

import (
	"time"
)

// MetricsMismatchedCount represented  metrics_unmatched entity
type MetricsMismatchedCount struct {
	RequestPath string    `json:"request_path"`
	Count       int64     `json:"count"`
	Since       time.Time `json:"since"`
}
