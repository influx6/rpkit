RPKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/rpkit)](https://goreportcard.com/report/github.com/gokit/rpkit)
[![Travis Build Status](https://travis-ci.org/gokit/rpkit.svg?branch=master)](https://travis-ci.org/gokit/rpkit#)
[![CircleCI](https://circleci.com/gh/gokit/rpkit.svg?style=svg)](https://circleci.com/gh/gokit/rpkit)

RPkit implements a code generator which automatically generates a RPC-style API package, which exposes methods of an interface as both http.Handlers with associated client code that directly make requests to such to perform action of said methods.

It is heavily inspired by the work on [Twirp](https://github.com/twitchtv/twirp/) by the [TwitchTV Team](https://github.com/twitchtv/people). Which inspires a beautiful approach for creating RPC-APIs that work with HTTP 1.1 and uses standard request-response with method names as targets.

RPKit takes these ideas and brings it into a more code friendly style, where there is no dependency on protobuff definitions, but rather on annotations marking interface declarations as code generation targets.

## Install

```
go get github.com/gokit/rpkit
```

## Examples

See [Examples](./examples) for demonstrations of packages generated using rpkit which creates a full RPC-style API with client code for a interface API definition/declaration.

## CLI

```bash
> rpkit generate
```

```bash
> rpkit
Usage: rpkit [flags] [command] 

⡿ COMMANDS:
	⠙ generate        Generates a rpc like API for interface types.

⡿ HELP:
	Run [command] help

⡿ OTHERS:
	Run 'rpkit flags' to print all flags of all commands.

⡿ WARNING:
	Uses internal flag package so flags must precede command name. 
	e.g 'rpkit -cmd.flag=4 run'

```


## How It works

1. Annotate all chosen interface with `@rp`
2. Run `rpkit generate` within package of annotated interfaces.

### Interface Based Annotation

You annotate any giving interface with `@rp` which marks giving interface has a target for code generation.

Sample below:

```go
// @rp
type Users interface {
	Create(NewUser) (User, error)
	UpdateEmail(SetUserEmail) error
}
```

## Important Rules

1. RPKit exposes a simple code generator using go's internal ast Package, but has specific requirements in the way it works with internal go package imports. Please always maintain the name the internal go package, that if your code uses the `context` package, then using the import alias `gctx "context"` will cause reference errors, as rpkit does not resolve this during it's code generation.
Aliasing can be done for external packages used by your code but not for internal Go packages.

