package typical

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of hotstone-seo
var Descriptor = typcore.Descriptor{
	Name:    "hotstone-seo",
	Version: "0.0.1",
	Package: "github.com/hotstone-seo/hotstone-seo",

	App: typapp.New(application).
		WithDependency(
			server,
			redis,
			postgres,
		).
		WithPrepare(
			redis,
			postgres,
		),

	Build: typbuild.New().
		WithCommands(
			docker,
			readme,
			postgres,
			redis,
		),

	Configuration: typcfg.New().
		WithConfigure(
			application,
			server,
			redis,
			postgres,
		),
}
