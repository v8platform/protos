package rasbinary

import (
	"bytes"
	"fmt"
	"github.com/v8platform/protos/extra/encoding/rasbinary/internal/set"
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

// UnmarshalOptions is a configurable RAS format parser.
type UnmarshalOptions struct {

	// ProtocolVersion uses to decode field.
	ProtocolVersion int32

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
	if err := dec.unmarshalMessage(m.ProtoReflect(), false); err != nil {
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
func (d decoder) unmarshalMessage(m pref.Message, skipTypeURL bool) error {

	if unmarshal := hasUnmarshaler(m); unmarshal != nil {

		n, err := unmarshal.UnmarshalRAS(d, d.opts.ProtocolVersion)
		d.n += n
		return err
	}

	var seenOneofs set.Ints

	fields := getFields(m)

	Each(fields, func(f field) bool {
		fd := f.fd

		if f.GetVersion() > d.opts.ProtocolVersion ||
			f.GetIgnore() || f.GetNoUnmarshal() {
			return true
		}

		switch {

		case fd.IsList():
			list := m.Mutable(fd).List()
			if err := d.unmarshalList(f, list, fd); err != nil {
				d.err = err
				return false
			}
		case fd.IsMap():
			mmap := m.Mutable(fd).Map()
			if err := d.unmarshalMap(f, mmap, fd); err != nil {
				d.err = err
				return false
			}
		default:
			// If field is a oneof, check if it has already been set.
			if od := fd.ContainingOneof(); od != nil {
				idx := uint64(od.Index())
				if seenOneofs.Has(idx) {
					d.err = fmt.Errorf("error decoding %s, oneof %v is already set", f.fd.FullName(), od.FullName())
					return false
				}
			}

			// Required or optional fields.
			if err := d.unmarshalSingular(f, m, fd); err != nil {
				d.err = err
				return false
			}
		}
		return true
	})

	return d.err
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
func (d decoder) unmarshalSingular(f field, m pref.Message, fd pref.FieldDescriptor) error {
	var val pref.Value
	var err error
	switch fd.Kind() {
	case pref.MessageKind, pref.GroupKind:

		if f.TypeField != nil {
			md := fd.Message()
			fd2 := m.Descriptor().Fields().ByNumber(pref.FieldNumber(f.GetTypeField()))

			val2 := m.Get(fd2)
			val := findExtensionByFullname(md, string(fd2.Enum().FullName()))

			if int32(val.Enum()) != int32(val2.Enum()) {
				return nil
			}
		}

		val = m.NewField(fd)
		err = d.unmarshalMessage(val.Message(), false)

	default:
		val, err = d.unmarshalScalar(f, m, fd)
	}

	if err != nil {
		return err
	}
	m.Set(fd, val)
	return nil
}

func findExtensionByFullname(md pref.Descriptor, fullname string) (val pref.Value) {
	md.Options().ProtoReflect().Range(func(descriptor pref.FieldDescriptor, value pref.Value) bool {
		typeFullName := descriptor.Enum().FullName()
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
func (d decoder) unmarshalScalar(f field, m pref.Message, fd pref.FieldDescriptor) (pref.Value, error) {

	kind := fd.Kind()
	decodeFn, _ := GetDecodeFunc(f.GetEncoder())

	switch kind {

	case pref.BoolKind:
		var val bool

		if decodeFn == nil {
			decodeFn = decodeBool
		}

		n, err := decodeFn(d, &val, f.Opts)
		d.n += n
		return pref.ValueOfBool(val), err

	case pref.Int32Kind, pref.Sint32Kind, pref.Sfixed32Kind:
		var val int32

		if decodeFn == nil {
			decodeFn = decodeUint32
		}

		n, err := decodeFn(d, &val, f.Opts)
		d.n += n
		return pref.ValueOfInt32(val), err

	case pref.Int64Kind, pref.Sint64Kind, pref.Sfixed64Kind:
		var val int64

		if decodeFn == nil {
			decodeFn = decodeUint64
		}

		n, err := decodeFn(d, &val, f.Opts)
		d.n += n
		return pref.ValueOfInt64(val), err

	case pref.Uint32Kind, pref.Fixed32Kind:
		var val uint32
		if decodeFn == nil {
			decodeFn = decodeUint32
		}
		n, err := decodeFn(d, &val, f.Opts)

		d.n += n
		return pref.ValueOfUint32(val), err

	case pref.Uint64Kind, pref.Fixed64Kind:
		var val uint64
		if decodeFn == nil {
			decodeFn = decodeUint64
		}
		n, err := decodeFn(d, &val, f.Opts)
		d.n += n
		return pref.ValueOfUint64(val), err

	case pref.FloatKind:
		var val float32

		if decodeFn == nil {
			decodeFn = decodeFloat32
		}

		n, err := decodeFn(d, &val, f.Opts)
		d.n += n
		return pref.ValueOfFloat32(val), err

	case pref.DoubleKind:

		var val float64

		if decodeFn == nil {
			decodeFn = decodeFloat64
		}

		n, err := decodeFn(d, &val, f.Opts)
		d.n += n
		return pref.ValueOfFloat64(val), err

	case pref.StringKind:

		var val string
		if decodeFn == nil {
			decodeFn = decodeString
		}
		n, err := decodeFn(d, &val, f.Opts)
		d.n += n
		return pref.ValueOfString(val), err

	case pref.BytesKind:
		if v, ok := d.unmarshalBytes(f, m, fd); ok {
			return v, nil
		}

	case pref.EnumKind:
		if v, ok := d.unmarshalEnum(f, m, fd); ok {
			return v, nil
		}

	default:
		panic(fmt.Sprintf("unmarshalScalar: invalid scalar kind %v", kind))
	}

	return pref.Value{}, fmt.Errorf("invalid value for <%v> field: %v", kind, f)
}

func (d decoder) unmarshalEnum(f field, m pref.Message, fd pref.FieldDescriptor) (pref.Value, bool) {

	var number int32
	decodeFn, _ := GetDecodeFunc(f.GetEncoder())
	if decodeFn == nil {
		decodeFn = decodeByte
	}
	n, err := decodeFn(d, &number, f.Opts)
	d.n += n
	if err != nil {
		d.err = err
		return pref.Value{}, false
	}

	return pref.ValueOfEnum(pref.EnumNumber(number)), true
}

func (d decoder) decodeValue(fn TypeDecoderFunc, into interface{}, opts ...map[string]string) (int, error) {

	return fn(d, into, opts...)

}

func (d decoder) unmarshalBytes(f field, m pref.Message, fd pref.FieldDescriptor) (pref.Value, bool) {

	var b []byte
	decodeFn, _ := GetDecodeFunc(f.GetEncoder())
	if decodeFn == nil {
		decodeFn = decodeBytes
	}
	if f.SizeField != nil {
		size := getFieldValueOfNumber(m, f.GetSizeField()).(int32)
		b = make([]byte, size)
	}

	n, err := decodeFn(d, b, f.Opts)
	d.n += n
	if err != nil {
		d.err = err
		return pref.Value{}, false
	}

	return pref.ValueOfBytes(b), true
}

func (d decoder) unmarshalMap(f field, mmap pref.Map, fd pref.FieldDescriptor) error {
	return nil
}

func (d decoder) unmarshalList(f field, list pref.List, fd pref.FieldDescriptor) error {
	return nil
}
