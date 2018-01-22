RPKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/rpkit)](https://goreportcard.com/report/github.com/gokit/rpkit)

RPkit implements a code generator which automatically generates a RPC-style API package, which exposes methods of an interface as both http.Handlers with associated client code that directly make requests to such to perform action of said methods.

It is heavily inspired by the work of the [TwitchTV Team](https://github.com/twitchtv/people) work on [Twirp](https://github.com/twitchtv/twirp/) which inspires a beautiful approach for creating RPC-APIs that easily work with HTTP 1.1 easily.

RPKit takes the idea and brings it into a more code friendly approach where there is no dependency to protobuff definition, but rather be generated from in-code annotations of interfaces.

## Install

```
go get github.com/gokit/rpkit
```

## CLI


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


