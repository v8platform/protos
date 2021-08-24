// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: ras/protocol/v1/packet.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Ras.Protocol.V1 {

  /// <summary>Holder for reflection information generated from ras/protocol/v1/packet.proto</summary>
  public static partial class PacketReflection {

    #region Descriptor
    /// <summary>File descriptor for ras/protocol/v1/packet.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static PacketReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChxyYXMvcHJvdG9jb2wvdjEvcGFja2V0LnByb3RvEg9yYXMucHJvdG9jb2wu",
            "djEaGWdvb2dsZS9wcm90b2J1Zi9hbnkucHJvdG8aIGdvb2dsZS9wcm90b2J1",
            "Zi9kZXNjcmlwdG9yLnByb3RvGhtyYXMvcHJvdG9jb2wvdjEvcGFyYW0ucHJv",
            "dG8aG3Jhcy9wcm90b2NvbC92MS90eXBlcy5wcm90bxoWcmFzL2VuY29kaW5n",
            "L3Jhcy5wcm90byLcAQoGUGFja2V0ElYKBHR5cGUYASABKA4yGy5yYXMucHJv",
            "dG9jb2wudjEuUGFja2V0VHlwZUIlmoSeAxJyYXNiaW5hcnk6ImJ5dGUsMSKC",
            "9eqUDggKBGJ5dGUQAVIEdHlwZRI5CgRzaXplGAIgASgFQiWahJ4DEnJhc2Jp",
            "bmFyeToic2l6ZSwyIoL16pQOCAoEc2l6ZRACUgRzaXplEj8KBGRhdGEYAyAB",
            "KAxCK5qEngMTcmFzYmluYXJ5OiJieXRlcywzIoL16pQODQoFYnl0ZXMQAzgB",
            "UAJSBGRhdGFCyQEKE2NvbS5yYXMucHJvdG9jb2wudjFCC1BhY2tldFByb3Rv",
            "UAFaO2dpdGh1Yi5jb20vdjhwbGF0Zm9ybS9wcm90b3MvZ2VuL3Jhcy9wcm90",
            "b2NvbC92MTtwcm90b2NvbHYxogIDUlBYqgIPUmFzLlByb3RvY29sLlYxygIP",
            "UmFzXFByb3RvY29sXFYx4gIbUmFzXFByb3RvY29sXFYxXEdQQk1ldGFkYXRh",
            "6gIRUmFzOjpQcm90b2NvbDo6VjGC9eqUDgYIARABGAFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Protobuf.WellKnownTypes.AnyReflection.Descriptor, global::Google.Protobuf.Reflection.DescriptorReflection.Descriptor, global::Ras.Protocol.V1.ParamReflection.Descriptor, global::Ras.Protocol.V1.TypesReflection.Descriptor, global::Ras.Encoding.RasReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Ras.Protocol.V1.Packet), global::Ras.Protocol.V1.Packet.Parser, new[]{ "Type", "Size", "Data" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  /// <summary>
  /// Порядок кодирования/декодирования в формат RAS
  /// </summary>
  public sealed partial class Packet : pb::IMessage<Packet>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Packet> _parser = new pb::MessageParser<Packet>(() => new Packet());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Packet> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Ras.Protocol.V1.PacketReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Packet() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Packet(Packet other) : this() {
      type_ = other.type_;
      size_ = other.size_;
      data_ = other.data_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Packet Clone() {
      return new Packet(this);
    }

    /// <summary>Field number for the "type" field.</summary>
    public const int TypeFieldNumber = 1;
    private global::Ras.Protocol.V1.PacketType type_ = global::Ras.Protocol.V1.PacketType.Negotiate;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Ras.Protocol.V1.PacketType Type {
      get { return type_; }
      set {
        type_ = value;
      }
    }

    /// <summary>Field number for the "size" field.</summary>
    public const int SizeFieldNumber = 2;
    private int size_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Size {
      get { return size_; }
      set {
        size_ = value;
      }
    }

    /// <summary>Field number for the "data" field.</summary>
    public const int DataFieldNumber = 3;
    private pb::ByteString data_ = pb::ByteString.Empty;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pb::ByteString Data {
      get { return data_; }
      set {
        data_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Packet);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Packet other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Type != other.Type) return false;
      if (Size != other.Size) return false;
      if (Data != other.Data) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (Type != global::Ras.Protocol.V1.PacketType.Negotiate) hash ^= Type.GetHashCode();
      if (Size != 0) hash ^= Size.GetHashCode();
      if (Data.Length != 0) hash ^= Data.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (Type != global::Ras.Protocol.V1.PacketType.Negotiate) {
        output.WriteRawTag(8);
        output.WriteEnum((int) Type);
      }
      if (Size != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Size);
      }
      if (Data.Length != 0) {
        output.WriteRawTag(26);
        output.WriteBytes(Data);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (Type != global::Ras.Protocol.V1.PacketType.Negotiate) {
        output.WriteRawTag(8);
        output.WriteEnum((int) Type);
      }
      if (Size != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Size);
      }
      if (Data.Length != 0) {
        output.WriteRawTag(26);
        output.WriteBytes(Data);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (Type != global::Ras.Protocol.V1.PacketType.Negotiate) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Type);
      }
      if (Size != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Size);
      }
      if (Data.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeBytesSize(Data);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(Packet other) {
      if (other == null) {
        return;
      }
      if (other.Type != global::Ras.Protocol.V1.PacketType.Negotiate) {
        Type = other.Type;
      }
      if (other.Size != 0) {
        Size = other.Size;
      }
      if (other.Data.Length != 0) {
        Data = other.Data;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Type = (global::Ras.Protocol.V1.PacketType) input.ReadEnum();
            break;
          }
          case 16: {
            Size = input.ReadInt32();
            break;
          }
          case 26: {
            Data = input.ReadBytes();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 8: {
            Type = (global::Ras.Protocol.V1.PacketType) input.ReadEnum();
            break;
          }
          case 16: {
            Size = input.ReadInt32();
            break;
          }
          case 26: {
            Data = input.ReadBytes();
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
