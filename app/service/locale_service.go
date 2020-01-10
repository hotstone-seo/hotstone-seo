package service

import (
	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"go.uber.org/dig"
)

// LocaleService contain logic for LocaleController [mock]
type LocaleService interface {
	repository.LocaleRepo
}

// LocaleServiceImpl is implementation of LocaleService
type LocaleServiceImpl struct {
	dig.In
	repository.LocaleRepo
}

// NewLocaleService return new instance of LocaleService [autowire]
func NewLocaleService(impl LocaleServiceImpl) LocaleService {
	return &impl
}
