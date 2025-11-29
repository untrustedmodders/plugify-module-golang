//go:generate go run generator.go -package=main -output=.
//go:build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/untrustedmodders/go-plugify"
)

func main() {
	var (
		patterns     = flag.String("patterns", "./...", "Package patterns to analyze")
		output       = flag.String("output", "", "Output manifest file (default: <packagename>.pplugin)")
		name         = flag.String("name", "", "Plugin name (default: package name)")
		version      = flag.String("version", "1.0.0", "Plugin version")
		description  = flag.String("description", "", "Plugin description")
		author       = flag.String("author", "", "Plugin author")
		website      = flag.String("website", "", "Plugin website")
		license      = flag.String("license", "", "Plugin license")
		platforms    = flag.String("platforms", "", "Comma-separated list of platforms (e.g., windows,linux,darwin)")
		dependencies = flag.String("dependencies", "", "Comma-separated list of dependencies (e.g., plugin1,plugin2)")
		conflicts    = flag.String("conflicts", "", "Comma-separated list of conflicts (e.g., plugin3,plugin4)")
		entry        = flag.String("entry", "", "Plugin entry point (default: <packagename>)")
		target       = flag.String("package", "main", "Autoexports package (default: main)")
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

	// Parse comma-separated strings
	platformList := parseCommaSeparated(*platforms)
	dependencyList := parseCommaSeparated(*dependencies)
	conflictList := parseCommaSeparated(*conflicts)

	// Call the generator with error handling
	err := plugify.Generate(*patterns, *output, *name, *version, *description, *author, *website, *license, platformList, dependencyList, conflictList, *entry, *target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating plugin manifest: %v\n", err)
		os.Exit(1)
	}
}

// parseCommaSeparated parses a comma-separated string into a slice of trimmed strings
func parseCommaSeparated(input string) []string {
	if input == "" {
		return nil
	}

	parts := strings.Split(input, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}
