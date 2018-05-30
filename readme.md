RPKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/rpkit)](https://goreportcard.com/report/github.com/gokit/rpkit)
[![Travis Build Status](https://travis-ci.org/gokit/rpkit.svg?branch=master)](https://travis-ci.org/gokit/rpkit#)
[![CircleCI](https://circleci.com/gh/gokit/rpkit.svg?style=svg)](https://circleci.com/gh/gokit/rpkit)

RPkit implements a code generator which automatically generates a RPC-style API package for you in Go and Javascript (Browser and Nodejs compatible code), which exposes methods of an interface as RPC styled methods with associated client functions. Enabling you rapidly build and prototype your API with minimal effort.

It is heavily inspired by the work on [Twirp](https://github.com/twitchtv/twirp/) by the [TwitchTV Team](https://github.com/twitchtv/people). Which inspires a beautiful approach for creating RPC-APIs that work with HTTP 1.1 and uses standard request-response with method names as targets.

RPKit takes these ideas and brings it into a friendly annotation style used in Go source ocde, where there is no dependency on protobuf definitions.

## Install

```
go get github.com/gokit/rpkit
```

## Examples
When using RPkit, you simply need to have `go get` the package first, then create a annotate interface within your target golang
source package. e.g


```go
package users

import "context"

type NewUser struct {
	Name string `json:"name"`
}

type User struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Addr string  `json:"addr"`
	Cid  float64 `json:"cid"`
}

//@rp(js:server, js:client)
//@implement
type UserService interface {
	Poke()
	PokeAgain() error
	Get(context.Context) (int, error)
	Create(context.Context, NewUser) (User, error)
	GetBy(context.Context, string) (int, error)
	CreateUser(NewUser) (User, error)
	GetUser(int) (User, error)
	GetUsers(context.Context) []User
}
```

If you look into the source code above, the annotation that marks the `UserService` interface has a target is `@rp`.
By specifying to `@rp` the `js:server` and `js:client`, the annotation will generate both javascript code that can be
used for both the server and client usage.


Once done, navigate to your source file directory so we could code generate the source files for both go and javascript.

See [Examples](./examples) for generated contents.

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

1. RPKit exposes a simple code generator using go's internal ast Package, but has specific requirements in the way it works with internal go package imports. Please always maintain the name of the internal go package (i.e if your code uses the `context` package, then using the import alias `gctx "context"` will cause reference errors, as rpkit does not resolve this during it's code generation). Aliasing can be done for external packages used by your code but not for internal Go packages.

