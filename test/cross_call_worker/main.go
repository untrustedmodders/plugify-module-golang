package main

import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
	"runtime/debug"
)

func main() {}

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("Go: OnPluginStart")
	})

	plugify.OnPluginEnd(func() {
		fmt.Println("Go: OnPluginEnd")
	})

	plugify.OnPluginPanic(func() []byte {
		return debug.Stack() // workaround for could not import runtime/debug inside plugify package
	})
}
