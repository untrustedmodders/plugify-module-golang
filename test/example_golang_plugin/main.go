package main

import "C"
import (
	"fmt"
	"plugify-plugin/plugify"
)

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("Go:OnPluginStart")
	})

	plugify.OnPluginEnd(func() {
		fmt.Println("Go:OnPluginEnd")
	})
}

//export SayHello
func SayHello(count int64) bool {
	if count > 10 {
		return false
	}

	fmt.Printf("Hello %d times!\n", count)

	return true
}

func main() {}
