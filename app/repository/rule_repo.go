package repository

import "time"

type Rule struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name" validate:"required"`
	UrlPattern   string    `json:"url_pattern" validate:"required"`
	Exclusion    *string   `json:"exclusion"`
	IdDataSource int64     `validate:"required"`
	UpdatedAt    time.Time `json:"-"`
	CreatedAt    time.Time `json:"-"`
}
