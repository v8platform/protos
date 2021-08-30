// Code generated by github.com/v8platform/protoc-gen-go-ras. DO NOT EDIT.

package protocolv1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	io "io"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.
// codec256.io.

func (x *NegotiateMessage) GetPacketType() PacketType {
	return PacketType_PACKET_TYPE_NEGOTIATE
}

func (x *NegotiateMessage) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Magic opts: encoder:"int32"  order:1
	if err := codec256.ParseInt(reader, &x.Magic); err != nil {
		return err
	}
	// decode x.Protocol opts: encoder:"short"  order:2
	if err := codec256.ParseShort(reader, &x.Protocol); err != nil {
		return err
	}
	// decode x.Version opts: encoder:"short"  order:3
	if err := codec256.ParseShort(reader, &x.Version); err != nil {
		return err
	}
	return nil
}
func (x *NegotiateMessage) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Magic opts: encoder:"int32"  order:1
	if err := codec256.FormatInt(writer, x.Magic); err != nil {
		return err
	}
	// decode x.Protocol opts: encoder:"short"  order:2
	if err := codec256.FormatShort(writer, x.Protocol); err != nil {
		return err
	}
	// decode x.Version opts: encoder:"short"  order:3
	if err := codec256.FormatShort(writer, x.Version); err != nil {
		return err
	}
	return nil
}
func (x *ConnectMessage) GetPacketType() PacketType {
	return PacketType_PACKET_TYPE_CONNECT
}

func (x *ConnectMessage) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Params opts: order:1
	// TODO generate map
	return nil
}
func (x *ConnectMessage) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Params opts: order:1
	// TODO generate map
	return nil
}
func (x *DisconnectMessage) GetPacketType() PacketType {
	return PacketType_PACKET_TYPE_DISCONNECT
}

func (x *DisconnectMessage) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Params opts: order:1
	// TODO generate map
	return nil
}
func (x *DisconnectMessage) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Params opts: order:1
	// TODO generate map
	return nil
}
func (x *ConnectMessageAck) GetPacketType() PacketType {
	return PacketType_PACKET_TYPE_CONNECT_ACK
}

func (x *ConnectMessageAck) Parse(reader io.Reader, version int32) error {
	return nil
}
func (x *ConnectMessageAck) Formatter(writer io.Writer, version int32) error {
	return nil
}