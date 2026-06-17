package main

import (
	"fmt"

	"github.com/untrustedmodders/go-plugify"
)

var plugin plugify.Plugin

func OnPluginStart() error {
	fmt.Println("Go: OnPluginStart")
	return nil
}

/*func OnPluginUpdate() {
    fmt.Println("Example: OnPluginUpdate")
    return nil
}*/

func OnPluginEnd() error {
	fmt.Println("Go: OnPluginEnd")
	return nil
}

func main() {}

func init() {
	plugin = plugify.NewPlugin("cross_call_worker", OnPluginStart, nil, OnPluginEnd)
}
