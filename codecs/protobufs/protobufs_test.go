package protobufs_test

import (
	"reflect"
	"testing"

	"bytes"
	"context"

	"github.com/gokit/rpkit/codecs/helpers/tests"
	"github.com/gokit/rpkit/codecs/protobufs"
	"github.com/gokit/rpkit/codecs/protobufs/rails"
	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	tests.VerifyCodecsWithMaps(t, protobufs.DefaultProtobufEncoder, protobufs.DefaultProtobufTargetDecoder)
	tests.VerifyCodecsWithStructs(t, protobufs.DefaultProtobufEncoder, protobufs.DefaultProtobufTargetDecoder)
}

func TestCodecWithProtoMessage(t *testing.T) {
	var r rails.Rails
	r.Route = 3434
	r.Name = "ishuku"

	var encoded bytes.Buffer
	encodingError := protobufs.DefaultProtobufEncoder.Encode(context.Background(), &encoded, &r)
	assert.NoError(t, encodingError)
	assert.NotEmpty(t, encoded)

	var rs rails.Rails
	err := protobufs.DefaultProtobufTargetDecoder.Decode(context.Background(), &encoded, &rs)
	assert.NoError(t, err)
	assert.NotNil(t, rs)

	assert.True(t, reflect.DeepEqual(r, rs))
}
