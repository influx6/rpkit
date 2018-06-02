package gobs_test

import (
	"testing"

	"github.com/gokit/rpkit/codecs/gobs"
	"github.com/gokit/rpkit/codecs/helpers/tests"
)

func TestCodec(t *testing.T) {
	tests.VerifyCodecsWithMaps(t, gobs.DefaultGOBSEncoder, gobs.DefaultGOBSTargetDecoder)
	tests.VerifyCodecsWithStructs(t, gobs.DefaultGOBSEncoder, gobs.DefaultGOBSTargetDecoder)
}
