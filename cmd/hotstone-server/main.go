package main

import "github.com/typical-go/typical-go/pkg/typapp"
import "github.com/hotstone-seo/hotstone-server/typical"
import _ "github.com/hotstone-seo/hotstone-server/internal/dependency"

func main() {
	typapp.Run(typical.Context)
}
