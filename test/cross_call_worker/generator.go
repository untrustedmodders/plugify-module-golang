//go:generate generator.go -package=main -output=.
//go:build ignore

package main

import "github.com/untrustedmodders/go-plugify"

func main() {
	plugify.Generate("main", "./")
}
