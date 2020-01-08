package main

import (
	"github.com/hotstone-seo/hotstone-seo/typical"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

func main() {
	typbuildtool.Run(&typical.Descriptor)
}
