// Package sereal implements the Encoder and Decoders for use with the json format. See https://github.com/Sereal/Sereal
package sereal

import (
	"bytes"
	"context"
	"io"

	"sync"

	"github.com/Sereal/Sereal/Go/sereal"
)

var (
	// DefaultSerealEncoder provides a package-level json encoder for use.
	DefaultSerealEncoder SerealEncoder

	// DefaultSerealTargetDecoder provides a package-level json target decoder for use.
	DefaultSerealTargetDecoder SerealUnmarshalDecoder

	bytesPool = sync.Pool{New: func() interface{} {
		return new(bytes.Buffer)
	}}
)

// SerealEncoder implements a wrapper over the encoding/json SerealEncoder to
// match the Encoder interface. This allow us use Sereal has a encoder for
// any method encoding as needed.
// See // See https://github.com/Sereal/Sereal.
type SerealEncoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (SerealEncoder) Encode(ctx context.Context, w io.Writer, payload interface{}) error {
	data, err := sereal.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

// SerealUnmarshalDecoder implements a wrapper over the encoding/json SerealEncoder to
// all incoming data into a interface{} type.
type SerealUnmarshalDecoder struct{}

// Encode implements the necessary logic to use json for encoding to provided
// target of type interface{}.
func (SerealUnmarshalDecoder) Decode(ctx context.Context, r io.Reader, data interface{}) error {
	b := bytesPool.Get().(*bytes.Buffer)
	defer bytesPool.Put(b)

	_, err := io.Copy(b, r)
	if err != nil {
		return err
	}

	err = sereal.Unmarshal(b.Bytes(), data)
	return err
}
