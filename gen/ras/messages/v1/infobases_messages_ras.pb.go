// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package messagesv1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	v1 "github.com/v8platform/protos/gen/v8platform/serialize/v1"
	io "io"
)

func (x *GetInfobasesSummaryRequest) GetMessageType() MessageType {
	return MessageType_GET_INFOBASES_SHORT_REQUEST
}

func (x *GetInfobasesSummaryRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetInfobasesSummaryRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetInfobasesSummaryResponse) GetMessageType() MessageType {
	return MessageType_GET_INFOBASES_SHORT_RESPONSE
}

func (x *GetInfobasesSummaryResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Infobases opts: order:1
	var size_Infobases int
	if err := codec256.ParseSize(reader, &size_Infobases); err != nil {
		return err
	}
	for i := 0; i < size_Infobases; i++ {
		val := &v1.InfobaseSummaryInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Infobases = append(x.Infobases, val)
	}
	return nil
}
func (x *GetInfobasesSummaryResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Infobases opts: order:1
	if err := codec256.FormatSize(writer, len(x.Infobases)); err != nil {
		return err
	}
	for i := 0; i < len(x.Infobases); i++ {
		if err := x.Infobases[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *GetInfobasesRequest) GetMessageType() MessageType {
	return MessageType_GET_INFOBASES_REQUEST
}

func (x *GetInfobasesRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetInfobasesRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetInfobasesResponse) GetMessageType() MessageType {
	return MessageType_GET_INFOBASES_RESPONSE
}

func (x *GetInfobasesResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Infobases opts: order:1
	var size_Infobases int
	if err := codec256.ParseSize(reader, &size_Infobases); err != nil {
		return err
	}
	for i := 0; i < size_Infobases; i++ {
		val := &v1.InfobaseInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Infobases = append(x.Infobases, val)
	}
	return nil
}
func (x *GetInfobasesResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Infobases opts: order:1
	if err := codec256.FormatSize(writer, len(x.Infobases)); err != nil {
		return err
	}
	for i := 0; i < len(x.Infobases); i++ {
		if err := x.Infobases[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *GetInfobaseInfoRequest) GetMessageType() MessageType {
	return MessageType_GET_INFOBASE_INFO_REQUEST
}

func (x *GetInfobaseInfoRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.InfobaseId opts: encoder:"uuid" order:2
	if err := codec256.ParseUUID(reader, &x.InfobaseId); err != nil {
		return err
	}
	return nil
}
func (x *GetInfobaseInfoRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.InfobaseId opts: encoder:"uuid" order:2
	if err := codec256.FormatUuid(writer, x.InfobaseId); err != nil {
		return err
	}
	return nil
}
func (x *GetInfobaseInfoResponse) GetMessageType() MessageType {
	return MessageType_GET_INFOBASE_INFO_RESPONSE
}

func (x *GetInfobaseInfoResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Info opts: order:1
	x.Info = &v1.InfobaseInfo{}
	if err := x.Info.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *GetInfobaseInfoResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Info opts: order:1
	if err := x.Info.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *CreateInfobaseRequest) GetMessageType() MessageType {
	return MessageType_CREATE_INFOBASE_REQUEST
}

func (x *CreateInfobaseRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.Info opts: order:2
	x.Info = &v1.InfobaseInfo{}
	if err := x.Info.Parse(reader, version); err != nil {
		return err
	}

	// decode x.Mode opts: encoder:"int" order:3
	if err := codec256.ParseInt(reader, &x.Mode); err != nil {
		return err
	}
	return nil
}
func (x *CreateInfobaseRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.Info opts: order:2
	if err := x.Info.Formatter(writer, version); err != nil {
		return err
	}
	// decode x.Mode opts: encoder:"int" order:3
	if err := codec256.FormatInt(writer, x.Mode); err != nil {
		return err
	}
	return nil
}
func (x *CreateInfobaseResponse) GetMessageType() MessageType {
	return MessageType_CREATE_INFOBASE_RESPONSE
}

func (x *CreateInfobaseResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.InfobaseId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.InfobaseId); err != nil {
		return err
	}
	return nil
}
func (x *CreateInfobaseResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.InfobaseId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.InfobaseId); err != nil {
		return err
	}
	return nil
}
func (x *DropInfobaseRequest) GetMessageType() MessageType {
	return MessageType_DROP_INFOBASE_REQUEST
}

func (x *DropInfobaseRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.InfobaseId opts: encoder:"uuid" order:2
	if err := codec256.ParseUUID(reader, &x.InfobaseId); err != nil {
		return err
	}
	// decode x.Mode opts: encoder:"int" order:3
	if err := codec256.ParseInt(reader, &x.Mode); err != nil {
		return err
	}
	return nil
}
func (x *DropInfobaseRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.InfobaseId opts: encoder:"uuid" order:2
	if err := codec256.FormatUuid(writer, x.InfobaseId); err != nil {
		return err
	}
	// decode x.Mode opts: encoder:"int" order:3
	if err := codec256.FormatInt(writer, x.Mode); err != nil {
		return err
	}
	return nil
}
func (x *UpdateInfobaseSummaryRequest) GetMessageType() MessageType {
	return MessageType_UPDATE_INFOBASE_SHORT_REQUEST
}

func (x *UpdateInfobaseSummaryRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.Summary opts: order:2
	x.Summary = &v1.InfobaseSummaryInfo{}
	if err := x.Summary.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *UpdateInfobaseSummaryRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.Summary opts: order:2
	if err := x.Summary.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *UpdateInfobaseRequest) GetMessageType() MessageType {
	return MessageType_UPDATE_INFOBASE_REQUEST
}

func (x *UpdateInfobaseRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.Info opts: order:2
	x.Info = &v1.InfobaseInfo{}
	if err := x.Info.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *UpdateInfobaseRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.Info opts: order:2
	if err := x.Info.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
