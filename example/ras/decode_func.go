package ras

import (
	"encoding/binary"
	"fmt"
	"github.com/k0kubun/pp"
	uuid "github.com/satori/go.uuid"

	pb "google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"math"
	"reflect"
	"strings"
	"time"
)

var decoderFunc = map[string]TypeDecoderFunc{}

type TypeDecoderFunc func(r io.Reader, into interface{}, opts ...map[string]string) (int, error)

func init() {
	RegisterDecoderType("time", decodeTime)
	RegisterDecoderType("type", decodeType)
	RegisterDecoderType("bool", decodeBool)
	RegisterDecoderType("byte int8 uint8", decodeByte)
	RegisterDecoderType("char short int16 uint16", decodeUint16)
	RegisterDecoderType("int int32 uint32", decodeUint32)
	RegisterDecoderType("int64 uint64 long", decodeUint64)
	RegisterDecoderType("float32", decodeFloat32)
	RegisterDecoderType("float64 double", decodeFloat64)
	RegisterDecoderType("string", decodeString)
	RegisterDecoderType("null-size", decodeNullableSize)
	RegisterDecoderType("size", decodeSize)
	RegisterDecoderType("bytes", decodeBytes)
	RegisterDecoderType("uuid", decodeUUID)
}

func RegisterDecoderType(name string, dec TypeDecoderFunc) {

	names := strings.Fields(strings.ToLower(name))

	for _, s := range names {
		decoderFunc[s] = dec
	}
}

func GetDecodeFunc(name string) (TypeDecoderFunc, bool) {
	fn, ok := decoderFunc[name]
	return fn, ok
}

func DecodeValue(codec string, r io.Reader, into interface{}) (int, error) {

	typeEncoderFunc, ok := decoderFunc[codec]
	if !ok {
		return 0, fmt.Errorf("unknown encoder <%s>", codec)
	}

	return typeEncoderFunc(r, into)
}

func decodeBytes(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {
	//chunkSize := 4096

	//options := map[string]string{}
	//if len(opts) > 0 {
	//	options = opts[0]
	//}

	data, ok := into.([]byte)

	if !ok {
		return 0, &TypeDecodeError{"bytes",
			fmt.Sprintf("convert from <%s> unsupporsed", reflect.TypeOf(into))}
	}

	size := len(data)

	readLength := 0
	n := 0
	var err error

	pp.Println("size", size)
	for readLength < len(data) {
		n, err = r.Read(data[readLength:])
		readLength += n

		if err != nil {
			return 0, err
		}
	}

	return readLength, nil
}

func decodeUUID(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {

	buf := make([]byte, 16)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"uuid",
			err.Error(),
		}
	}

	u, err := uuid.FromBytes(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"uuid",
			err.Error(),
		}
	}

	switch typed := into.(type) {
	case []byte:
		copy(typed, buf)
	case *[]byte:
		*typed = buf
	case *string:
		*typed = u.String()
	case *uuid.UUID:
		*typed = u
	default:
		return n, &TypeDecodeError{"uuid",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}

	return n, nil
}

func decodeTime(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {

	buf := make([]byte, 8)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"time",
			err.Error(),
		}
	}

	val := binary.BigEndian.Uint64(buf)
	ticks := int64(val)
	timeT := (ticks - AgeDelta) / 10

	timestamp := time.Unix(0, timeT*int64(time.Millisecond)).UnixNano()

	switch typed := into.(type) {
	case *uint64:
		*typed = uint64(timestamp)
	case *int64:
		*typed = timestamp
	case *time.Time:
		*typed = time.Unix(0, timestamp)
	case *pb.Timestamp:
		*typed = *pb.New(time.Unix(0, timestamp))
	default:
		return n, &TypeDecodeError{"time",
			fmt.Sprintf("decode time to <%s> unsupporsed", typed)}
	}
	return n, nil
}

func decodeType(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {

	buf := make([]byte, 1)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"type",
			err.Error(),
		}
	}

	b1 := buf[0]
	cur := b1 & 0xFF

	switch typed := into.(type) {
	case *byte:
		*typed = cur
	case *int32:
		*typed = int32(cur)
	case *int:
		*typed = int(cur)
	case *uint32:
		*typed = uint32(cur)
	case *uint:
		*typed = uint(cur)
	case *int64:
		*typed = int64(cur)
	case *uint64:
		*typed = uint64(cur)
	default:
		return n, &TypeDecodeError{"type",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}
	return n, nil
}

func decodeByte(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {
	buf := make([]byte, 1)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"byte",
			err.Error(),
		}
	}

	b1 := buf[0]

	switch typed := into.(type) {
	case *byte:
		*typed = b1
	case *int8:
		*typed = int8(b1)
	case *int32:
		*typed = int32(b1)
	case *int:
		*typed = int(b1)
	case *uint32:
		*typed = uint32(b1)
	case *uint:
		*typed = uint(b1)
	case *int64:
		*typed = int64(b1)
	case *uint64:
		*typed = uint64(b1)
	default:
		return n, &TypeDecodeError{"byte",
			fmt.Sprintf("decode byte to <%s> unsupporsed", typed)}
	}
	return n, nil
}

func decodeBool(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {
	buf := make([]byte, 1)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"bool",
			err.Error(),
		}
	}

	b1 := buf[0]

	var val bool

	switch b1 {
	case TRUE_BYTE:
		val = true
	case FALSE_BYTE:
		val = false
	}

	switch typed := into.(type) {
	case *bool:
		*typed = val
	case *int:
		if val {
			*typed = 1
		} else {
			*typed = 0
		}
	default:
		return n, &TypeDecodeError{"bool",
			fmt.Sprintf("decode byte to <%s> unsupporsed", typed)}
	}
	return n, nil

}

func decodeUint16(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {

	buf := make([]byte, 2)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"uint16",
			err.Error(),
		}
	}

	val := binary.BigEndian.Uint16(buf)

	switch typed := into.(type) {
	case *int:
		*typed = int(val)
	case *uint16:
		*typed = val
	case *int16:
		*typed = int16(val)
	case *uint32:
		*typed = uint32(val)
	case *int32:
		*typed = int32(val)
	case *uint64:
		*typed = uint64(val)
	case *int64:
		*typed = int64(val)
	default:
		return n, &TypeDecodeError{"uint16",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}
	return n, nil

}

func decodeUint32(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {
	buf := make([]byte, 4)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"uint32",
			err.Error(),
		}
	}

	val := binary.BigEndian.Uint32(buf)

	switch typed := into.(type) {
	case *int:
		*typed = int(val)
	case *uint16:
		*typed = uint16(val)
	case *int16:
		*typed = int16(val)
	case *uint32:
		*typed = uint32(val)
	case *int32:
		*typed = int32(val)
	case *uint64:
		*typed = uint64(val)
	case *int64:
		*typed = int64(val)
	default:
		return n, &TypeDecodeError{"uint32",
			fmt.Sprintf("convert to <%s> unsupporsed", reflect.TypeOf(typed))}
	}
	return n, nil

}

func decodeUint64(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {
	buf := make([]byte, 8)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"uint64",
			err.Error(),
		}
	}

	val := binary.BigEndian.Uint64(buf)

	switch typed := into.(type) {
	case *int:
		*typed = int(val)
	case *uint16:
		*typed = uint16(val)
	case *int16:
		*typed = int16(val)
	case *uint32:
		*typed = uint32(val)
	case *int32:
		*typed = int32(val)
	case *uint64:
		*typed = uint64(val)
	case *int64:
		*typed = int64(val)
	default:
		return n, &TypeDecodeError{"uint64",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}
	return n, nil

}

func decodeFloat32(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {
	buf := make([]byte, 4)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"float32",
			err.Error(),
		}
	}

	val := math.Float32frombits(binary.BigEndian.Uint32(buf))

	switch typed := into.(type) {
	case *float32:
		*typed = float32(val)
	case *float64:
		*typed = float64(val)
	default:
		return n, &TypeDecodeError{"float32",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}
	return n, nil
}

func decodeFloat64(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {

	buf := make([]byte, 8)
	n, err := r.Read(buf)
	if err != nil {
		return n, &TypeDecodeError{
			"float64",
			err.Error(),
		}
	}

	val := math.Float64frombits(binary.BigEndian.Uint64(buf))

	switch typed := into.(type) {
	case *float32:
		*typed = float32(val)
	case *float64:
		*typed = float64(val)
	default:
		return n, &TypeDecodeError{"float64",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}
	return n, nil
}

func decodeString(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {

	var size, total int

	n, err := decodeNullableSize(r, &size)
	if err != nil {
		return n, err
	}
	total += n
	buf := make([]byte, size)
	nRead, err := r.Read(buf)
	total += nRead
	if err != nil {
		return total, &TypeDecodeError{"string",
			fmt.Sprintf("read bytes<%d> err: <%s>", n, err.Error())}
	}

	switch typed := into.(type) {
	case *string:
		*typed = string(buf)
	case *[]byte:
		*typed = buf
	case []byte:
		copy(typed, buf)
	default:
		return total, &TypeDecodeError{"string",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}

	return total, nil
}

func decodeNullableSize(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {

	readByte := func(fnName string) (int, byte, error) {
		buf := make([]byte, 1)
		n, err := r.Read(buf)
		if err != nil {
			return n, 0, &TypeDecodeError{
				fnName,
				err.Error(),
			}
		}
		b1 := buf[0]
		return n, b1, err
	}

	size := 0
	var total int
	n, b1, err := readByte("nullableSize")
	total += n
	if err != nil {
		return total, err
	}

	cur := int(b1 & 0xFF)
	if (cur & 0xFFFFFF80) == 0x0 {
		size = cur & 0x3F
		if cur&0x40 == 0x0 {
			return total, applyNullableSize(size, into)
		}

		shift := NULL_SHIFT
		n, b1, err := readByte("nullableSize")
		total += n
		if err != nil {
			return total, err
		}
		cur := int(b1 & 0xFF)
		size += (cur & 0x7F) << NULL_SHIFT
		shift += MAX_SHIFT

		for (cur & 0xFFFFFF80) != 0x0 {

			n, b1, err := readByte("nullableSize")
			total += n
			if err != nil {
				return total, err
			}

			cur = int(b1 & 0xFF)
			size += (cur & 0x7F) << shift
			shift += MAX_SHIFT

		}
		return total, applyNullableSize(size, into)
	}

	if (cur & 0x7F) != 0x0 {
		return total, &TypeDecodeError{
			"nullableSize",
			"null expected",
		}
	}

	return total, applyNullableSize(size, into)
}

func applyNullableSize(val int, into interface{}) error {
	switch typed := into.(type) {
	case *int:
		*typed = int(val)
	case *uint16:
		*typed = uint16(val)
	case *int16:
		*typed = int16(val)
	case *uint32:
		*typed = uint32(val)
	case *int32:
		*typed = int32(val)
	case *uint64:
		*typed = uint64(val)
	case *int64:
		*typed = int64(val)
	default:
		return &TypeDecodeError{"nullableSize",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}
	return nil
}

func decodeSize(r io.Reader, into interface{}, opts ...map[string]string) (int, error) {

	readByte := func(fnName string) (int, byte, error) {
		buf := make([]byte, 1)
		n, err := r.Read(buf)
		if err != nil {
			return n, 0, &TypeDecodeError{
				fnName,
				err.Error(),
			}
		}
		b1 := buf[0]
		return n, b1, err
	}
	var total int
	ff := 0xFFFFFF80
	n, b1, err := readByte("size")
	total += n
	if err != nil {
		return total, err
	}

	cur := int(b1 & 0xFF)
	size := cur & 0x7F
	for shift := MAX_SHIFT; (cur & ff) != 0x0; {

		n, b1, err = readByte("size")
		total += n
		if err != nil {
			return total, err
		}

		cur = int(b1 & 0xFF)
		size += (cur & 0x7F) << shift
		shift += MAX_SHIFT
	}

	switch typed := into.(type) {
	case *int:
		*typed = int(size)
	case *uint16:
		*typed = uint16(size)
	case *int16:
		*typed = int16(size)
	case *uint32:
		*typed = uint32(size)
	case *int32:
		*typed = int32(size)
	case *uint64:
		*typed = uint64(size)
	case *int64:
		*typed = int64(size)
	default:
		return total, &TypeDecodeError{"size",
			fmt.Sprintf("convert to <%s> unsupporsed", typed)}
	}
	return total, nil
}

type TypeDecodeError struct {
	Mame string
	Msg  string
}

func (e *TypeDecodeError) Error() string {
	// if e.Type == nil {
	// 	return "ras: Decode(nil)"
	// }

	// if e.Type.Kind() != reflect.Ptr {
	// 	return "ras: Decode(non-pointer " + e.Type.String() + ")"
	// }
	return "ras: (decoderFunc " + e.Mame + ") " + e.Msg + ""
}
