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
// Encoder and Decoder Types
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

// TargetDecoder defines an interface representing a target decoder which
// expects a type to decode into.
type TargetDecoder interface {
	Decode(io.Reader, interface{}) error
}

//****************************************************************************
// users.NewUser Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// NewUserTargetEncoder implements a encoder for the users.NewUser type
// using a provided function.
type NewUserTargetEncoder struct {
	EncoderFunc func(io.Writer, users.NewUser) error
}

// Encode implements the encoding function for type users.NewUser used in gokit.rpkit.examples.users by
// calling the underline encoding function to handle the work.
func (td NewUserTargetEncoder) Encode(w io.Writer, payload users.NewUser) error {
	return td.EncoderFunc(w, payload)
}

// NewUserTargetDecoder implements a decoder for the users.NewUser type
// using a provided function.
type NewUserTargetDecoder struct {
	DecoderFunc func(io.Reader) (users.NewUser, error)
}

// Decode implements the decoding function for type users.NewUser used in gokit.rpkit.examples.users by
// calling the underline decoding function to handle the work.
func (td NewUserTargetDecoder) Decode(r io.Reader) (users.NewUser, error) {
	return td.DecoderFunc(r)
}

// NewUserTypeEncoder implements a encoder for the users.NewUser type.
type NewUserTypeEncoder struct {
	Encoder Encoder
}

// Encode implements the encode function for type users.NewUser used in gokit.rpkit.examples.users by
// calling the underline Encoder to handle the work.
func (en NewUserTypeEncoder) Encode(w io.Writer, payload users.NewUser) error {
	return en.Encoder.Encode(w, payload)
}

// NewUserTypeDecoder implements a decoder for the users.NewUser type.
type NewUserTypeDecoder struct {
	Decoder TargetDecoder
}

// Decode implements the decode function for type users.NewUser used in gokit.rpkit.examples.users by
// calling the underline Decoder to handle the work.
func (td NewUserTypeDecoder) Decode(r io.Reader) (users.NewUser, error) {
	var res users.NewUser
	err := td.Decoder.Decode(r, &res)
	return res, err
}

//****************************************************************************
// users.User Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// UserTargetEncoder implements a encoder for the users.User type
// using a provided function.
type UserTargetEncoder struct {
	EncoderFunc func(io.Writer, users.User) error
}

// Encode implements the encoding function for type users.User used in gokit.rpkit.examples.users by
// calling the underline encoding function to handle the work.
func (td UserTargetEncoder) Encode(w io.Writer, payload users.User) error {
	return td.EncoderFunc(w, payload)
}

// UserTargetDecoder implements a decoder for the users.User type
// using a provided function.
type UserTargetDecoder struct {
	DecoderFunc func(io.Reader) (users.User, error)
}

// Decode implements the decoding function for type users.User used in gokit.rpkit.examples.users by
// calling the underline decoding function to handle the work.
func (td UserTargetDecoder) Decode(r io.Reader) (users.User, error) {
	return td.DecoderFunc(r)
}

// UserTypeEncoder implements a encoder for the users.User type.
type UserTypeEncoder struct {
	Encoder Encoder
}

// Encode implements the encode function for type users.User used in gokit.rpkit.examples.users by
// calling the underline Encoder to handle the work.
func (en UserTypeEncoder) Encode(w io.Writer, payload users.User) error {
	return en.Encoder.Encode(w, payload)
}

// UserTypeDecoder implements a decoder for the users.User type.
type UserTypeDecoder struct {
	Decoder TargetDecoder
}

// Decode implements the decode function for type users.User used in gokit.rpkit.examples.users by
// calling the underline Decoder to handle the work.
func (td UserTypeDecoder) Decode(r io.Reader) (users.User, error) {
	var res users.User
	err := td.Decoder.Decode(r, &res)
	return res, err
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

// JSONDecoder implements a wrapper over the encoding/json JSONEncoder to
// all incoming data into a interface{} type.
type JSONDecoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (JSONDecoder) Decode(r io.Reader) (interface{}, error) {
	var data interface{}
	err := json.NewDecoder(r).Decode(&data)
	return data, err
}

// JSONTargetDecoder implements a wrapper over the encoding/json JSONEncoder to
// all incoming data into a interface{} type.
type JSONTargetDecoder struct{}

// Encode implements the necessary logic to use json for encoding to provided
// target of type interface{}.
func (JSONTargetDecoder) Decode(r io.Reader, data interface{}) error {
	return json.NewDecoder(r).Decode(data)
}
