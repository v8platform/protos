package main

// type MessageFieldExtension struct {
// 	*encoding.EncodingFieldOptions
// }
//
// func GetMessageExtensionFor(messageDesc protoreflect.MessageDescriptor) *MessageFieldExtension {
// 	opts := messageDesc.Options().(*descriptorpb.MessageOptions)
// 	if opts == nil || !proto.HasExtension(opts, encoding.E_Field) {
// 		return nil
// 	}
//
// 	ext := proto.GetExtension(opts, encoding.E_Field).(*encoding.EncodingFieldOptions)
//
// 	return &MessageFieldExtension{ext}
// }
