package main

import (
	"example_plugin/cross_call_worker"
	"fmt"

	"github.com/untrustedmodders/go-plugify"
)

var plugin plugify.Plugin

func OnPluginStart() error {
	fmt.Println("Example: OnPluginStart")

	strs := cross_call_worker.CallFunc33(func(a *any) {
		fmt.Println("Example: CallFunc33")
		*a = "some string lul"
	})
	fmt.Println("Example CallFunc33: ", strs)
	str := cross_call_worker.CallFuncEnum(func(p1 cross_call_worker.Example, p2 *[]cross_call_worker.Example) []cross_call_worker.Example {
		fmt.Println("Example: CallFuncEnum")
		p1 = cross_call_worker.Example_Third
		*p2 = []cross_call_worker.Example{
			cross_call_worker.Example_Third,
			cross_call_worker.Example_Third,
			cross_call_worker.Example_Third,
		}
		return []cross_call_worker.Example{cross_call_worker.Example_Forth}
	})
	fmt.Println("Example CallFuncEnum: ", str)

	for i := 0; i < 122; i++ {
		cross_call_worker.NoParamReturnVoid()
	}

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
	plugin = plugify.NewPlugin("example_plugin", OnPluginStart, nil, OnPluginEnd)
}
