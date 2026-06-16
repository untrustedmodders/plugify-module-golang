package main

import (
	"example_plugin/cross_call_worker"
	"fmt"
	"runtime/debug"

	"github.com/untrustedmodders/go-plugify"
)

var plugin plugify.PluginInfo

func OnPluginStart() error {
	fmt.Println("Example: OnPluginStart")
	str := cross_call_worker.CallFuncEnum(func(p1 cross_call_worker.Example, p2 *[]cross_call_worker.Example) []cross_call_worker.Example {
		fmt.Println("Example: CallFuncEnum")
		return []cross_call_worker.Example{cross_call_worker.Example_Forth}
	})
	fmt.Println("Example after: ", str)
	return nil
}

/*func OnPluginUpdate() {
	fmt.Println("Example: OnPluginUpdate")
	return nil
}*/

func OnPluginEnd() error {
	fmt.Println("Example: OnPluginEnd")
	return nil
}

func main() {}

func init() {
	var buildInfo, _ = debug.ReadBuildInfo()
	plugin = plugify.NewPlugin(buildInfo, OnPluginStart, nil, OnPluginEnd)
}
