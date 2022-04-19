# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/goutil?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/goutil)](https://github.com/whale-clouds/goutil)
[![GoDoc](https://godoc.org/github.com/whale-clouds/goutil?status.svg)](https://pkg.go.dev/github.com/whale-clouds/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/whale-clouds/goutil)](https://goreportcard.com/report/github.com/whale-clouds/goutil)
[![Unit-Tests](https://github.com/whale-clouds/goutil/workflows/Unit-Tests/badge.svg)](https://github.com/whale-clouds/goutil/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/goutil/badge.svg?branch=master)](https://coveralls.io/github/gookit/goutil?branch=master)

💪 Useful utils for the Go: string, array/slice, map, format, CLI, ENV, filesystem, testing and more.

- `arrutil` Array/Slice util functions. eg: check, convert
- `dump`  Simple variable printing tool, printing slice, map will automatically wrap each element and display the call location
- `cliutil` Command-line util functions. eg: read input, exec command, cmdline parse/build
- `envutil` ENV util for current runtime env information. eg: get one, get info, parse var
- `fmtutil` Format data util functions
- `fsutil` filesystem util functions. eg: file and dir check, operate
- `jsonutil` JSON util functions.
- `maputil` map util functions. eg: convert, sub-value get, simple merge
- `mathutil` math util functions. eg: convert, math calc, random
- `netutil` network util functions
- `strutil` string util functions. eg: bytes, check, convert, encode, format and more
- `sysutil` system util functions. eg: sysenv, exec, user, process
- `testutil` test help util functions. eg: http test, mock ENV value

> **[中文说明](README.zh-CN.md)**

## GoDoc

- [Godoc for github](https://pkg.go.dev/github.com/whale-clouds/goutil)

## Packages
{{pgkFuncs}}
## Code Check & Testing

```bash
gofmt -w -l ./
golint ./...
go test ./...
```

## Gookit packages

  - [gookit/ini](https://github.com/gookit/ini) Go config management, use INI files
  - [gookit/rux](https://github.com/gookit/rux) Simple and fast request router for golang HTTP
  - [gookit/gcli](https://github.com/gookit/gcli) Build CLI application, tool library, running CLI commands
  - [gookit/slog](https://github.com/gookit/slog) Lightweight, easy to extend, configurable logging library written in Go
  - [gookit/color](https://github.com/gookit/color) A command-line color library with true color support, universal API methods and Windows support
  - [gookit/event](https://github.com/gookit/event) Lightweight event manager and dispatcher implements by Go
  - [gookit/cache](https://github.com/gookit/cache) Generic cache use and cache manager for golang. support File, Memory, Redis, Memcached.
  - [gookit/config](https://github.com/gookit/config) Go config management. support JSON, YAML, TOML, INI, HCL, ENV and Flags
  - [gookit/filter](https://github.com/gookit/filter) Provide filtering, sanitizing, and conversion of golang data
  - [gookit/validate](https://github.com/gookit/validate) Use for data validation and filtering. support Map, Struct, Form data
  - [gookit/goutil](https://github.com/whale-clouds/goutil) Some utils for the Go: string, array/slice, map, format, cli, env, filesystem, test and more
  - More, please see https://github.com/gookit

## License

[MIT](LICENSE)
