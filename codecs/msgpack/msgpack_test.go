package msgpack_test

import (
	"testing"

	"github.com/gokit/rpkit/codecs/helpers/tests"
	"github.com/gokit/rpkit/codecs/msgpack"
)

func TestCodec(t *testing.T) {
	tests.VerifyCodecsWithMaps(t, msgpack.DefaultMsgPackEncoder, msgpack.DefaultMsgPackTargetDecoder)
	tests.VerifyCodecsWithStructs(t, msgpack.DefaultMsgPackEncoder, msgpack.DefaultMsgPackTargetDecoder)
}
