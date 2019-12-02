// +build typical

package main

import "github.com/typical-go/typical-go/pkg/typbuildtool"
import "github.com/hotstone-seo/hotstone-server/typical"
import _ "github.com/hotstone-seo/hotstone-server/internal/dependency"

func main() {
	typbuildtool.Run(typical.Context)
}
