package tests

import (
	"bytes"
	"context"
	"io"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Encoder defines an interface representing a generic encoder which
//expects a interface{} type has it's encoding target.
type Encoder interface {
	Encode(context.Context, io.Writer, interface{}) error
}

//TargetDecoder defines an interface representing a target decoder which
//expects a type to decode into.
type UnmarshalDecoder interface {
	Decode(context.Context, io.Reader, interface{}) error
}

func VerifyCodecsWithMaps(t *testing.T, encoder Encoder, decoder UnmarshalDecoder) {
	testData := map[string]interface{}{
		"name": "Shelly",
	}

	var encoded bytes.Buffer
	encodingError := encoder.Encode(context.Background(), &encoded, testData)
	assert.NoError(t, encodingError)
	assert.NotEmpty(t, encoded)

	var res map[string]interface{}
	err := decoder.Decode(context.Background(), &encoded, &res)
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(testData, res))
}

type human struct {
	Name string
}

func VerifyCodecsWithStructs(t *testing.T, encoder Encoder, decoder UnmarshalDecoder) {
	testData := human{Name: "Shelly"}

	var encoded bytes.Buffer
	encodingError := encoder.Encode(context.Background(), &encoded, testData)
	assert.NoError(t, encodingError)
	assert.NotEmpty(t, encoded)

	var res human
	err := decoder.Decode(context.Background(), &encoded, &res)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	assert.True(t, reflect.DeepEqual(testData, res))
}
