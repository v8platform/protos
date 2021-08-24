package rasbinary

import (
	extpb "github.com/v8platform/protos/gen/ras/encoding"
	"google.golang.org/protobuf/reflect/protoreflect"
	"sort"
)

func RangeFields(m protoreflect.Message,
	serviceVersion int32,
	fn func(fd protoreflect.FieldDescriptor, value protoreflect.Value, opts *extpb.EncodingFieldOptions) bool) {

	md := m.Descriptor()

	mFields := md.Fields()
	var fields []field

	for i := 0; i < mFields.Len(); i++ {
		fd := mFields.Get(i)

		if fd.IsExtension() {
			continue
		}

		encoderOptions, ok := GetEncodingFieldOptions(fd)
		if !ok ||
			encoderOptions.GetOrder() == 0 ||
			encoderOptions.GetVersion() > serviceVersion {
			continue
		}

		fields = append(fields, field{
			opts:  encoderOptions,
			fd:    fd,
			value: m.Get(fd),
		})
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].opts.GetOrder() < fields[j].opts.GetOrder()
	})

	for _, f := range fields {
		if ok := fn(f.fd, f.value, f.opts); !ok {
			break
		}
	}
}

type field struct {
	n     protoreflect.FieldNumber
	opts  *extpb.EncodingFieldOptions
	fd    protoreflect.FieldDescriptor
	value protoreflect.Value
}
