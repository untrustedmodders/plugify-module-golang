package main

import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
)

func main() {}

func init() {
	plugify.OnPluginStart(func() error {
		fmt.Println("Go: OnPluginStart")
		return nil
	})

	plugify.OnPluginEnd(func() error {
		fmt.Println("Go: OnPluginEnd")
		return nil
	})
}
