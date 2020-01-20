package service

import (
	"github.com/hotstone-seo/hotstone-seo/app/repository"
	"go.uber.org/dig"
)

// TagService contain logic for TagController [mock]
type TagService interface {
	repository.TagRepo
}

// TagServiceImpl is implementation of TagService
type TagServiceImpl struct {
	dig.In
	repository.TagRepo
}

// NewTagService return new instance of TagService [constructor]
func NewTagService(impl TagServiceImpl) TagService {
	return &impl
}
