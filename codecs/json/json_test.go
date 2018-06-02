package json_test

import (
	"testing"

	"github.com/gokit/rpkit/codecs/helpers/tests"
	"github.com/gokit/rpkit/codecs/json"
)

func TestCodec(t *testing.T) {
	tests.VerifyCodecsWithMaps(t, json.DefaultJSONEncoder, json.DefaultJSONTargetDecoder)
	tests.VerifyCodecsWithStructs(t, json.DefaultJSONEncoder, json.DefaultJSONTargetDecoder)
}
