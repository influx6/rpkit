// Package fn implements the Encoder and Decoders for use with the provided functions.
package fn

import (
	"context"
	"errors"
	"io"
)

// FuncEncoder implements a wrapper over the custom encoding function
// that handles the logic to encode giving type.
type FuncEncoder struct {
	EncodeFunc func(context.Context, io.Writer, interface{}) error
}

// Encode calls the necessary function internally to encode type.
func (f FuncEncoder) Encode(ctx context.Context, w io.Writer, t interface{}) error {
	if f.EncodeFunc == nil {
		return errors.New("not implemented")
	}
	return f.EncodeFunc(ctx, w, t)
}

// FuncDecoder implements a wrapper over the custom decoding function
// that handles the logic to decode giving data into a type from provided reader.
type FuncDecoder struct {
	DecodeFunc func(context.Context, io.Reader) (interface{}, error)
}

// Decode calls the necessary function internally to decode data to type.
func (f FuncDecoder) Decode(ctx context.Context, r io.Reader) (interface{}, error) {
	if f.DecodeFunc == nil {
		return nil, errors.New("not implemented")
	}
	return f.DecodeFunc(ctx, r)
}
