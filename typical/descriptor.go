package typical

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of hotstone-seo
var Descriptor = typcore.Descriptor{
	Name:    "hotstone-seo",
	Version: "0.0.1",

	App: typapp.New(serverApp).
		AppendProvider(
			socialLogin,
		).
		AppendDependency(
			redis,
			postgres,
		).
		AppendPreparer(
			redis,
			postgres,
		),

	BuildTool: typbuildtool.New().
		AppendCommander(
			docker,
			readme,
			postgres,
			redis,
		),

	Configuration: typcfg.New().
		AppendConfigurer(
			serverApp,
			redis,
			postgres,
			socialLogin,
		),
}
