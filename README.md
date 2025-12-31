# YaegiDemo

Demo project showing how to use [Yaegi](https://github.com/traefik/yaegi) to load and execute Go code at runtime.

English | [中文](README.zh_CN.md)

## Overview

The HTTP handler is defined in `rule/rule.go` and hot-reloaded by Yaegi.

## Requirements

- Go 1.22+

## Quick Start

Start the demo server:

```bash
make run
```

Send a request:

```bash
curl localhost:8080 -i
```

Example response (truncated):

```bash
HTTP/1.1 200 OK
Cache-Control: no-cache
Content-Type: text/plain; charset=utf-8
Via: useless-helper
Date: Thu, 23 Nov 2023 12:24:31 GMT
Content-Length: 241

Hello!
Your Content-Type is
Your User-Agent is curl/8.1.2
Your Host is localhost:8080
Your RemoteAddr is 127.0.0.1:61784
Your RequestURI is /
Your Method is GET
Your Proto is HTTP/1.1
Your URL is /
Code above is wrote by GitHub Copilot.

:)
```

## Hot Reload

Edit `rule/rule.go` and save; the server response updates without restart.

## Configuration

- `--rule` (default: `./rule/rule.go`): rule file path

## Development

- `make generate`: regenerate Yaegi symbols
- `make build`: build binary into `output/yaegi_demo`

## Project Tree

```text
.
├── cmd
│   └── main.go         # Program entry
├── internal
│   ├── engine          # Yaegi execution & reload
│   ├── helper          # Demo helpers
│   └── symbol          # Interpreter symbols
├── rule
│   └── rule.go          # Hot-reloaded handler
├── Makefile             # Dev commands
├── README.md            # Project overview
├── README.zh_CN.md      # Chinese README
├── go.mod               # Module definition
└── go.sum               # Dependency checksums
```
