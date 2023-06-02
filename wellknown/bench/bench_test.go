package wrapperspb

import (
	"bytes"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

func Benchmark(b *testing.B) {
	input := Double(12345)
	want, _ := proto.Marshal(input)
	var got []byte
	b.Run("std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			got, _ = proto.Marshal(input)
		}
		if !bytes.Equal(want, got) {
			b.Fatalf("bad value: %v vs %v", want, got)
		}
	})
	b.Run("vt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			got, _ = input.MarshalVT()
		}
		if !bytes.Equal(want, got) {
			b.Fatalf("bad value: %v vs %v", want, got)
		}
	})
	b.Run("hand-std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			gotv, _ := input.ProtoMethods().Marshal(protoiface.MarshalInput{
				Message: input.ProtoReflect(),
			})
			got = gotv.Buf
		}
		if !bytes.Equal(want, got) {
			b.Fatalf("bad value: %v vs %v", want, got)
		}
	})
	b.Run("hand-vt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			got, _ = input.FastHandMarshal()
		}
		if !bytes.Equal(want, got) {
			b.Fatalf("bad value: %v vs %v", want, got)
		}
	})
}
