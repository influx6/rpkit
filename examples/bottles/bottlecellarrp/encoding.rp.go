package bottlecellarrp

// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.

import (
	"encoding/json"
	"errors"
	"io"

	cellar "github.com/gokit/rpkit/examples/bottles/cellar"
)

// errors ...
var (
	ErrDecodedUnknownType = errors.New("decoder returns unknown type")
)

//****************************************************************************
// Encoders
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

//****************************************************************************
// RP: Output Returning No Error Decoders
// Method: Cellar
// Source: github.com/gokit/rpkit/examples/bottles
// Handler: bottles.BottleCellar.Cellar
//****************************************************************************

// CellarDecoder implements the CellarEncoder by providing a wrapper around
// a implementation of the Encoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type CellarEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by CellarEncoder.
func (ew CellarEncoderWrapper) Encode(w io.Writer, payload cellar.Cellar) error {
	return ew.Encoder.Encode(w, payload)
}

// CellarClientDecoderWrapper implements the CellarClientDecoder by providing a wrapper around
// a implementation of the ClientDecoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type CellarClientDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by CellarClientDecoder.
func (cdw CellarClientDecoderWrapper) Decode(r io.Reader) (cellar.Cellar, error) {
	recs, err := cdw.Decoder.Decode(r)
	if err != nil {
		return cellar.Cellar{}, err
	}

	if wantedRecs, ok := recs.(cellar.Cellar); ok {
		return wantedRecs, nil
	}

	return cellar.Cellar{}, ErrDecodedUnknownType
}

// CellarClientJSONDecoder implements a wrapper over the encoding/json JSONDecoder to
// match the Decoder interface. This allow us use JSON has a decoding for any method decoding as
// needed. This specifically decodes incoming data into cellar.Cellar type.
type CellarClientJSONDecoder struct{}

// Decode implements the necessary logic to use json for decoding.
func (CellarClientJSONDecoder) Decode(r io.Reader) (interface{}, error) {
	var res cellar.Cellar
	err := json.NewDecoder(r).Decode(&res)
	return res, err
}

//****************************************************************************
// RP: Input Returning Only Error Encoders
// Method: AddBottle
// Source: github.com/gokit/rpkit/examples/bottles
// Handler: bottles.BottleCellar.AddBottle
//****************************************************************************
// AddBottleDecoder implements the AddBottleDecoder by providing a wrapper around
// a implementation of the Decoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type AddBottleDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by AddBottleDecoder.
func (dw AddBottleDecoderWrapper) Decode(r io.Reader) (string, error) {

	recs, err := dw.Decoder.Decode(r)
	if err != nil {
		return "", err
	}

	if wantedRecs, ok := recs.(string); ok {
		return wantedRecs, nil
	}

	return "", ErrDecodedUnknownType
}

// AddBottleJSONDecoder implements a wrapper over the encoding/json JSONDecoder to
// match the Decoder interface. This allow us use JSON has a decoding for any method decoding as
// needed. This specifically decodes incoming data into wanted type.
type AddBottleJSONDecoder struct{}

// Decode implements the necessary logic to use json for decoding.
func (AddBottleJSONDecoder) Decode(r io.Reader) (interface{}, error) {

	var res string
	err := json.NewDecoder(r).Decode(&res)
	return res, err
}

// AddBottleClientEncoderWrapper implements the AddBottleClientEncoder by providing a wrapper around
// a implementation of the ClientEncoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type AddBottleClientEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by AddBottleClientEncoder.
func (cew AddBottleClientEncoderWrapper) Encode(w io.Writer, payload string) error {
	return cew.Encoder.Encode(w, payload)
}
