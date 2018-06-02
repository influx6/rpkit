// Package json implements the Encoder and Decoders for use with the json format.
package json

import (
	"context"
	"encoding/json"
	"io"
)

var (
	// DefaultJSONEncoder provides a package-level json encoder for use.
	DefaultJSONEncoder JSONEncoder

	// DefaultJSONDecoder provides a package-level json decoder for use.
	DefaultJSONDecoder JSONDecoder

	// DefaultJSONTargetDecoder provides a package-level json target decoder for use.
	DefaultJSONTargetDecoder JSONUnmarshalDecoder
)

// JSONEncoder implements a wrapper over the encoding/json JSONEncoder to
// match the Encoder interface. This allow us use JSON has a encoder for
// any method encoding as needed.
type JSONEncoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (JSONEncoder) Encode(ctx context.Context, w io.Writer, payload interface{}) error {
	return json.NewEncoder(w).Encode(payload)
}

// JSONDecoder implements a wrapper over the encoding/json JSONEncoder to
// all incoming data into a interface{} type.
type JSONDecoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (JSONDecoder) Decode(ctx context.Context, r io.Reader) (interface{}, error) {
	var data interface{}
	err := json.NewDecoder(r).Decode(&data)
	return data, err
}

// JSONUnmarshalDecoder implements a wrapper over the encoding/json JSONEncoder to
// all incoming data into a interface{} type.
type JSONUnmarshalDecoder struct{}

// Encode implements the necessary logic to use json for encoding to provided
// target of type interface{}.
func (JSONUnmarshalDecoder) Decode(ctx context.Context, r io.Reader, data interface{}) error {
	return json.NewDecoder(r).Decode(data)
}
