package userservicerp

// All code below are auto-generated and should not be edited by hand.
// See https://github.com/gokit/rpkit for more info.
  
import (
	"encoding/json"
	"errors"
	"context"
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
type Encoder interface{
	Encode(context.Context, io.Writer, interface{}) error
}

// Decoder defines an interface representing a generic decoder which
// returning a interface{} type has it's response.
type Decoder interface{
	Decode(context.Context, io.Reader) (interface{}, error)
}

// TargetDecoder defines an interface representing a target decoder which
// expects a type to decode into.
type TargetDecoder interface{
	Decode(context.Context, io.Reader, interface{}) error
}


//****************************************************************************
// int Target Function Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// IntTargetEncoder implements a encoder for the int type
// using a provided function.
type IntTargetEncoder struct {
	EncoderFunc func(context.Context, io.Writer, int) error
}

// Encode implements the encoding function for type int used in users by
// calling the underline encoding function to handle the work.
func (td IntTargetEncoder) Encode(ctx context.Context, w io.Writer, payload int) error {
	return td.EncoderFunc(ctx, w, payload)
}

// IntTargetDecoder implements a decoder for the int type
// using a provided function.
type IntTargetDecoder struct {
	DecoderFunc func(context.Context,io.Reader) (int, error)
}

// Decode implements the decoding function for type int used in users by
// calling the underline decoding function to handle the work.
func (td IntTargetDecoder) Decode(ctx context.Context, r io.Reader) (int, error) {
	return td.DecoderFunc(ctx, r)
}

//****************************************************************************
// int Type Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// IntTypeEncoder implements a encoder for the int type.
type IntTypeEncoder struct {
	Encoder Encoder
}

// Encode implements the encode function for type int used in users by
// calling the underline Encoder to handle the work.
func (en IntTypeEncoder) Encode(ctx context.Context, w io.Writer, payload int) error {
	return en.Encoder.Encode(ctx, w, payload)
}


// IntTypeDecoder implements a decoder for the int type.
type IntTypeDecoder struct {
	Decoder Decoder
}

// Decode implements the decode function for type int used in users by
// calling the underline Decoder to handle the work.
func (td IntTypeDecoder) Decode(ctx context.Context,r io.Reader) (int, error) {
	recs, err := td.Decoder.Decode(ctx, r)
	if err != nil {
		return recs.(int), err
	}

	// Type convert, so we have the right type, which is int.
	if erecs, ok := recs.(int); ok {
		return erecs, nil
	}

    return recs.(int), ErrDecodedUnknownType
}


//****************************************************************************
// string Target Function Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// StringTargetEncoder implements a encoder for the string type
// using a provided function.
type StringTargetEncoder struct {
	EncoderFunc func(context.Context, io.Writer, string) error
}

// Encode implements the encoding function for type string used in users by
// calling the underline encoding function to handle the work.
func (td StringTargetEncoder) Encode(ctx context.Context, w io.Writer, payload string) error {
	return td.EncoderFunc(ctx, w, payload)
}

// StringTargetDecoder implements a decoder for the string type
// using a provided function.
type StringTargetDecoder struct {
	DecoderFunc func(context.Context,io.Reader) (string, error)
}

// Decode implements the decoding function for type string used in users by
// calling the underline decoding function to handle the work.
func (td StringTargetDecoder) Decode(ctx context.Context, r io.Reader) (string, error) {
	return td.DecoderFunc(ctx, r)
}

//****************************************************************************
// string Type Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// StringTypeEncoder implements a encoder for the string type.
type StringTypeEncoder struct {
	Encoder Encoder
}

// Encode implements the encode function for type string used in users by
// calling the underline Encoder to handle the work.
func (en StringTypeEncoder) Encode(ctx context.Context, w io.Writer, payload string) error {
	return en.Encoder.Encode(ctx, w, payload)
}


// StringTypeDecoder implements a decoder for the string type.
type StringTypeDecoder struct {
	Decoder Decoder
}

// Decode implements the decode function for type string used in users by
// calling the underline Decoder to handle the work.
func (td StringTypeDecoder) Decode(ctx context.Context,r io.Reader) (string, error) {
	recs, err := td.Decoder.Decode(ctx, r)
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
// users.NewUser Target Function Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// NewUserTargetEncoder implements a encoder for the users.NewUser type
// using a provided function.
type NewUserTargetEncoder struct {
	EncoderFunc func(context.Context, io.Writer, users.NewUser) error
}

// Encode implements the encoding function for type users.NewUser used in users by
// calling the underline encoding function to handle the work.
func (td NewUserTargetEncoder) Encode(ctx context.Context, w io.Writer, payload users.NewUser) error {
	return td.EncoderFunc(ctx, w, payload)
}

// NewUserTargetDecoder implements a decoder for the users.NewUser type
// using a provided function.
type NewUserTargetDecoder struct {
	DecoderFunc func(context.Context,io.Reader) (users.NewUser, error)
}

// Decode implements the decoding function for type users.NewUser used in users by
// calling the underline decoding function to handle the work.
func (td NewUserTargetDecoder) Decode(ctx context.Context, r io.Reader) (users.NewUser, error) {
	return td.DecoderFunc(ctx, r)
}

//****************************************************************************
// users.NewUser Type Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// NewUserTypeEncoder implements a encoder for the users.NewUser type.
type NewUserTypeEncoder struct {
	Encoder Encoder
}

// Encode implements the encode function for type users.NewUser used in users by
// calling the underline Encoder to handle the work.
func (en NewUserTypeEncoder) Encode(ctx context.Context, w io.Writer, payload users.NewUser) error {
	return en.Encoder.Encode(ctx, w, payload)
}


// NewUserTypeDecoder implements a decoder for the users.NewUser type.
type NewUserTypeDecoder struct {
	Decoder TargetDecoder
}

// Decode implements the decode function for type users.NewUser used in users by
// calling the underline Decoder to handle the work.
func (td NewUserTypeDecoder) Decode(ctx context.Context,r io.Reader) (users.NewUser, error) {  
	var res users.NewUser
	err := td.Decoder.Decode(ctx, r, &res)
	return res, err
}


//****************************************************************************
// users.User Target Function Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// UserTargetEncoder implements a encoder for the users.User type
// using a provided function.
type UserTargetEncoder struct {
	EncoderFunc func(context.Context, io.Writer, users.User) error
}

// Encode implements the encoding function for type users.User used in users by
// calling the underline encoding function to handle the work.
func (td UserTargetEncoder) Encode(ctx context.Context, w io.Writer, payload users.User) error {
	return td.EncoderFunc(ctx, w, payload)
}

// UserTargetDecoder implements a decoder for the users.User type
// using a provided function.
type UserTargetDecoder struct {
	DecoderFunc func(context.Context,io.Reader) (users.User, error)
}

// Decode implements the decoding function for type users.User used in users by
// calling the underline decoding function to handle the work.
func (td UserTargetDecoder) Decode(ctx context.Context, r io.Reader) (users.User, error) {
	return td.DecoderFunc(ctx, r)
}

//****************************************************************************
// users.User Type Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
// Used By: users.UserService
//****************************************************************************

// UserTypeEncoder implements a encoder for the users.User type.
type UserTypeEncoder struct {
	Encoder Encoder
}

// Encode implements the encode function for type users.User used in users by
// calling the underline Encoder to handle the work.
func (en UserTypeEncoder) Encode(ctx context.Context, w io.Writer, payload users.User) error {
	return en.Encoder.Encode(ctx, w, payload)
}


// UserTypeDecoder implements a decoder for the users.User type.
type UserTypeDecoder struct {
	Decoder TargetDecoder
}

// Decode implements the decode function for type users.User used in users by
// calling the underline Decoder to handle the work.
func (td UserTypeDecoder) Decode(ctx context.Context,r io.Reader) (users.User, error) {  
	var res users.User
	err := td.Decoder.Decode(ctx, r, &res)
	return res, err
}



//****************************************************************************
// JSON Encoders / Decoders Implementations
// Source: github.com/gokit/rpkit/examples/users
//****************************************************************************

// JSONEncoder implements a wrapper over the encoding/json JSONEncoder to
// match the Encoder interface. This allow us use JSON has a encoder for
// any method encoding as needed.
type JSONEncoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (JSONEncoder) Encode(ctx context.Context,w io.Writer, payload interface{}) error {
	return json.NewEncoder(w).Encode(payload)
}

// JSONDecoder implements a wrapper over the encoding/json JSONEncoder to
// all incoming data into a interface{} type.
type JSONDecoder struct{}

// Encode implements the necessary logic to use json for encoding.
func (JSONDecoder) Decode(ctx context.Context,r io.Reader) (interface{}, error) {
	var data interface{}
	err := json.NewDecoder(r).Decode(&data)
	return data, err
}

// JSONTargetDecoder implements a wrapper over the encoding/json JSONEncoder to
// all incoming data into a interface{} type.
type JSONTargetDecoder struct{}

// Encode implements the necessary logic to use json for encoding to provided
// target of type interface{}.
func (JSONTargetDecoder) Decode(ctx context.Context, r io.Reader, data interface{}) error {
	return json.NewDecoder(r).Decode(data)
}
