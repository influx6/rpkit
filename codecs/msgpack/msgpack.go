// Package msgpack implements the Encoder and Decoders for use with the MessagePack format.
package msgpack

import (
	"context"
	"io"

	"github.com/vmihailenco/msgpack"
)

var (
	// DefaultMsgPackEncoder provides a package-level json encoder for use.
	DefaultMsgPackEncoder MsgPackEncoder

	// DefaultMsgPackTargetDecoder provides a package-level json target decoder for use.
	DefaultMsgPackTargetDecoder MsgPackUnmarshalDecoder
)

// MsgPackEncoder implements a wrapper over the encoding/json MsgPackEncoder to
// match the Encoder interface. This allow us use MsgPack has a encoder for
// any method encoding as needed.
type MsgPackEncoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (MsgPackEncoder) Encode(ctx context.Context, w io.Writer, payload interface{}) error {
	return msgpack.NewEncoder(w).Encode(payload)
}

// MsgPackUnmarshalDecoder implements a wrapper over the encoding/json MsgPackEncoder to
// all incoming data into a interface{} type.
type MsgPackUnmarshalDecoder struct{}

// Encode implements the necessary logic to use json for encoding to provided
// target of type interface{}.
func (MsgPackUnmarshalDecoder) Decode(ctx context.Context, r io.Reader, data interface{}) error {
	return msgpack.NewDecoder(r).Decode(data)
}
