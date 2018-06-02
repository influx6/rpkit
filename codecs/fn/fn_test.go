package fn_test

import (
	"reflect"
	"testing"

	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/gokit/rpkit/codecs/fn"
	"github.com/stretchr/testify/assert"
)

func TestFnCodec(t *testing.T) {
	encoder := fn.FuncEncoder{
		EncodeFunc: func(ctx context.Context, w io.Writer, i interface{}) error {
			return json.NewEncoder(w).Encode(i)
		},
	}

	decoder := fn.FuncDecoder{
		DecodeFunc: func(ctx context.Context, r io.Reader) (interface{}, error) {
			var res map[string]interface{}
			err := json.NewDecoder(r).Decode(&res)
			return res, err
		},
	}

	testData := map[string]interface{}{
		"name": "Shelly",
	}

	var encoded bytes.Buffer
	encodingError := encoder.Encode(context.Background(), &encoded, testData)
	assert.NoError(t, encodingError)
	assert.NotEmpty(t, encoded)

	res, err := decoder.Decode(context.Background(), &encoded)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	resMap, ok := res.(map[string]interface{})
	assert.True(t, ok)
	assert.NotNil(t, resMap)

	assert.True(t, reflect.DeepEqual(testData, resMap))
}
