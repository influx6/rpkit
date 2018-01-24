package iboxcellarrp

// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gokit/rpkit/examples/bottles"
)

// errors ...
var (
	ErrDecodedUnknownType = errors.New("decoder returns unknown type")
)

//****************************************************************************
// Encoder and Decoder Types
// Source: github.com/gokit/rpkit/examples/bottles
//****************************************************************************

// Encoder defines an interface representing a generic encoder which
// expects a interface{} type has it's encoding target.
type Encoder interface {
	Encode(io.Writer, interface{}) error
}

// Decoder defines an interface representing a generic decoder which
// returning a interface{} type has it's response.
type Decoder interface {
	Decode(io.Reader) (interface{}, error)
}

// TargetDecoder defines an interface representing a target decoder which
// expects a type to decode into.
type TargetDecoder interface {
	Decode(io.Reader, interface{}) error
}

//****************************************************************************
// bottles.IBox Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/bottles
// Used By: bottles.IBoxCellar
//****************************************************************************

// IBoxTypeEncoder implements a encoder for the bottles.IBox type.
type IBoxTypeEncoder struct {
	Encoder Encoder
}

// Encode implements the encode function for type bottles.IBox used in gokit.rpkit.examples.bottles by
// calling the underline Encoder to handle the work.
func (en IBoxTypeEncoder) Encode(w io.Writer, payload bottles.IBox) error {
	return en.Encoder.Encode(w, payload)
}

// IBoxTypeDecoder implements a decoder for the bottles.IBox type.
type IBoxTypeDecoder struct {
	Decoder Decoder
}

// Decode implements the decode function for type bottles.IBox used in gokit.rpkit.examples.bottles by
// calling the underline Decoder to handle the work.
func (td IBoxTypeDecoder) Decode(r io.Reader) (bottles.IBox, error) {
	recs, err := td.Decoder.Decode(r)
	if err != nil {
		return recs.(bottles.IBox), err
	}

	// Type convert, so we have the right type, which is bottles.IBox.
	if erecs, ok := recs.(bottles.IBox); ok {
		return erecs, nil
	}

	return recs.(bottles.IBox), ErrDecodedUnknownType
}

//****************************************************************************
// string Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/bottles
// Used By: bottles.IBoxCellar
//****************************************************************************

// StringTypeEncoder implements a encoder for the string type.
type StringTypeEncoder struct {
	Encoder Encoder
}

// Encode implements the encode function for type string used in gokit.rpkit.examples.bottles by
// calling the underline Encoder to handle the work.
func (en StringTypeEncoder) Encode(w io.Writer, payload string) error {
	return en.Encoder.Encode(w, payload)
}

// StringTypeDecoder implements a decoder for the string type.
type StringTypeDecoder struct {
	Decoder Decoder
}

// Decode implements the decode function for type string used in gokit.rpkit.examples.bottles by
// calling the underline Decoder to handle the work.
func (td StringTypeDecoder) Decode(r io.Reader) (string, error) {
	recs, err := td.Decoder.Decode(r)
	if err != nil {
		return recs.(string), err
	}

	// Type convert, so we have the right type, which is string.
	if erecs, ok := recs.(string); ok {
		return erecs, nil
	}

	return recs.(string), ErrDecodedUnknownType
}

//****************************************************************************
// Base Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/bottles
//****************************************************************************

// JSONEncoder implements a wrapper over the encoding/json JSONEncoder to
// match the Encoder interface. This allow us use JSON has a encoder for
// any method encoding as needed.
type JSONEncoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (JSONEncoder) Encode(w io.Writer, payload interface{}) error {
	return json.NewEncoder(w).Encode(payload)
}

// JSONMapDecoder implements a wrapper over the encoding/json JSONEncoder to
// all incoming data into a interface{} type.
type JSONMapDecoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (JSONMapDecoder) Decode(r io.Reader) (interface{}, error) {
	var data interface{}
	err := json.NewDecoder(r).Decode(&data)
	return data, err
}
