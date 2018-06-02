// Package protobufs implements the Encoder and Decoders for use with the protobuf format. See https://github.com/gogo/protobuf
package protobufs

import (
	"context"
	"io"

	"encoding/json"

	"bytes"

	"github.com/gogo/protobuf/proto"
)

var (
	// DefaultProtobufEncoder provides a package-level json encoder for use.
	DefaultProtobufEncoder ProtobufEncoder

	// DefaultProtobufTargetDecoder provides a package-level json target decoder for use.
	DefaultProtobufTargetDecoder ProtobufUnmarshalDecoder
)

// ProtobufEncoder implements a wrapper over the encoding with protobuffers definitions to
// match the Encoder interface. This allow us use Protobuf has a encoder for
// any method encoding as needed.
type ProtobufEncoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (ProtobufEncoder) Encode(ctx context.Context, w io.Writer, payload interface{}) error {
	if pb, ok := payload.(proto.Message); ok {
		data, err := proto.Marshal(pb)
		if err != nil {
			return err
		}

		_, err = w.Write(data)
		return err
	}

	// if not a proto file, default to json,
	return json.NewEncoder(w).Encode(payload)
}

// ProtobufUnmarshalDecoder implements a wrapper over the encoding/json ProtobufEncoder to
// all incoming data into a interface{} type.
type ProtobufUnmarshalDecoder struct{}

// Encode implements the necessary logic to use json for encoding to provided
// target of type interface{}.
func (ProtobufUnmarshalDecoder) Decode(ctx context.Context, r io.Reader, data interface{}) error {
	if pb, ok := data.(proto.Message); ok {
		var b bytes.Buffer
		_, err := io.Copy(&b, r)
		if err != nil {
			return err
		}

		err = proto.Unmarshal(b.Bytes(), pb)
		if err != nil {
			return err
		}

		return err
	}

	return json.NewDecoder(r).Decode(data)
}
