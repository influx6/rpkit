// Package gobs implements the Encoder and Decoders for use with the gop format.
package gobs

import (
	"context"
	"encoding/gob"
	"io"
)

var (
	// DefaultGOBSEncoder provides a package-level json encoder for use.
	DefaultGOBSEncoder GOBSEncoder

	// DefaultGOBSDecoder provides a package-level json decoder for use.
	DefaultGOBSDecoder GOBSDecoder

	// DefaultGOBSTargetDecoder provides a package-level json target decoder for use.
	DefaultGOBSTargetDecoder GOBSUnmarshalDecoder
)

// GOBSEncoder implements a wrapper over the encoding/json GOBSEncoder to
// match the Encoder interface. This allow us use GOBS has a encoder for
// any method encoding as needed.
type GOBSEncoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (GOBSEncoder) Encode(ctx context.Context, w io.Writer, payload interface{}) error {
	return gob.NewEncoder(w).Encode(payload)
}

// GOBSDecoder implements a wrapper over the encoding/json GOBSEncoder to
// all incoming data into a interface{} type.
type GOBSDecoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (GOBSDecoder) Decode(ctx context.Context, r io.Reader) (interface{}, error) {
	var data interface{}
	err := gob.NewDecoder(r).Decode(&data)
	return data, err
}

// GOBSUnmarshalDecoder implements a wrapper over the encoding/json GOBSEncoder to
// all incoming data into a interface{} type.
type GOBSUnmarshalDecoder struct{}

// Encode implements the necessary logic to use json for encoding to provided
// target of type interface{}.
func (GOBSUnmarshalDecoder) Decode(ctx context.Context, r io.Reader, data interface{}) error {
	return gob.NewDecoder(r).Decode(data)
}
