package sereal_test

import (
	"testing"

	"github.com/gokit/rpkit/codecs/helpers/tests"
	"github.com/gokit/rpkit/codecs/sereal"
)

func TestCodec(t *testing.T) {
	tests.VerifyCodecsWithMaps(t, sereal.DefaultSerealEncoder, sereal.DefaultSerealTargetDecoder)
	tests.VerifyCodecsWithStructs(t, sereal.DefaultSerealEncoder, sereal.DefaultSerealTargetDecoder)
}
