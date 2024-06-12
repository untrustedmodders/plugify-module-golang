# Go Language Module for Plugify

The Go Language Module for Plugify enables developers to write plugins in Go for the Plugify framework. This module provides a seamless integration for Go plugins, allowing them to be dynamically loaded and managed by the Plugify core.

## Features

- **Go Plugin Support**: Write your plugins in Go and integrate them effortlessly with the Plugify framework.
- **Automatic Exporting**: Easily export and import methods between plugins and the language module.
- **Initialization and Cleanup**: Handle plugin initialization, startup, and cleanup with dedicated module events.
- **Interoperability**: Communicate with plugins written in other languages through auto-generated interfaces.

## Getting Started

### Prerequisites

- Go Compiler
- Plugify Framework Installed

### Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/untrustedmodders/go-lang-module.git
    cd go-lang-module
    git submodule update --init --recursive
    ```

2. Build the Go language module:

    ```bash
    mkdir build && cd build
    cmake ..
    cmake --build .
    ```

### Usage

1. **Integration with Plugify**

   Ensure that your Go language module is available in the same directory as your Plugify setup.

2. **Write Go Plugins**

   Develop your plugins in Go using the Plugify Go API. Refer to the [Plugify Go Plugin Guide](https://docs.plugify.io/go-plugin-guide) for detailed instructions.

3. **Build and Install Plugins**

   Compile your Go plugins and place the shared libraries in a directory accessible to the Plugify core.

4. **Run Plugify**

   Start the Plugify framework, and it will dynamically load your Go plugins.

## Example

### Initialize your module

```sh
go mod init example.com/my-go-plugin
```

### Get the go-plugify module

Note that you need to include the v in the version tag.

```sh
go get github.com/untrustedmodders/go-plugify@v1.0.0
```

```go
package main

import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
)

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("OnPluginStart")
	})

	plugify.OnPluginEnd(func() {
		fmt.Println("OnPluginEnd")
	})
}

func main() {}
```

## Documentation

For comprehensive documentation on writing plugins in Go using the Plugify framework, refer to the [Plugify Documentation](https://docs.plugify.io).

## Contributing

Feel free to contribute by opening issues or submitting pull requests. We welcome your feedback and ideas!

## License

This Go Language Module for Plugify is licensed under the [MIT License](LICENSE).
