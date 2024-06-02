// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: loginBonus/masterLoginBonus/master_login_bonus.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Api.Game {

  /// <summary>Holder for reflection information generated from loginBonus/masterLoginBonus/master_login_bonus.proto</summary>
  public static partial class MasterLoginBonusReflection {

    #region Descriptor
    /// <summary>File descriptor for loginBonus/masterLoginBonus/master_login_bonus.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MasterLoginBonusReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CjRsb2dpbkJvbnVzL21hc3RlckxvZ2luQm9udXMvbWFzdGVyX2xvZ2luX2Jv",
            "bnVzLnByb3RvEghhcGkuZ2FtZSKBAQoQTWFzdGVyTG9naW5Cb251cxIxChVt",
            "YXN0ZXJfbG9naW5fYm9udXNfaWQYASABKANSEm1hc3RlckxvZ2luQm9udXNJ",
            "ZBImCg9tYXN0ZXJfZXZlbnRfaWQYAiABKANSDW1hc3RlckV2ZW50SWQSEgoE",
            "bmFtZRgDIAEoCVIEbmFtZUK/AQoMY29tLmFwaS5nYW1lQhVNYXN0ZXJMb2dp",
            "bkJvbnVzUHJvdG9QAVpXZ2l0aHViLmNvbS9nYW1lLWNvcmUvZ2Mtc2VydmVy",
            "L2FwaS9nYW1lL3ByZXNlbnRhdGlvbi9zZXJ2ZXIvbG9naW5Cb251cy9tYXN0",
            "ZXJMb2dpbkJvbnVzogIDQUdYqgIIQXBpLkdhbWXKAghBcGlcR2FtZeICFEFw",
            "aVxHYW1lXEdQQk1ldGFkYXRh6gIJQXBpOjpHYW1lYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Api.Game.MasterLoginBonus), global::Api.Game.MasterLoginBonus.Parser, new[]{ "MasterLoginBonusId", "MasterEventId", "Name" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  [global::System.Diagnostics.DebuggerDisplayAttribute("{ToString(),nq}")]
  public sealed partial class MasterLoginBonus : pb::IMessage<MasterLoginBonus>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MasterLoginBonus> _parser = new pb::MessageParser<MasterLoginBonus>(() => new MasterLoginBonus());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<MasterLoginBonus> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Api.Game.MasterLoginBonusReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MasterLoginBonus() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MasterLoginBonus(MasterLoginBonus other) : this() {
      masterLoginBonusId_ = other.masterLoginBonusId_;
      masterEventId_ = other.masterEventId_;
      name_ = other.name_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MasterLoginBonus Clone() {
      return new MasterLoginBonus(this);
    }

    /// <summary>Field number for the "master_login_bonus_id" field.</summary>
    public const int MasterLoginBonusIdFieldNumber = 1;
    private long masterLoginBonusId_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long MasterLoginBonusId {
      get { return masterLoginBonusId_; }
      set {
        masterLoginBonusId_ = value;
      }
    }

    /// <summary>Field number for the "master_event_id" field.</summary>
    public const int MasterEventIdFieldNumber = 2;
    private long masterEventId_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long MasterEventId {
      get { return masterEventId_; }
      set {
        masterEventId_ = value;
      }
    }

    /// <summary>Field number for the "name" field.</summary>
    public const int NameFieldNumber = 3;
    private string name_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as MasterLoginBonus);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(MasterLoginBonus other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (MasterLoginBonusId != other.MasterLoginBonusId) return false;
      if (MasterEventId != other.MasterEventId) return false;
      if (Name != other.Name) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (MasterLoginBonusId != 0L) hash ^= MasterLoginBonusId.GetHashCode();
      if (MasterEventId != 0L) hash ^= MasterEventId.GetHashCode();
      if (Name.Length != 0) hash ^= Name.GetHashCode();
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
      if (MasterLoginBonusId != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(MasterLoginBonusId);
      }
      if (MasterEventId != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(MasterEventId);
      }
      if (Name.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Name);
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
      if (MasterLoginBonusId != 0L) {
        output.WriteRawTag(8);
        output.WriteInt64(MasterLoginBonusId);
      }
      if (MasterEventId != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(MasterEventId);
      }
      if (Name.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Name);
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
      if (MasterLoginBonusId != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(MasterLoginBonusId);
      }
      if (MasterEventId != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(MasterEventId);
      }
      if (Name.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(MasterLoginBonus other) {
      if (other == null) {
        return;
      }
      if (other.MasterLoginBonusId != 0L) {
        MasterLoginBonusId = other.MasterLoginBonusId;
      }
      if (other.MasterEventId != 0L) {
        MasterEventId = other.MasterEventId;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
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
            MasterLoginBonusId = input.ReadInt64();
            break;
          }
          case 16: {
            MasterEventId = input.ReadInt64();
            break;
          }
          case 26: {
            Name = input.ReadString();
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
            MasterLoginBonusId = input.ReadInt64();
            break;
          }
          case 16: {
            MasterEventId = input.ReadInt64();
            break;
          }
          case 26: {
            Name = input.ReadString();
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