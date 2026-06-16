package main

import (
	"fmt"
	"runtime/debug"

	"github.com/untrustedmodders/go-plugify"
)

var plugin plugify.PluginInfo

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
	var buildInfo, _ = debug.ReadBuildInfo()
	fmt.Printf("Plug: Version: %v\n", buildInfo)
	plugin = plugify.NewPlugin(buildInfo, OnPluginStart, nil, OnPluginEnd)
}
