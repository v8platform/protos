package main

import (
	"encoding/binary"
	"fmt"
	uuid "github.com/satori/go.uuid"
	pb "google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"math"
	"reflect"
	"strings"
	"time"
)

var encoderFunc = map[string]TypeEncoderFunc{}

type TypeEncoderFunc func(r io.Writer, value interface{}) (int, error)

func init() {
	RegisterEncoderType("time", encodeTime)
	RegisterEncoderType("type", encodeType)
	RegisterEncoderType("bool", encodeBool)
	RegisterEncoderType("byte int8 uint8", encodeByte)
	RegisterEncoderType("char short int16 uint16", encodeUint16)
	RegisterEncoderType("int int32 uint32", encodeUint32)
	RegisterEncoderType("int64 uint64 long", encodeUint64)
	RegisterEncoderType("float32", encodeFloat32)
	RegisterEncoderType("float64 double", encodeFloat64)
	RegisterEncoderType("string", encodeString)
	RegisterEncoderType("null-size", encodeNullableSize)
	RegisterEncoderType("size", encodeSize)
	RegisterEncoderType("uuid", encodeUuid)
	RegisterEncoderType("bytes", encodeBytes)
}

func encodeBytes(r io.Writer, value interface{}) (int, error) {
	switch val := value.(type) {
	case []byte:
		return writeBuf("bytes", r, val)
	case *[]byte:
		return writeBuf("bytes", r, *val)
	case *uuid.UUID:
		return writeBuf("bytes", r, val.Bytes())
	case uuid.UUID:
		return writeBuf("bytes", r, val.Bytes())
	case string:
		return writeBuf("bytes", r, []byte(val))
	case *string:
		return writeBuf("bytes", r, []byte(*val))
	default:
		return 0, &TypeEncoderError{"bytes", "unknown bytes type"}
	}
}

func EncodeValue(encoder string, r io.Writer, value interface{}) (int, error) {

	typeEncoderFunc, ok := encoderFunc[encoder]
	if !ok {
		return 0, fmt.Errorf("unknown encoder <%s>", encoder)
	}

	return typeEncoderFunc(r, value)
}

func encodeUuid(r io.Writer, value interface{}) (int, error) {

	switch val := value.(type) {
	case []byte:
		return writeBuf("uuid", r, val)
	case *[]byte:
		return writeBuf("uuid", r, *val)
	case *uuid.UUID:
		return writeBuf("uuid", r, val.Bytes())
	case uuid.UUID:
		return writeBuf("uuid", r, val.Bytes())
	case string:
		return writeBuf("uuid", r, uuid.FromStringOrNil(val).Bytes())
	case *string:
		return writeBuf("uuid", r, uuid.FromStringOrNil(*val).Bytes())
	default:
		return 0, &TypeEncoderError{"uuid", "unknown uuid type"}
	}

}

func RegisterEncoderType(name string, dec TypeEncoderFunc) {

	names := strings.Fields(strings.ToLower(name))

	for _, s := range names {
		encoderFunc[s] = dec
	}
}

func encodeTime(w io.Writer, value interface{}) (int, error) {
	var val int64

	switch tVal := value.(type) {
	case int64:
		val = int64(tVal)
	case uint64:
		val = int64(tVal)
	case *int64:
		val = int64(*tVal)
	case *uint64:
		val = int64(*tVal)
	case time.Time:
		val = tVal.UnixNano()
	case *time.Time:
		val = tVal.UnixNano()
	case pb.Timestamp:
		val = tVal.AsTime().UnixNano()
	case *pb.Timestamp:
		val = tVal.AsTime().UnixNano()
	default:
		return 0, &TypeEncoderError{"time", "TODO"}
	}
	ticks := val / int64(time.Millisecond)
	ticks = ticks*10 + AgeDelta

	return encodeUint64(w, ticks)

}

func encodeUint16(w io.Writer, value interface{}) (int, error) {
	var val uint16

	switch tVal := value.(type) {
	case int16:
		val = uint16(tVal)
	case uint16:
		val = uint16(tVal)
	case int32:
		val = uint16(tVal)
	case uint32:
		val = uint16(tVal)
	case *int32:
		val = uint16(*tVal)
	case *uint32:
		val = uint16(*tVal)
	case *int16:
		val = uint16(*tVal)
	case *uint16:
		val = uint16(*tVal)
	default:
		return 0, &TypeEncoderError{"uint16", "TODO"}
	}
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, val)
	return writeBuf("uint16", w, buf)

}

func encodeUint32(w io.Writer, value interface{}) (int, error) {
	var val uint32

	switch tVal := value.(type) {
	case int:
		val = uint32(tVal)
	case uint:
		val = uint32(tVal)
	case *int:
		val = uint32(*tVal)
	case *uint:
		val = uint32(*tVal)
	case int32:
		val = uint32(tVal)
	case uint32:
		val = uint32(tVal)
	case *int32:
		val = uint32(*tVal)
	case *uint32:
		val = uint32(*tVal)
	default:
		return 0, &TypeEncoderError{"uint32", "TODO"}
	}
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, val)
	return writeBuf("uint32", w, buf)

}

func encodeUint64(w io.Writer, value interface{}) (int, error) {
	var val uint64

	switch tVal := value.(type) {
	case int64:
		val = uint64(tVal)
	case uint64:
		val = uint64(tVal)
	case *int64:
		val = uint64(*tVal)
	case *uint64:
		val = uint64(*tVal)
	default:
		return 0, &TypeEncoderError{"uint64", fmt.Sprintf("%s", reflect.TypeOf(tVal))}
	}
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	return writeBuf("uint64", w, buf)

}

func encodeFloat32(w io.Writer, value interface{}) (int, error) {
	var val float32

	switch tVal := value.(type) {
	case float32:
		val = tVal
	case *float32:
		val = *tVal
	default:
		return 0, &TypeEncoderError{"float32", "TODO"}
	}
	return encodeUint32(w, math.Float32bits(val))
}

func encodeFloat64(w io.Writer, value interface{}) (int, error) {
	var val float64

	switch tVal := value.(type) {
	case float64:
		val = tVal
	case *float64:
		val = *tVal
	default:
		return 0, &TypeEncoderError{"float64", "TODO"}
	}
	return encodeUint64(w, math.Float64bits(val))

}

func encodeString(w io.Writer, value interface{}) (int, error) {
	var val []byte

	switch tVal := value.(type) {
	case []byte:
		val = tVal
	case *[]byte:
		val = *tVal
	case string:
		val = []byte(tVal)
	case *string:
		val = []byte(*tVal)
	default:
		return 0, &TypeEncoderError{"string", "TODO"}
	}

	if len(val) == 0 {
		n, err := writeNull(w)
		if err != nil {
			return 0, err
		}
		return n, nil
	}

	size := len(val)
	n, err := encodeNullableSize(w, size)
	if err != nil {
		return 0, err
	}

	bufN, err := writeBuf("string", w, val)
	if err != nil {
		return bufN + n, err
	}

	return bufN + n, nil

}

func encodeType(w io.Writer, value interface{}) (int, error) {
	var val byte

	switch tVal := value.(type) {
	case int8:
		val = byte(tVal)
	case uint8:
		val = byte(tVal)
	case *int8:
		val = byte(*tVal)
	case *uint8:
		val = byte(*tVal)
	case int:
		val = byte(tVal)
	case uint:
		val = byte(tVal)
	case *int:
		val = byte(*tVal)
	case *uint:
		val = byte(*tVal)
	case int32:
		val = byte(tVal)
	case uint32:
		val = byte(tVal)
	case *int32:
		val = byte(*tVal)
	case *uint32:
		val = byte(*tVal)
	default:
		return 0, &TypeEncoderError{"type", "TODO"}
	}

	if val == NULL_BYTE {
		return writeNull(w)
	}
	return writeBuf("type", w, []byte{val})

}

func encodeBool(w io.Writer, value interface{}) (int, error) {
	var val byte

	switch tVal := value.(type) {
	case int:
		val = byte(tVal)
	case *int:
		val = byte(*tVal)
	case bool:
		val = FALSE_BYTE
		if tVal {
			val = TRUE_BYTE
		}
	case *bool:
		val = FALSE_BYTE
		if *tVal {
			val = TRUE_BYTE
		}
	default:
		return 0, &TypeEncoderError{"bool", "TODO"}
	}

	return writeBuf("bool", w, []byte{val})

}

func encodeByte(w io.Writer, value interface{}) (int, error) {
	var val byte

	switch tVal := value.(type) {
	case int8:
		val = byte(tVal)
	case uint8:
		val = byte(tVal)
	case *int8:
		val = byte(*tVal)
	case *uint8:
		val = byte(*tVal)
	case int:
		val = byte(tVal)
	case uint:
		val = byte(tVal)
	case *int:
		val = byte(*tVal)
	case *uint:
		val = byte(*tVal)
	case int32:
		val = byte(tVal)
	case uint32:
		val = byte(tVal)
	case *int32:
		val = byte(*tVal)
	case *uint32:
		val = byte(*tVal)
	default:
		return 0, &TypeEncoderError{"byte", "TODO"}
	}

	if val == NULL_BYTE {
		return writeNull(w)
	}
	return writeBuf("byte", w, []byte{val})

}

func encodeSize(w io.Writer, value interface{}) (int, error) {

	val, err := castToInt("size", value)
	if err != nil {
		return 0, err
	}

	var b1 int

	msb := val >> MAX_SHIFT
	if msb != 0 {
		b1 = -128
	} else {
		b1 = 0
	}

	var total int
	n, err := writeBuf("size", w, []byte{byte(b1 | (val & 0x7F))})
	if err != nil {
		return n, err
	}
	total += n
	for val = msb; val > 0; val = msb {

		msb >>= MAX_SHIFT
		if msb != 0 {
			b1 = -128
		} else {
			b1 = 0
		}

		n, err := writeBuf("size", w, []byte{byte(b1 | (val & 0x7F))})
		if err != nil {
			return n, err
		}
		total += n
	}

	return total, err
}

func encodeNullableSize(w io.Writer, value interface{}) (int, error) {

	val, err := castToInt("null-size", value)
	if err != nil {
		return 0, err
	}

	var b1 int

	msb := val >> NULL_SHIFT
	if msb != 0 {
		b1 = NULL_NEXT_MASK
	} else {
		b1 = 0
	}

	var total int

	n, err := writeBuf("null-size", w, []byte{byte(b1 | (val & 0x7F))})
	if err != nil {
		return n, err
	}
	total += n

	for val = msb; val > 0; val = msb {

		msb >>= MAX_SHIFT
		if msb != 0 {
			b1 = NEXT_MASK
		} else {
			b1 = 0
		}

		n, err := writeBuf("null-size", w, []byte{byte(b1 | (val & 0x7F))})
		if err != nil {
			return n, err
		}
		total += n
	}

	return total, nil
}

func writeNull(w io.Writer) (int, error) {
	return writeBuf("write null", w, []byte{0x00})
}

func writeBuf(fnName string, w io.Writer, buf []byte) (int, error) {

	n, err := w.Write(buf)
	if err != nil {
		return n, &EncoderWriteError{fnName, n, err}
	}

	return n, nil
}

func castToInt(fnName string, value interface{}) (int, error) {
	var val int

	switch tVal := value.(type) {
	case int:
		val = int(tVal)
	case uint:
		val = int(tVal)
	case *int:
		val = int(*tVal)
	case *uint:
		val = int(*tVal)
	case int32:
		val = int(tVal)
	case uint32:
		val = int(tVal)
	case *int32:
		val = int(*tVal)
	case *uint32:
		val = int(*tVal)
	case int64:
		val = int(tVal)
	case uint64:
		val = int(tVal)
	case *int64:
		val = int(*tVal)
	case *uint64:
		val = int(*tVal)
	default:
		return 0, &TypeEncoderError{fnName, "TODO"}
	}

	return val, nil
}

type EncoderWriteError struct {
	Mame   string
	writeN int
	err    error
}

func (e *EncoderWriteError) Error() string {

	return "ras: (encoderFunc " + e.Mame + ") write" + e.err.Error() + ""
}

type TypeEncoderError struct {
	Mame string
	Msg  string
}

func (e *TypeEncoderError) Error() string {
	// if e.Type == nil {
	// 	return "ras: Decode(nil)"
	// }

	// if e.Type.Kind() != reflect.Ptr {
	// 	return "ras: Decode(non-pointer " + e.Type.String() + ")"
	// }
	return "ras: (encoderFunc " + e.Mame + ") " + e.Msg + ""
}
