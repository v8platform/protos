package main

import (
	"github.com/v8platform/protos/gen/ras/encoding"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type MessageFieldExtension struct {
	*encoding.EncodingFieldOptions
}

func GetMessageExtensionFor(messageDesc protoreflect.MessageDescriptor) *MessageFieldExtension {
	opts := messageDesc.Options().(*descriptorpb.MessageOptions)
	if opts == nil || !proto.HasExtension(opts, encoding.E_Field) {
		return nil
	}

	ext := proto.GetExtension(opts, encoding.E_Field).(*encoding.EncodingFieldOptions)

	return &MessageFieldExtension{ext}
}
