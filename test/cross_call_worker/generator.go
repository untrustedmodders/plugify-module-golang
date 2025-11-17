//go:generate generator.go -package=main -output=.
//go:build ignore

package main

import (
	"flag"
	"fmt"
	"github.com/untrustedmodders/go-plugify"
	"os"
)

func main() {
	var (
		patterns    = flag.String("patterns", "./...", "Package patterns to analyze")
		output      = flag.String("output", "", "Output manifest file (default: <packagename>.pplugin)")
		name        = flag.String("name", "", "Plugin name (default: package name)")
		version     = flag.String("version", "1.0.0", "Plugin version")
		description = flag.String("description", "", "Plugin description")
		author      = flag.String("author", "", "Plugin author")
		website     = flag.String("website", "", "Plugin website")
		license     = flag.String("license", "", "Plugin license")
		entry       = flag.String("entry", "", "Plugin entry point (default: <packagename>)")
		target      = flag.String("package", "main", "Autoexports package (default: main)")
	)

	flag.Parse()

	// Log what we're doing
	fmt.Println("Starting plugin manifest generation...")
	fmt.Printf("Package patterns: %s\n", *patterns)
	if *output != "" {
		fmt.Printf("Output file: %s\n", *output)
	}
	if *name != "" {
		fmt.Printf("Plugin name: %s\n", *name)
	}
	fmt.Printf("Version: %s\n", *version)

	// Call the generator with error handling
	err := plugify.Generate(*patterns, *output, *name, *version, *description, *author, *website, *license, *entry, *target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating plugin manifest: %v\n", err)
		os.Exit(1)
	}
}
