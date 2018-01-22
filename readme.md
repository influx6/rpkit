RPKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/rpkit)](https://goreportcard.com/report/github.com/gokit/rpkit)

Mgokit implements a code generator which automatically generates go packages for mongodb implementations for annotated structs.

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
type Users struct {
	Create(NewUser) (User, error)
	UpdateEmail(SetUserEmail) error
}
```


