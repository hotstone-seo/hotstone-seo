// +build typical

package main

import "github.com/typical-go/typical-go/pkg/typprebuilder"
import "github.com/hotstone-seo/hotstone-server/typical"

func main() {
	typprebuilder.Run(typical.Context)
}
