package ras

import (
	extpb "github.com/v8platform/protos/gen/ras/encoding"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func getFields(m protoreflect.Message) []field {
	var fields []field

	md := m.Descriptor()

	mFields := md.Fields()

	for i := 0; i < mFields.Len(); i++ {
		fd := mFields.Get(i)

		encoderOptions := proto.GetExtension(fd.Options(), extpb.E_Field).(*extpb.EncodingFieldOptions)
		if encoderOptions == nil {
			continue
		}
		fields = append(fields, field{
			encoderOptions,
			fd,
			m.Get(fd),
		})
	}

	return fields
}

func Each(list []field, fn func(f field) bool) {

	for _, f := range list {
		if !fn(f) {
			break
		}
	}
}

type field struct {
	*extpb.EncodingFieldOptions
	fd    protoreflect.FieldDescriptor
	value protoreflect.Value
}
