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
// RP: Output with Returning Error Decoders
// Method: Raise
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Raise
//****************************************************************************

// RaiseDecoder implements the RaiseEncoder by providing a wrapper around
// a implementation of the Encoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type RaiseEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by RaiseEncoder.
func (ew RaiseEncoderWrapper) Encode(w io.Writer, payload string) error {
	return ew.Encoder.Encode(w, payload)
}

// RaiseClientDecoderWrapper implements the RaiseClientDecoder by providing a wrapper around
// a implementation of the ClientDecoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type RaiseClientDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by RaiseClientDecoder.
func (cdw RaiseClientDecoderWrapper) Decode(r io.Reader) (string, error) {
	recs, err := cdw.Decoder.Decode(r)
	if err != nil {
		return "", err
	}

	if wantedRecs, ok := recs.(string); ok {
		return wantedRecs, nil
	}

	return "", ErrDecodedUnknownType
}

// RaiseClientJSONDecoder implements a wrapper over the encoding/json JSONDecoder to
// match the Decoder interface. This allow us use JSON has a decoding for any method decoding as
// needed. This specifically decodes incoming data into string type.
type RaiseClientJSONDecoder struct{}

// Decode implements the necessary logic to use json for decoding.
func (RaiseClientJSONDecoder) Decode(r io.Reader) (interface{}, error) {
	var res string
	err := json.NewDecoder(r).Decode(&res)
	return res, err
}

//****************************************************************************
// RP: Input Returning Only Error Encoders
// Method: Delete
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Delete
//****************************************************************************
// DeleteDecoder implements the DeleteDecoder by providing a wrapper around
// a implementation of the Decoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type DeleteDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by DeleteDecoder.
func (dw DeleteDecoderWrapper) Decode(r io.Reader) (string, error) {
	recs, err := dw.Decoder.Decode(r)
	if err != nil {
		return "", err
	}

	if wantedRecs, ok := recs.(string); ok {
		return wantedRecs, nil
	}

	return "", ErrDecodedUnknownType
}

// DeleteJSONDecoder implements a wrapper over the encoding/json JSONDecoder to
// match the Decoder interface. This allow us use JSON has a decoding for any method decoding as
// needed. This specifically decodes incoming data into wanted type.
type DeleteJSONDecoder struct{}

// Decode implements the necessary logic to use json for decoding.
func (DeleteJSONDecoder) Decode(r io.Reader) (interface{}, error) {
	var res string
	err := json.NewDecoder(r).Decode(&res)
	return res, err
}

// DeleteClientEncoderWrapper implements the DeleteClientEncoder by providing a wrapper around
// a implementation of the ClientEncoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type DeleteClientEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by DeleteClientEncoder.
func (cew DeleteClientEncoderWrapper) Encode(w io.Writer, payload string) error {
	return cew.Encoder.Encode(w, payload)
}

//****************************************************************************
// RP: Input Returning Only Error Encoders
// Method: Update
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Update
//****************************************************************************
// UpdateDecoder implements the UpdateDecoder by providing a wrapper around
// a implementation of the Decoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type UpdateDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by UpdateDecoder.
func (dw UpdateDecoderWrapper) Decode(r io.Reader) (users.UpdateUser, error) {
	recs, err := dw.Decoder.Decode(r)
	if err != nil {
		return users.UpdateUser{}, err
	}

	if wantedRecs, ok := recs.(users.UpdateUser); ok {
		return wantedRecs, nil
	}

	return users.UpdateUser{}, ErrDecodedUnknownType
}

// UpdateJSONDecoder implements a wrapper over the encoding/json JSONDecoder to
// match the Decoder interface. This allow us use JSON has a decoding for any method decoding as
// needed. This specifically decodes incoming data into wanted type.
type UpdateJSONDecoder struct{}

// Decode implements the necessary logic to use json for decoding.
func (UpdateJSONDecoder) Decode(r io.Reader) (interface{}, error) {
	var res users.UpdateUser
	err := json.NewDecoder(r).Decode(&res)
	return res, err
}

// UpdateClientEncoderWrapper implements the UpdateClientEncoder by providing a wrapper around
// a implementation of the ClientEncoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type UpdateClientEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by UpdateClientEncoder.
func (cew UpdateClientEncoderWrapper) Encode(w io.Writer, payload users.UpdateUser) error {
	return cew.Encoder.Encode(w, payload)
}

//****************************************************************************
// RP: Input And Output Returning Error Encoders/Decoders
// Method: Get
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.Get
//****************************************************************************

// GetDecoder implements the GetDecoder by providing a wrapper around
// a implementation of the Decoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type GetDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by GetDecoder.
func (dw GetDecoderWrapper) Decode(r io.Reader) (string, error) {
	recs, err := dw.Decoder.Decode(r)
	if err != nil {
		return "", err
	}

	if wantedRecs, ok := recs.(string); ok {
		return wantedRecs, nil
	}

	return "", ErrDecodedUnknownType
}

// GetJSONDecoder implements a wrapper over the encoding/json JSONDecoder to
// match the Decoder interface. This allow us use JSON has a decoding for any method decoding as
// needed. This specifically decodes incoming data into wanted type.
type GetJSONDecoder struct{}

// Decode implements the necessary logic to use json for decoding.
func (GetJSONDecoder) Decode(r io.Reader) (interface{}, error) {
	var res string
	err := json.NewDecoder(r).Decode(&res)
	return res, err
}

// GetClientDecoderWrapper implements the GetClientDecoder by providing a wrapper around
// a implementation of the ClientDecoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type GetClientDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by GetClientDecoder.
func (cdw GetClientDecoderWrapper) Decode(r io.Reader) (users.User, error) {
	recs, err := cdw.Decoder.Decode(r)
	if err != nil {
		return users.User{}, err
	}

	if wantedRecs, ok := recs.(users.User); ok {
		return wantedRecs, nil
	}

	return users.User{}, ErrDecodedUnknownType
}

// GetDecoder implements the GetEncoder by providing a wrapper around
// a implementation of the Encoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type GetEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by GetEncoder.
func (ew GetEncoderWrapper) Encode(w io.Writer, payload users.User) error {
	return ew.Encoder.Encode(w, payload)
}

// GetClientEncoderWrapper implements the GetClientEncoder by providing a wrapper around
// a implementation of the ClientEncoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type GetClientEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by GetClientEncoder.
func (cew GetClientEncoderWrapper) Encode(w io.Writer, payload string) error {
	return cew.Encoder.Encode(w, payload)
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

//****************************************************************************
// RP: Input And Output Returning Error Encoders/Decoders
// Method: GetAll
// Source: github.com/gokit/rpkit/examples/users
// Handler: users.UserService.GetAll
//****************************************************************************

// GetAllDecoder implements the GetAllDecoder by providing a wrapper around
// a implementation of the Decoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type GetAllDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by GetAllDecoder.
func (dw GetAllDecoderWrapper) Decode(r io.Reader) (string, error) {
	recs, err := dw.Decoder.Decode(r)
	if err != nil {
		return "", err
	}

	if wantedRecs, ok := recs.(string); ok {
		return wantedRecs, nil
	}

	return "", ErrDecodedUnknownType
}

// GetAllJSONDecoder implements a wrapper over the encoding/json JSONDecoder to
// match the Decoder interface. This allow us use JSON has a decoding for any method decoding as
// needed. This specifically decodes incoming data into wanted type.
type GetAllJSONDecoder struct{}

// Decode implements the necessary logic to use json for decoding.
func (GetAllJSONDecoder) Decode(r io.Reader) (interface{}, error) {
	var res string
	err := json.NewDecoder(r).Decode(&res)
	return res, err
}

// GetAllClientDecoderWrapper implements the GetAllClientDecoder by providing a wrapper around
// a implementation of the ClientDecoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Decoder which behaves in such a manner.
type GetAllClientDecoderWrapper struct {
	Decoder Decoder
}

// Decode implements the Decode function required by GetAllClientDecoder.
func (cdw GetAllClientDecoderWrapper) Decode(r io.Reader) ([]users.User, error) {
	recs, err := cdw.Decoder.Decode(r)
	if err != nil {
		return nil, err
	}

	if wantedRecs, ok := recs.([]users.User); ok {
		return wantedRecs, nil
	}

	return nil, ErrDecodedUnknownType
}

// GetAllDecoder implements the GetAllEncoder by providing a wrapper around
// a implementation of the Encoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type GetAllEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by GetAllEncoder.
func (ew GetAllEncoderWrapper) Encode(w io.Writer, payload []users.User) error {
	return ew.Encoder.Encode(w, payload)
}

// GetAllClientEncoderWrapper implements the GetAllClientEncoder by providing a wrapper around
// a implementation of the ClientEncoder interface which returns a `interface{}` type. This allows us
// wrap around things like the JSON.Encoder which behaves in such a manner.
type GetAllClientEncoderWrapper struct {
	Encoder Encoder
}

// Encode implements the Encode function required by GetAllClientEncoder.
func (cew GetAllClientEncoderWrapper) Encode(w io.Writer, payload string) error {
	return cew.Encoder.Encode(w, payload)
}
