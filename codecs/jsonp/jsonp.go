// Package jsonp implements the Encoder and Decoders for use with the protobufs jsonp format.
package jsonp

//var (
//	// DefaultJSONPEncoder provides a package-level json encoder for use.
//	DefaultJSONPEncoder JSONPEncoder
//
//	// DefaultJSONPDecoder provides a package-level json decoder for use.
//	DefaultJSONPDecoder JSONPDecoder
//
//	// DefaultJSONPTargetDecoder provides a package-level json target decoder for use.
//	DefaultJSONPTargetDecoder JSONPUnmarshalDecoder
//)
//
//// JSONPEncoder implements a wrapper over the encoding/json JSONPEncoder to
//// match the Encoder interface. This allow us use JSONP has a encoder for
//// any method encoding as needed.
//type JSONPEncoder struct{}
//
//// Encode implements the necessary logic to use json for encoding.
//func (JSONPEncoder) Encode(ctx context.Context, w io.Writer, payload interface{}) error {
//	data, err := jsonp.Marshal(payload)
//	if err != nil {
//		return err
//	}
//
//	_, err = w.Write(data)
//	return err
//}
//
//// JSONPDecoder implements a wrapper over the encoding/json JSONPEncoder to
//// all incoming data into a interface{} type.
//type JSONPDecoder struct{}
//
//// Encode implements the necessary logic to use json for encoding.
//func (JSONPDecoder) Decode(ctx context.Context, r io.Reader) (interface{}, error) {
//	var data interface{}
//	err := jsonp.Unmarshal(r).Decode(&data)
//	return data, err
//}
//
//// JSONPUnmarshalDecoder implements a wrapper over the encoding/json JSONPEncoder to
//// all incoming data into a interface{} type.
//type JSONPUnmarshalDecoder struct{}
//
//// Encode implements the necessary logic to use json for encoding to provided
//// target of type interface{}.
//func (JSONPUnmarshalDecoder) Decode(ctx context.Context, r io.Reader, data interface{}) error {
//	return msgpack.NewDecoder(r).Decode(data)
//}
