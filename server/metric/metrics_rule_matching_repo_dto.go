package metric

import (
	"time"
)

// NotMatchedReport contain sumary of total count of not-matching rule
type NotMatchedReport struct {
	URL       string    `json:"url"`
	Count     int64     `json:"count"`
	FirstSeen time.Time `json:"first_seen"`
	LastSeen  time.Time `json:"last_seen"`
}

// DailyReport contain sumary of daily hit
type DailyReport struct {
	Date     time.Time `json:"date"`
	HitCount int64     `json:"count"`
}
