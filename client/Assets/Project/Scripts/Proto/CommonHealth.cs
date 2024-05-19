// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: health/commonHealth/common_health.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Api.Game {

  /// <summary>Holder for reflection information generated from health/commonHealth/common_health.proto</summary>
  public static partial class CommonHealthReflection {

    #region Descriptor
    /// <summary>File descriptor for health/commonHealth/common_health.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static CommonHealthReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CidoZWFsdGgvY29tbW9uSGVhbHRoL2NvbW1vbl9oZWFsdGgucHJvdG8SCGFw",
            "aS5nYW1lGjFoZWFsdGgvY29tbW9uSGVhbHRoL2NvbW1vbl9oZWFsdGhfdHlw",
            "ZV9lbnVtLnByb3RvIokBCgxDb21tb25IZWFsdGgSGwoJaGVhbHRoX2lkGAEg",
            "ASgDUghoZWFsdGhJZBISCgRuYW1lGAIgASgJUgRuYW1lEkgKEmNvbW1vbl9o",
            "ZWFsdGhfdHlwZRgDIAEoDjIaLmFwaS5nYW1lLkNvbW1vbkhlYWx0aFR5cGVS",
            "EGNvbW1vbkhlYWx0aFR5cGVCtQEKDGNvbS5hcGkuZ2FtZUIRQ29tbW9uSGVh",
            "bHRoUHJvdG9QAVpRZ2l0aHViLmNvbS9nYW1lLWNvcmUvZ2FtZS1zZXJ2ZXIv",
            "YXBpL2dhbWUvcHJlc2VudGF0aW9uL3NlcnZlci9oZWFsdGgvY29tbW9uSGVh",
            "bHRoogIDQUdYqgIIQXBpLkdhbWXKAghBcGlcR2FtZeICFEFwaVxHYW1lXEdQ",
            "Qk1ldGFkYXRh6gIJQXBpOjpHYW1lYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Api.Game.CommonHealthTypeEnumReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Api.Game.CommonHealth), global::Api.Game.CommonHealth.Parser, new[]{ "HealthId", "Name", "CommonHealthType" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  [global::System.Diagnostics.DebuggerDisplayAttribute("{ToString(),nq}")]
  public sealed partial class CommonHealth : pb::IMessage<CommonHealth>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<CommonHealth> _parser = new pb::MessageParser<CommonHealth>(() => new CommonHealth());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<CommonHealth> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Api.Game.CommonHealthReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CommonHealth() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CommonHealth(CommonHealth other) : this() {
      healthId_ = other.healthId_;
      name_ = other.name_;
      commonHealthType_ = other.commonHealthType_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CommonHealth Clone() {
      return new CommonHealth(this);
    }

    /// <summary>Field number for the "health_id" field.</summary>
    public const int HealthIdFieldNumber = 1;
    private long healthId_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long HealthId {
      get { return healthId_; }
      set {
        healthId_ = value;
      }
    }

    /// <summary>Field number for the "name" field.</summary>
    public const int NameFieldNumber = 2;
    private string name_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "common_health_type" field.</summary>
    public const int CommonHealthTypeFieldNumber = 3;
    private global::Api.Game.CommonHealthType commonHealthType_ = global::Api.Game.CommonHealthType.CommonNone;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Api.Game.CommonHealthType CommonHealthType {
      get { return commonHealthType_; }
      set {
        commonHealthType_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as CommonHealth);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(CommonHealth other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (HealthId != other.HealthId) return false;
      if (Name != other.Name) return false;
      if (CommonHealthType != other.CommonHealthType) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (HealthId != 0L) hash ^= HealthId.GetHashCode();
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (CommonHealthType != global::Api.Game.CommonHealthType.CommonNone) hash ^= CommonHealthType.GetHashCode();
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
      if (HealthId != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(HealthId);
      }
      if (Name.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Name);
      }
      if (CommonHealthType != global::Api.Game.CommonHealthType.CommonNone) {
        output.WriteRawTag(24);
        output.WriteEnum((int) CommonHealthType);
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
      if (HealthId != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(HealthId);
      }
      if (Name.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Name);
      }
      if (CommonHealthType != global::Api.Game.CommonHealthType.CommonNone) {
        output.WriteRawTag(24);
        output.WriteEnum((int) CommonHealthType);
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
      if (HealthId != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(HealthId);
      }
      if (Name.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (CommonHealthType != global::Api.Game.CommonHealthType.CommonNone) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) CommonHealthType);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(CommonHealth other) {
      if (other == null) {
        return;
      }
      if (other.HealthId != 0L) {
        HealthId = other.HealthId;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.CommonHealthType != global::Api.Game.CommonHealthType.CommonNone) {
        CommonHealthType = other.CommonHealthType;
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
            HealthId = input.ReadInt64();
            break;
          }
          case 18: {
            Name = input.ReadString();
            break;
          }
          case 24: {
            CommonHealthType = (global::Api.Game.CommonHealthType) input.ReadEnum();
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
            HealthId = input.ReadInt64();
            break;
          }
          case 18: {
            Name = input.ReadString();
            break;
          }
          case 24: {
            CommonHealthType = (global::Api.Game.CommonHealthType) input.ReadEnum();
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
