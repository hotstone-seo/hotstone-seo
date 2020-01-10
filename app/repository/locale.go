package repository

import (
	"context"
	"time"
)

// Locale represented  locale entity
type Locale struct {
	ID          int64     `json:"id"`
	LangCode    string    `json:"lang_code"`
	CountryCode string    `json:"country_code"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// LocaleRepo to handle locales entity [mock]
type LocaleRepo interface {
	FindOne(context.Context, int64) (*Locale, error)
	Find(context.Context) ([]*Locale, error)
	Insert(context.Context, Locale) (lastInsertID int64, err error)
	Delete(context.Context, int64) error
	Update(context.Context, Locale) error
}

// NewLocaleRepo return new instance of LocaleRepo [autowire]
func NewLocaleRepo(impl CachedLocaleRepoImpl) LocaleRepo {
	return &impl
}
