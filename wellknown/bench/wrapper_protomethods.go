package wrapperspb

import (
	"encoding/binary"
	"math"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

var methods = Marshaller(func(ms protoreflect.ProtoMessage) ([]byte, error) {
	return ms.(*DoubleValue).FastHandMarshal()
})

func (m *DoubleValue) FastHandMarshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	size := 9
	d := make([]byte, size)
	d[0] = 0x9
	binary.LittleEndian.PutUint64(d[1:], math.Float64bits(m.Value))
	return d, nil
}
func (_ *DoubleValue) ProtoMethods() *protoiface.Methods {
	return methods
}

func Marshaller(marshal func(message protoreflect.ProtoMessage) ([]byte, error)) *protoiface.Methods {
	return &protoiface.Methods{
		Size: nil,
		Marshal: func(in protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			out, err := marshal(in.Message.Interface())
			if in.Buf != nil {
				out = append(in.Buf, out...)
			}
			return protoiface.MarshalOutput{
				Buf: out,
			}, err
		},
	}
}
