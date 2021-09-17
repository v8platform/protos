// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package messagesv1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	v1 "github.com/v8platform/protos/gen/v8platform/serialize/v1"
	io "io"
)

func (x *GetWorkingServersRequest) GetMessageType() MessageType {
	return MessageType_GET_WORKING_SERVERS_REQUEST
}

func (x *GetWorkingServersRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingServersRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingServersResponse) GetMessageType() MessageType {
	return MessageType_GET_WORKING_SERVERS_RESPONSE
}

func (x *GetWorkingServersResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Servers opts: order:1
	x.Servers = &v1.ServerInfo{}
	if err := x.Servers.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *GetWorkingServersResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Servers opts: order:1
	if err := x.Servers.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingServerInfoRequest) GetMessageType() MessageType {
	return MessageType_GET_WORKING_SERVER_INFO_REQUEST
}

func (x *GetWorkingServerInfoRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.ServerId opts: encoder:"uuid" order:2
	if err := codec256.ParseUUID(reader, &x.ServerId); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingServerInfoRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.ServerId opts: encoder:"uuid" order:2
	if err := codec256.FormatUuid(writer, x.ServerId); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingServerInfoResponse) GetMessageType() MessageType {
	return MessageType_GET_WORKING_SERVER_INFO_RESPONSE
}

func (x *GetWorkingServerInfoResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Info opts: order:1
	x.Info = &v1.ServerInfo{}
	if err := x.Info.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *GetWorkingServerInfoResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Info opts: order:1
	if err := x.Info.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *AddWorkingServerRequest) GetMessageType() MessageType {
	return MessageType_REG_WORKING_SERVER_REQUEST
}

func (x *AddWorkingServerRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.Info opts: order:2
	x.Info = &v1.ServerInfo{}
	if err := x.Info.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *AddWorkingServerRequest) Formatter(writer io.Writer, version int32) error {
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
func (x *AddWorkingServerResponse) GetMessageType() MessageType {
	return MessageType_REG_WORKING_SERVER_RESPONSE
}

func (x *AddWorkingServerResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ServerId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ServerId); err != nil {
		return err
	}
	return nil
}
func (x *AddWorkingServerResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ServerId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ServerId); err != nil {
		return err
	}
	return nil
}
func (x *DeleteWorkingServerRequest) GetMessageType() MessageType {
	return MessageType_UNREG_WORKING_SERVER_REQUEST
}

func (x *DeleteWorkingServerRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.ServerId opts: encoder:"uuid" order:2
	if err := codec256.ParseUUID(reader, &x.ServerId); err != nil {
		return err
	}
	return nil
}
func (x *DeleteWorkingServerRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.ServerId opts: encoder:"uuid" order:2
	if err := codec256.FormatUuid(writer, x.ServerId); err != nil {
		return err
	}
	return nil
}
