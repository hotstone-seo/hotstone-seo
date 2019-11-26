package typical

import (
	"github.com/hotstone-seo/hotstone-server/app"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// Context of Project
var Context = &typctx.Context{
	Name:      "hotstone-server",
	Version:   "0.0.1",
	Package:   "github.com/hotstone-seo/hotstone-server",
	AppModule: app.Module(),
	Releaser: typrls.Releaser{
		Targets: []typrls.Target{"linux/amd64", "darwin/amd64"},
	},
}
