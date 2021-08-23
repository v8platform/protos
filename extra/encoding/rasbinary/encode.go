package rasbinary

import (
	"bytes"
	"fmt"
	"google.golang.org/protobuf/proto"
	pref "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"io"
)

// Marshal writes the given proto.Message in RAS format using default options.
// Do not depend on the output being stable. It may change over time across
// different versions of the program.
func Marshal(m proto.Message) ([]byte, error) {
	return MarshalOptions{}.Marshal(m)
}

// MarshalOptions is a configurable RAS format marshaler.
type MarshalOptions struct {

	// ProtocolVersion uses to encode field.
	ProtocolVersion int32

	// Resolver is used for looking up types when expanding google.protobuf.Any
	// messages. If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver interface {
		protoregistry.ExtensionTypeResolver
		protoregistry.MessageTypeResolver
	}
}

// Marshal marshals the given proto.Message in the RAS format using options in
// MarshalOptions. Do not depend on the output being stable. It may change over
// time across different versions of the program.
func (o MarshalOptions) Marshal(m proto.Message) ([]byte, error) {
	return o.marshal(m)
}

// marshal is a centralized function that all marshal operations go through.
// For profiling purposes, avoid changing the name of this function or
// introducing other code paths for marshal that do not go through this.
func (o MarshalOptions) marshal(m proto.Message) ([]byte, error) {

	if o.Resolver == nil {
		o.Resolver = protoregistry.GlobalTypes
	}

	// Treat nil message interface as an empty message,
	// in which case the output in an empty RAS object.
	if m == nil {
		return []byte{}, nil
	}

	buf := &bytes.Buffer{}

	enc := encoder{buf, o}
	if err := enc.marshalMessage(m.ProtoReflect(), ""); err != nil {
		return nil, err
	}
	return buf.Bytes(), proto.CheckInitialized(m)
}

type encoder struct {
	io.Writer
	opts MarshalOptions
}

// unpopulatedFieldRanger wraps a protoreflect.Message and modifies its Range
// method to additionally iterate over unpopulated fields.
type unpopulatedFieldRanger struct{ pref.Message }

func (m unpopulatedFieldRanger) Range(f func(pref.FieldDescriptor, pref.Value) bool) {
	fds := m.Descriptor().Fields()
	for i := 0; i < fds.Len(); i++ {
		fd := fds.Get(i)
		if m.Has(fd) || fd.ContainingOneof() != nil {
			continue // ignore populated fields and fields within a oneofs
		}

		v := m.Get(fd)
		if !f(fd, v) {
			return
		}
	}
	m.Message.Range(f)
}

type Marshaller interface {
	MarshalRAS(writer io.Writer, version int32) (n int, err error)
}

func hasMarshaller(m pref.Message) Marshaller {

	v := pref.ValueOfMessage(m)
	if u, ok := v.Interface().(Marshaller); ok {
		return u
	}

	return nil
}

// marshalMessage marshals the fields in the given protoreflect.Message.
// If the typeURL is non-empty, then a synthetic "@type" field is injected
// containing the URL as the value.
func (e encoder) marshalMessage(m pref.Message, typeURL string) error {

	if marshaller := hasMarshaller(m); marshaller != nil {
		_, err := marshaller.MarshalRAS(e, e.opts.ProtocolVersion)
		return err
	}

	fields := getFields(m)

	var err error
	Each(fields, func(f field) bool {

		fd := f.fd
		v := f.value

		if f.GetVersion() > e.opts.ProtocolVersion ||
			f.GetIgnore() || f.GetNoMarshal() {
			return true
		}

		switch {
		case fd.IsList():
			if err = e.marshalList(f, v.List(), fd); err != nil {
				return false
			}

		case fd.IsMap():

			if err = e.marshalMap(v.Map(), fd); err != nil {
				return false
			}

		case fd.Kind() == pref.MessageKind, fd.Kind() == pref.GroupKind:

			if f.TypeField != nil {
				md := fd.Message()
				fd2 := m.Descriptor().Fields().ByNumber(pref.FieldNumber(f.GetTypeField()))

				val2 := m.Get(fd2)
				val := findExtensionByFullname(md, string(fd2.Enum().FullName()))

				if int32(val.Enum()) != int32(val2.Enum()) {
					return true
				}
			}

			if err = e.marshalMessage(v.Message(), ""); err != nil {
				return false
			}
		default:
			if err = e.marshalSingular(f, v, fd); err != nil {
				return false
			}
		}

		return true
	})
	return err
}

// // marshalValue marshals the given protoreflect.Value.
// func (e encoder) marshalValue(f field, val pref.Value, fd pref.FieldDescriptor) error {
// 	switch {
// 	case fd.IsList():
// 		return e.marshalList(f, val.List(), fd)
// 	case fd.IsMap():
// 		return e.marshalMap(val.Map(), fd)
// 	default:
// 		return e.marshalSingular(f, val, fd)
// 	}
// }

// marshalSingular marshals the given non-repeated field value. This includes
// all scalar types, enums, messages, and groups.
func (e encoder) marshalSingular(f field, val pref.Value, fd pref.FieldDescriptor) error {

	if !val.IsValid() {
		return nil
	}

	encodeFunc, _ := GetEncodeFunc(f.GetEncoder())

	switch kind := fd.Kind(); kind {
	case pref.BoolKind:

		if encodeFunc == nil {
			encodeFunc = encodeBool
		}

		_, err := encodeFunc(e, val.Bool())
		if err != nil {
			return err
		}

	case pref.StringKind:

		if encodeFunc == nil {
			encodeFunc = encodeString
		}

		_, err := encodeFunc(e, val.String())
		if err != nil {
			return err
		}

	case pref.Int32Kind, pref.Sint32Kind, pref.Sfixed32Kind,
		pref.Uint32Kind, pref.Fixed32Kind:

		if encodeFunc == nil {
			encodeFunc = encodeUint32
		}

		_, err := encodeFunc(e, int32(val.Int()))
		if err != nil {
			return err
		}

	case pref.Int64Kind, pref.Sint64Kind, pref.Uint64Kind,
		pref.Sfixed64Kind, pref.Fixed64Kind:

		if encodeFunc == nil {
			encodeFunc = encodeUint64
		}

		_, err := encodeFunc(e, val.Int())
		if err != nil {
			return err
		}

	case pref.FloatKind:
		// Encoder.WriteFloat handles the special numbers NaN and infinites.
		if encodeFunc == nil {
			encodeFunc = encodeFloat32
		}

		_, err := encodeFunc(e, val.Float())
		if err != nil {
			return err
		}

	case pref.DoubleKind:
		// Encoder.WriteFloat handles the special numbers NaN and infinites.
		if encodeFunc == nil {
			encodeFunc = encodeFloat64
		}

		_, err := encodeFunc(e, val.Float())
		if err != nil {
			return err
		}

	case pref.BytesKind:
		if encodeFunc == nil {
			encodeFunc = encodeBytes
		}

		_, err := encodeFunc(e, val.Bytes())
		if err != nil {
			return err
		}

	case pref.EnumKind:
		if encodeFunc == nil {
			encodeFunc = encodeByte
		}

		_, err := encodeFunc(e, int(val.Enum()))
		if err != nil {
			return err
		}

	default:
		panic(fmt.Sprintf("%v has unknown kind: %v", fd.FullName(), kind))
	}

	return nil
}

// marshalList marshals the given protoreflect.List.
func (e encoder) marshalList(f field, list pref.List, fd pref.FieldDescriptor) error {

	_, err := encodeSize(e, list.Len())
	if err != nil {
		return err
	}
	for i := 0; i < list.Len(); i++ {
		item := list.Get(i)
		if err := e.marshalSingular(f, item, fd); err != nil {
			return err
		}
	}
	return nil
}

// marshalMap marshals given protoreflect.Map.
func (e encoder) marshalMap(mmap pref.Map, fd pref.FieldDescriptor) error {

	_, err := encodeSize(e, mmap.Len())
	if err != nil {
		return err
	}

	// panic("TODO marshalMap")

	// e.StartObject()
	// defer e.EndObject()
	//
	// var err error
	// order.RangeEntries(mmap, order.GenericKeyOrder, func(k pref.MapKey, v pref.Value) bool {
	// 	if err = e.WriteName(k.String()); err != nil {
	// 		return false
	// 	}
	// 	if err = e.marshalSingular(v, fd.MapValue()); err != nil {
	// 		return false
	// 	}
	// 	return true
	// })
	// return err

	return nil
}
