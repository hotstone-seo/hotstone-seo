package repository

import (
	"time"
)

// MetricsUnmatchedCount represented  metrics_unmatched entity
type MetricsUnmatchedCount struct {
	RequestPath string    `json:"request_path"`
	Count       int64     `json:"count"`
	Since       time.Time `json:"since"`
}
