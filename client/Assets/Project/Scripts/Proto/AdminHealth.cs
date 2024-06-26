// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: health/adminHealth/admin_health.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Api.Game {

  /// <summary>Holder for reflection information generated from health/adminHealth/admin_health.proto</summary>
  public static partial class AdminHealthReflection {

    #region Descriptor
    /// <summary>File descriptor for health/adminHealth/admin_health.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static AdminHealthReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiVoZWFsdGgvYWRtaW5IZWFsdGgvYWRtaW5faGVhbHRoLnByb3RvEghhcGku",
            "Z2FtZRoqaGVhbHRoL2FkbWluSGVhbHRoL2FkbWluX2hlYWx0aF9lbnVtLnBy",
            "b3RvIoUBCgtBZG1pbkhlYWx0aBIbCgloZWFsdGhfaWQYASABKANSCGhlYWx0",
            "aElkEhIKBG5hbWUYAiABKAlSBG5hbWUSRQoRYWRtaW5faGVhbHRoX2VudW0Y",
            "AyABKA4yGS5hcGkuZ2FtZS5BZG1pbkhlYWx0aEVudW1SD2FkbWluSGVhbHRo",
            "RW51bUKwAQoMY29tLmFwaS5nYW1lQhBBZG1pbkhlYWx0aFByb3RvUAFaTWdp",
            "dGh1Yi5jb20vZ2FtZS1jb3JlL2djLXNlcnZlci9hcGkvZ2FtZS9wcmVzZW50",
            "YXRpb24vcHJvdG8vaGVhbHRoL2FkbWluSGVhbHRoogIDQUdYqgIIQXBpLkdh",
            "bWXKAghBcGlcR2FtZeICFEFwaVxHYW1lXEdQQk1ldGFkYXRh6gIJQXBpOjpH",
            "YW1lYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Api.Game.AdminHealthEnumReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Api.Game.AdminHealth), global::Api.Game.AdminHealth.Parser, new[]{ "HealthId", "Name", "AdminHealthEnum" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  [global::System.Diagnostics.DebuggerDisplayAttribute("{ToString(),nq}")]
  public sealed partial class AdminHealth : pb::IMessage<AdminHealth>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<AdminHealth> _parser = new pb::MessageParser<AdminHealth>(() => new AdminHealth());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<AdminHealth> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Api.Game.AdminHealthReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public AdminHealth() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public AdminHealth(AdminHealth other) : this() {
      healthId_ = other.healthId_;
      name_ = other.name_;
      adminHealthEnum_ = other.adminHealthEnum_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public AdminHealth Clone() {
      return new AdminHealth(this);
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

    /// <summary>Field number for the "admin_health_enum" field.</summary>
    public const int AdminHealthEnumFieldNumber = 3;
    private global::Api.Game.AdminHealthEnum adminHealthEnum_ = global::Api.Game.AdminHealthEnum.AdminNone;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Api.Game.AdminHealthEnum AdminHealthEnum {
      get { return adminHealthEnum_; }
      set {
        adminHealthEnum_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as AdminHealth);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(AdminHealth other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (HealthId != other.HealthId) return false;
      if (Name != other.Name) return false;
      if (AdminHealthEnum != other.AdminHealthEnum) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (HealthId != 0L) hash ^= HealthId.GetHashCode();
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (AdminHealthEnum != global::Api.Game.AdminHealthEnum.AdminNone) hash ^= AdminHealthEnum.GetHashCode();
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
      if (AdminHealthEnum != global::Api.Game.AdminHealthEnum.AdminNone) {
        output.WriteRawTag(24);
        output.WriteEnum((int) AdminHealthEnum);
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
      if (AdminHealthEnum != global::Api.Game.AdminHealthEnum.AdminNone) {
        output.WriteRawTag(24);
        output.WriteEnum((int) AdminHealthEnum);
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
      if (AdminHealthEnum != global::Api.Game.AdminHealthEnum.AdminNone) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) AdminHealthEnum);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(AdminHealth other) {
      if (other == null) {
        return;
      }
      if (other.HealthId != 0L) {
        HealthId = other.HealthId;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.AdminHealthEnum != global::Api.Game.AdminHealthEnum.AdminNone) {
        AdminHealthEnum = other.AdminHealthEnum;
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
            AdminHealthEnum = (global::Api.Game.AdminHealthEnum) input.ReadEnum();
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
            AdminHealthEnum = (global::Api.Game.AdminHealthEnum) input.ReadEnum();
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
