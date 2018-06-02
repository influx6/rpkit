package sereal

import (
	"context"
	"encoding/gob"
	"io"
)

var (
	// DefaultSerealEncoder provides a package-level json encoder for use.
	DefaultSerealEncoder SerealEncoder

	// DefaultSerealDecoder provides a package-level json decoder for use.
	DefaultSerealDecoder SerealDecoder

	// DefaultSerealTargetDecoder provides a package-level json target decoder for use.
	DefaultSerealTargetDecoder SerealUnmarshalDecoder
)

// SerealEncoder implements a wrapper over the encoding/json SerealEncoder to
// match the Encoder interface. This allow us use Sereal has a encoder for
// any method encoding as needed.
type SerealEncoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (SerealEncoder) Encode(ctx context.Context, w io.Writer, payload interface{}) error {
	sereal.NewEncoder()
	return gob.NewEncoder(w).Encode(payload)
}

// SerealDecoder implements a wrapper over the encoding/json SerealEncoder to
// all incoming data into a interface{} type.
type SerealDecoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (SerealDecoder) Decode(ctx context.Context, r io.Reader) (interface{}, error) {
	var data interface{}
	err := gob.NewDecoder(r).Decode(&data)
	return data, err
}

// SerealUnmarshalDecoder implements a wrapper over the encoding/json SerealEncoder to
// all incoming data into a interface{} type.
type SerealUnmarshalDecoder struct{}

// Encode implements the necessary logic to use json for encoding to provided
// target of type interface{}.
func (SerealUnmarshalDecoder) Decode(ctx context.Context, r io.Reader, data interface{}) error {
	return gob.NewDecoder(r).Decode(data)
}
