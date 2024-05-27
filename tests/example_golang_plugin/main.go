package main

import "C"
import (
	"fmt"
	"plugify-plugin/cpptest"
	"plugify-plugin/plugify"
)

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("Go:OnPluginStart")

		// Define some random values
        p1 := true
        p2 := int16(-123)
        p3 := int8(45)
        p4 := int16(678)
        p5 := int32(-9876)
        p6 := int64(1234567890)
        p7 := uint8(200)
        p8 := uint16(60000)
        p9 := uint32(3000000000)
        p10 := uint64(12345678901234567890)
        var p11 uintptr = 0 // Example of a pointer, replace with the actual address if needed
        p12 := float32(3.14)
        p13 := float64(6.28)

        // Call ParamAllPrimitives with the random values
        result := cpptest.ParamAllPrimitives(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13)

        fmt.Println("Result:", result)
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
