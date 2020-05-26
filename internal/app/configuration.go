package app

import (
	"github.com/hotstone-seo/hotstone-seo/internal/app/config"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	configName = "APP"
)

// Configuration of server
func Configuration() *typgo.Configuration {
	return &typgo.Configuration{
		Name: configName,
		Spec: &config.Config{},
	}
}
