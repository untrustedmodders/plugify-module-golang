package main

import (
	"fmt"

	"github.com/untrustedmodders/go-plugify"
)

var plugin plugify.Plugin

func onPluginStart() error {
	fmt.Println("Go: OnPluginStart")
	return nil
}

/*func onPluginUpdate() {
    fmt.Println("Go: OnPluginUpdate")
    return nil
}*/

func onPluginEnd() error {
	fmt.Println("Go: OnPluginEnd")
	return nil
}

func main() {}

func init() {
	plugin = plugify.NewPlugin("cross_call_worker", onPluginStart, nil, onPluginEnd)
}
