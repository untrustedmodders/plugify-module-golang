[![Русский](https://img.shields.io/badge/Русский-%F0%9F%87%B7%F0%9F%87%BA-green?style=for-the-badge)](/managed/README_ru.md)

# Go Single-Runtime Host

This directory contains the **single-runtime host** for the Go Language Module. Instead of loading each Go plugin as a C shared library (the default mode), this host allows all Go plugins to run inside **one shared Go runtime** using Go's native [`plugin`](https://pkg.go.dev/plugin) package.

## How it works

In the default mode, each plugin is a standalone `c-shared` binary with its own Go runtime, garbage collector, and thread pool. The single-runtime host changes this: the host itself is loaded as a C shared library by Plugify, and all Go plugins are then loaded into that host as native Go plugins (`.so` files built with `-buildmode=plugin`). They all share the same Go runtime instance.

```
Plugify Core
    └── libplugify-module-golang.so   (language module)
    └── libplugify-host-golang.so     (this — loaded via LD_PRELOAD)
            ├── plugin_a.so           (Go plugin, buildmode=plugin)
            ├── plugin_b.so           (Go plugin, buildmode=plugin)
            └── plugin_c.so           (Go plugin, buildmode=plugin)
```

## Advantages

- Go plugins can **call each other directly** without going through Plugify's marshalling layer.
- All plugins share **one garbage collector** — lower memory overhead and fewer GC pauses.
- **Fewer OS threads** are allocated compared to running N independent Go runtimes.

## Limitations & Warnings

This mode uses Go's `plugin` package, which carries significant caveats. Read these carefully before deciding to use it.

> Sourced from the [official Go `plugin` package documentation](https://pkg.go.dev/plugin#hdr-Warnings):

- **Platform support is limited.** Go plugins are only supported on **Linux, FreeBSD, and macOS**. This makes them unsuitable for portable applications.
- **Race detector support is poor.** Even simple race conditions may not be detected automatically. See [golang.org/issue/24245](https://go.dev/issue/24245).
- **Deployment is more complex.** All parts of the program must be available in the correct filesystem locations at runtime.
- **Initialization order is harder to reason about.** Some packages may not be initialized until long after the application has started.
- **Security risk.** Bugs in plugin-loading code could be exploited to load dangerous or untrusted libraries.
- **Strict toolchain compatibility required.** All plugins and the host **must** be compiled with the exact same Go toolchain version, build tags, and relevant flags/environment variables. Any mismatch will cause **runtime crashes**.
- **Identical dependencies required.** All common dependencies of the host and plugins must be built from exactly the same source code.

> In practice, this means the host and all plugins must be built together by the same person or build system. If that is already your situation, you may find it simpler to blank-import all plugins and compile a single static binary instead.

For use cases where these constraints are too restrictive, consider the default **C shared library** mode or traditional IPC mechanisms (sockets, RPC, shared memory).

## Building the Host

The runtime host is **not distributed** as a pre-built binary — you must build it yourself to ensure it matches your plugins' toolchain and dependencies exactly.

From this directory:

```sh
go build -buildmode=c-shared -o libplugify-host-golang.so ./
```

## Deployment

Place the resulting `libplugify-host-golang.so` in the **same `bin` folder** as the language module binary:

```
bin/
├── libplugify-module-golang.so   ← language module
└── libplugify-host-golang.so     ← this file
```

### LD_PRELOAD (Required)

The host library **must** be preloaded before Plugify starts. If it is initialised too late, the Go runtime inside the host may not be ready when plugins attempt to load, causing panic.

Set `LD_PRELOAD` before launching Plugify:

```sh
LD_PRELOAD=/path/to/bin/libplugify-host-golang.so ./plugify
```

Or export it in your shell session / service unit:

```sh
export LD_PRELOAD=/path/to/bin/libplugify-host-golang.so
```

## Building Plugins for Single-Runtime Mode

Plugins targeting the single-runtime host must be built differently from the default C shared library mode.

Use `-buildmode=plugin` and the `-tags=plugin` build tag so that the generated bindings are compiled correctly:

```sh
go build -buildmode=plugin -tags=plugin -o my_plugin.so ./
```

> **Important:** All plugins and the host must use the same Go version, the same module dependencies (pinned to identical versions and source), and the same build tags. Mixing is not supported and will crash.

## Choosing the Right Mode

| | C Shared (default) | Single Runtime (this) |
|---|---|---|
| Platform support | All platforms | Linux, FreeBSD, macOS only |
| Cross-plugin calls | Via Plugify marshalling | Direct Go calls |
| GC / threads | One per plugin | Shared across all plugins |
| Build complexity | Low | High — all must match exactly |
| Portability | High | Low |
| Recommended for | General use | Performance-critical, controlled environments |****