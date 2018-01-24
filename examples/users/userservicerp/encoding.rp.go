package userservicerp

// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gokit/rpkit/examples/users"
)

// errors ...
var (
	ErrDecodedUnknownType = errors.New("decoder returns unknown type")
)

//****************************************************************************
// Encoders
// Source: github.com/gokit/rpkit/examples/users
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
// Source: github.com/gokit/rpkit/examples/users
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
// RP: Input And Output Returning Error Encoders/Decoders
// Method: Create
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Create
//****************************************************************************

// CreateDecoder implements the CreateDecoder by providing a wrapper around
// a implementation of the Decoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type CreateDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by CreateDecoder.
func (dw CreateDecoderWrapper) Decode(r io.Reader) (users.NewUser, error) {
	recs, err := dw.Decoder.Decode(r)
	if err != nil {
		return users.NewUser{}, err
	}

	if wantedRecs, ok := recs.(users.NewUser); ok {
		return wantedRecs, nil
	}

	return users.NewUser{}, ErrDecodedUnknownType
}

// CreateJSONDecoder implements a wrapper over the encoding/json JSONDecoder to
// match the Decoder interface. This allow us use JSON has a decoding for any method decoding as
// needed. This specifically decodes incoming data into wanted type.
type CreateJSONDecoder struct{}

// Decode implements the necessary logic to use json for decoding.
func (CreateJSONDecoder) Decode(r io.Reader) (interface{}, error) {
	var res users.NewUser
	err := json.NewDecoder(r).Decode(&res)
	return res, err
}

// CreateClientDecoderWrapper implements the CreateClientDecoder by providing a wrapper around
// a implementation of the ClientDecoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type CreateClientDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by CreateClientDecoder.
func (cdw CreateClientDecoderWrapper) Decode(r io.Reader) (users.User, error) {
	recs, err := cdw.Decoder.Decode(r)
	if err != nil {
		return users.User{}, err
	}

	if wantedRecs, ok := recs.(users.User); ok {
		return wantedRecs, nil
	}

	return users.User{}, ErrDecodedUnknownType
}

// CreateDecoder implements the CreateEncoder by providing a wrapper around
// a implementation of the Encoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type CreateEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by CreateEncoder.
func (ew CreateEncoderWrapper) Encode(w io.Writer, payload users.User) error {
	return ew.Encoder.Encode(w, payload)
}

// CreateClientEncoderWrapper implements the CreateClientEncoder by providing a wrapper around
// a implementation of the ClientEncoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type CreateClientEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by CreateClientEncoder.
func (cew CreateClientEncoderWrapper) Encode(w io.Writer, payload users.NewUser) error {
	return cew.Encoder.Encode(w, payload)
}
