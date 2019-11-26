package app

import "fmt"

// Module of application
func Module() interface{} {
	return &module{}
}

type module struct{}

func (module) Action() interface{} {
	return func() {
		fmt.Println("Hello World")
	}
}
