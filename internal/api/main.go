package api

import (
	"github.com/hotstone-seo/hotstone-seo/internal/config"
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
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

// Main function to run server
func Main(s server, m urlstore.Worker) (err error) {
	if err = m.Start(); err != nil {
		return
	}
	if err = startServer(s); err != nil {
		return
	}
	return
}
