package rasbinary

import (
	"bytes"
	"fmt"
	"github.com/v8platform/protos/extra/encoding/rasbinary/internal/set"
	extpb "github.com/v8platform/protos/gen/ras/encoding"
	"google.golang.org/protobuf/proto"
	pref "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"io"
)

// Unmarshal reads the given []byte into the given proto.Message.
// The provided message must be mutable (e.g., a non-nil pointer to a message).
func Unmarshal(b []byte, m proto.Message) error {
	return UnmarshalOptions{}.Unmarshal(b, m)
}

// Unmarshal reads the given []byte into the given proto.Message.
// The provided message must be mutable (e.g., a non-nil pointer to a message).
func UnmarshalReader(r io.Reader, m proto.Message) error {
	return UnmarshalOptions{}.UnmarshalReader(r, m)
}

// UnmarshalOptions is a configurable RAS format parser.
type UnmarshalOptions struct {

	// ServiceVersion uses to decode field.
	ServiceVersion int32

	// Resolver is used for looking up types when unmarshaling
	// google.protobuf.Any messages or extension fields.
	// If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver interface {
		protoregistry.MessageTypeResolver
		protoregistry.ExtensionTypeResolver
	}
}

// Unmarshal reads the given []byte and populates the given proto.Message
// using options in the UnmarshalOptions object.
// It will clear the message first before setting the fields.
// If it returns an error, the given message may be partially set.
// The provided message must be mutable (e.g., a non-nil pointer to a message).
func (o UnmarshalOptions) Unmarshal(b []byte, m proto.Message) error {
	return o.unmarshal(b, m)
}

// UnmarshalReader reads the given []byte and populates the given proto.Message
// using options in the UnmarshalOptions object.
// It will clear the message first before setting the fields.
// If it returns an error, the given message may be partially set.
// The provided message must be mutable (e.g., a non-nil pointer to a message).
func (o UnmarshalOptions) UnmarshalReader(r io.Reader, m proto.Message) error {
	return o.unmarshalReader(r, m)
}

// unmarshal is a centralized function that all unmarshal operations go through.
// For profiling purposes, avoid changing the name of this function or
// introducing other code paths for unmarshal that do not go through this.
func (o UnmarshalOptions) unmarshal(b []byte, m proto.Message) error {
	return o.UnmarshalReader(bytes.NewBuffer(b), m)
}

// unmarshal is a centralized function that all unmarshal operations go through.
// For profiling purposes, avoid changing the name of this function or
// introducing other code paths for unmarshal that do not go through this.
func (o UnmarshalOptions) unmarshalReader(r io.Reader, m proto.Message) error {
	proto.Reset(m)

	if o.Resolver == nil {
		o.Resolver = protoregistry.GlobalTypes
	}

	dec := decoder{Reader: r, opts: o}
	if err := dec.unmarshalMessage(m.ProtoReflect(), nil); err != nil {
		return err
	}

	return proto.CheckInitialized(m)
}

type decoder struct {
	io.Reader
	opts UnmarshalOptions
	err  error
	n    int
}

// unmarshalMessage unmarshals a message into the given protoreflect.Message.
func (d decoder) unmarshalMessage(m pref.Message, opts *extpb.EncodingFieldOptions) error {

	if unmarshal := wellKnownTypeUnmarshaler(m.Descriptor().FullName()); unmarshal != nil {
		return unmarshal(d, m, opts)
	}

	if unmarshal := hasUnmarshaler(m); unmarshal != nil {
		_, err := unmarshal.UnmarshalRAS(d, d.opts.ServiceVersion)
		return err
	}

	var seenOneofs set.Ints
	var err error

	RangeFields(m, d.opts.ServiceVersion, func(fd pref.FieldDescriptor, value pref.Value, opts *extpb.EncodingFieldOptions) bool {

		switch {

		case fd.IsList():
			list := m.Mutable(fd).List()
			if err = d.unmarshalList(m, list, fd, opts); err != nil {
				return false
			}
		case fd.IsMap():
			mmap := m.Mutable(fd).Map()
			if err = d.unmarshalMap(m, mmap, fd, opts); err != nil {
				return false
			}
		default:
			// If field is a oneof, check if it has already been set.
			if od := fd.ContainingOneof(); od != nil {
				idx := uint64(od.Index())
				if seenOneofs.Has(idx) {
					err = fmt.Errorf("error decoding %s, oneof %v is already set", fd.FullName(), od.FullName())
					return false
				}
			}

			// Required or optional fields.
			if err = d.unmarshalSingular(m, fd, opts); err != nil {
				return false
			}
		}
		return true
	})

	if err != nil {
		return err
	}

	return nil
}

type Unmarshaler interface {
	UnmarshalRAS(reader io.Reader, version int32) (n int, err error)
}

func hasUnmarshaler(m pref.Message) Unmarshaler {

	v := pref.ValueOfMessage(m)
	if u, ok := v.Interface().(Unmarshaler); ok {
		return u
	}

	return nil
}

// unmarshalSingular unmarshals to the non-repeated field specified
// by the given FieldDescriptor.
func (d decoder) unmarshalSingular(m pref.Message, fd pref.FieldDescriptor, opts *extpb.EncodingFieldOptions) error {
	var val pref.Value
	var err error
	switch fd.Kind() {
	case pref.MessageKind, pref.GroupKind:

		if opts.TypeField != nil {
			md := fd.Message()
			fd2 := m.Descriptor().Fields().ByNumber(pref.FieldNumber(opts.GetTypeField()))

			val2 := m.Get(fd2)
			val := findExtensionByFullname(md, string(fd2.Enum().FullName()))

			if int32(val.Enum()) != int32(val2.Enum()) {
				return nil
			}
		}

		val = m.NewField(fd)

		err = d.unmarshalMessage(val.Message(), opts)

	default:
		val, err = d.unmarshalScalar(m, fd, opts)
	}

	if err != nil {
		return err
	}
	m.Set(fd, val)
	return nil
}

func findExtensionByFullname(md pref.Descriptor, fullname string) (val pref.Value) {
	md.Options().ProtoReflect().Range(func(fd pref.FieldDescriptor, value pref.Value) bool {

		if !fd.IsExtension() {
			return true
		}

		typeFullName := fd.Enum().FullName()
		if string(typeFullName) == fullname {

			val = value
			return false
		}

		return true
	})
	return
}

// unmarshalScalar unmarshals to a scalar/enum protoreflect.Value specified by
// the given FieldDescriptor.
func (d decoder) unmarshalScalar(m pref.Message, fd pref.FieldDescriptor, opts *extpb.EncodingFieldOptions) (pref.Value, error) {

	kind := fd.Kind()
	decodeFn, _ := GetDecodeFunc(opts.GetEncoder())

	switch kind {

	case pref.BoolKind:
		var val bool

		if decodeFn == nil {
			decodeFn = decodeBool
		}

		n, err := decodeFn(d, &val, opts.Opts)
		d.n += n
		return pref.ValueOfBool(val), err

	case pref.Int32Kind, pref.Sint32Kind, pref.Sfixed32Kind:
		var val int32

		if decodeFn == nil {
			decodeFn = decodeUint32
		}

		n, err := decodeFn(d, &val, opts.Opts)
		d.n += n
		return pref.ValueOfInt32(val), err

	case pref.Int64Kind, pref.Sint64Kind, pref.Sfixed64Kind:
		var val int64

		if decodeFn == nil {
			decodeFn = decodeUint64
		}

		n, err := decodeFn(d, &val, opts.Opts)
		d.n += n
		return pref.ValueOfInt64(val), err

	case pref.Uint32Kind, pref.Fixed32Kind:
		var val uint32
		if decodeFn == nil {
			decodeFn = decodeUint32
		}
		n, err := decodeFn(d, &val, opts.Opts)

		d.n += n
		return pref.ValueOfUint32(val), err

	case pref.Uint64Kind, pref.Fixed64Kind:
		var val uint64
		if decodeFn == nil {
			decodeFn = decodeUint64
		}
		n, err := decodeFn(d, &val, opts.Opts)
		d.n += n
		return pref.ValueOfUint64(val), err

	case pref.FloatKind:
		var val float32

		if decodeFn == nil {
			decodeFn = decodeFloat32
		}

		n, err := decodeFn(d, &val, opts.Opts)
		d.n += n
		return pref.ValueOfFloat32(val), err

	case pref.DoubleKind:

		var val float64

		if decodeFn == nil {
			decodeFn = decodeFloat64
		}

		n, err := decodeFn(d, &val, opts.Opts)
		d.n += n
		return pref.ValueOfFloat64(val), err

	case pref.StringKind:

		var val string
		if decodeFn == nil {
			decodeFn = decodeString
		}
		n, err := decodeFn(d, &val, opts.Opts)
		d.n += n
		return pref.ValueOfString(val), err

	case pref.BytesKind:

		return d.unmarshalBytes(m, fd, opts)

	case pref.EnumKind:
		return d.unmarshalEnum(m, fd, opts)

	default:
		panic(fmt.Errorf("invalid value for <%v> field: %v", kind, opts))
	}

}

func (d decoder) unmarshalEnum(m pref.Message, fd pref.FieldDescriptor, opts *extpb.EncodingFieldOptions) (pref.Value, error) {

	var number int32
	decodeFn, _ := GetDecodeFunc(opts.GetEncoder())
	if decodeFn == nil {
		decodeFn = decodeByte
	}
	_, err := decodeFn(d, &number, opts.Opts)
	if err != nil {
		return pref.Value{}, err
	}

	return pref.ValueOfEnum(pref.EnumNumber(number)), nil
}

func (d decoder) unmarshalBytes(m pref.Message, fd pref.FieldDescriptor, opts *extpb.EncodingFieldOptions) (pref.Value, error) {

	b := make([]byte, 0, 512)

	decodeFn, _ := GetDecodeFunc(opts.GetEncoder())

	if opts.SizeField != nil {
		size := getFieldValueOfNumber(m, opts.GetSizeField()).(int32)
		b = make([]byte, size)
	}

	if decodeFn == nil && opts.SizeField != nil {
		decodeFn = decodeBytes
	}

	if decodeFn == nil {
		all, err := io.ReadAll(d)
		if err != nil {
			return pref.Value{}, err
		}
		return pref.ValueOfBytes(all), nil
	}

	n, err := decodeFn(d, b, opts.Opts)
	d.n += n
	if err != nil {
		return pref.Value{}, err
	}

	return pref.ValueOfBytes(b), nil
}

func (d decoder) unmarshalMap(m pref.Message, mmap pref.Map, fd pref.FieldDescriptor, opts *extpb.EncodingFieldOptions) error {
	return nil
}

func (d decoder) unmarshalList(m pref.Message, list pref.List, fd pref.FieldDescriptor, opts *extpb.EncodingFieldOptions) error {

	var size int
	_, err := decodeSize(d, &size)
	if err != nil {
		return err
	}

	switch fd.Kind() {
	case pref.MessageKind, pref.GroupKind:
		for i := 0; i < size; i++ {

			val := list.NewElement()
			if err := d.unmarshalMessage(val.Message(), opts); err != nil {
				fmt.Println(err.Error())
				return err
			}
			list.Append(val)
		}
	default:
		for i := 0; i < size; i++ {

			val, err := d.unmarshalScalar(m, fd, opts)
			if err != nil {
				return err
			}
			list.Append(val)
		}
	}

	return nil
}
