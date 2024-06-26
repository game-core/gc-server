// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: health/masterHealth/master_health.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Api.Game {

  /// <summary>Holder for reflection information generated from health/masterHealth/master_health.proto</summary>
  public static partial class MasterHealthReflection {

    #region Descriptor
    /// <summary>File descriptor for health/masterHealth/master_health.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MasterHealthReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CidoZWFsdGgvbWFzdGVySGVhbHRoL21hc3Rlcl9oZWFsdGgucHJvdG8SCGFw",
            "aS5nYW1lGixoZWFsdGgvbWFzdGVySGVhbHRoL21hc3Rlcl9oZWFsdGhfZW51",
            "bS5wcm90byKJAQoMTWFzdGVySGVhbHRoEhsKCWhlYWx0aF9pZBgBIAEoA1II",
            "aGVhbHRoSWQSEgoEbmFtZRgCIAEoCVIEbmFtZRJIChJtYXN0ZXJfaGVhbHRo",
            "X2VudW0YAyABKA4yGi5hcGkuZ2FtZS5NYXN0ZXJIZWFsdGhFbnVtUhBtYXN0",
            "ZXJIZWFsdGhFbnVtQrIBCgxjb20uYXBpLmdhbWVCEU1hc3RlckhlYWx0aFBy",
            "b3RvUAFaTmdpdGh1Yi5jb20vZ2FtZS1jb3JlL2djLXNlcnZlci9hcGkvZ2Ft",
            "ZS9wcmVzZW50YXRpb24vcHJvdG8vaGVhbHRoL21hc3RlckhlYWx0aKICA0FH",
            "WKoCCEFwaS5HYW1lygIIQXBpXEdhbWXiAhRBcGlcR2FtZVxHUEJNZXRhZGF0",
            "YeoCCUFwaTo6R2FtZWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Api.Game.MasterHealthEnumReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Api.Game.MasterHealth), global::Api.Game.MasterHealth.Parser, new[]{ "HealthId", "Name", "MasterHealthEnum" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  [global::System.Diagnostics.DebuggerDisplayAttribute("{ToString(),nq}")]
  public sealed partial class MasterHealth : pb::IMessage<MasterHealth>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MasterHealth> _parser = new pb::MessageParser<MasterHealth>(() => new MasterHealth());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<MasterHealth> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Api.Game.MasterHealthReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MasterHealth() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MasterHealth(MasterHealth other) : this() {
      healthId_ = other.healthId_;
      name_ = other.name_;
      masterHealthEnum_ = other.masterHealthEnum_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MasterHealth Clone() {
      return new MasterHealth(this);
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

    /// <summary>Field number for the "master_health_enum" field.</summary>
    public const int MasterHealthEnumFieldNumber = 3;
    private global::Api.Game.MasterHealthEnum masterHealthEnum_ = global::Api.Game.MasterHealthEnum.MasterNone;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Api.Game.MasterHealthEnum MasterHealthEnum {
      get { return masterHealthEnum_; }
      set {
        masterHealthEnum_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as MasterHealth);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(MasterHealth other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (HealthId != other.HealthId) return false;
      if (Name != other.Name) return false;
      if (MasterHealthEnum != other.MasterHealthEnum) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (HealthId != 0L) hash ^= HealthId.GetHashCode();
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (MasterHealthEnum != global::Api.Game.MasterHealthEnum.MasterNone) hash ^= MasterHealthEnum.GetHashCode();
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
      if (MasterHealthEnum != global::Api.Game.MasterHealthEnum.MasterNone) {
        output.WriteRawTag(24);
        output.WriteEnum((int) MasterHealthEnum);
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
      if (MasterHealthEnum != global::Api.Game.MasterHealthEnum.MasterNone) {
        output.WriteRawTag(24);
        output.WriteEnum((int) MasterHealthEnum);
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
      if (MasterHealthEnum != global::Api.Game.MasterHealthEnum.MasterNone) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) MasterHealthEnum);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(MasterHealth other) {
      if (other == null) {
        return;
      }
      if (other.HealthId != 0L) {
        HealthId = other.HealthId;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.MasterHealthEnum != global::Api.Game.MasterHealthEnum.MasterNone) {
        MasterHealthEnum = other.MasterHealthEnum;
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
            MasterHealthEnum = (global::Api.Game.MasterHealthEnum) input.ReadEnum();
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
            MasterHealthEnum = (global::Api.Game.MasterHealthEnum) input.ReadEnum();
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
