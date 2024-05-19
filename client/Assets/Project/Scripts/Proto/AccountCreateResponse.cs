// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: account/account_create_response.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Api.Game {

  /// <summary>Holder for reflection information generated from account/account_create_response.proto</summary>
  public static partial class AccountCreateResponseReflection {

    #region Descriptor
    /// <summary>File descriptor for account/account_create_response.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static AccountCreateResponseReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiVhY2NvdW50L2FjY291bnRfY3JlYXRlX3Jlc3BvbnNlLnByb3RvEghhcGku",
            "Z2FtZRomYWNjb3VudC91c2VyQWNjb3VudC91c2VyX2FjY291bnQucHJvdG8i",
            "ZwoVQWNjb3VudENyZWF0ZVJlc3BvbnNlEj0KDHVzZXJfYWNjb3VudBgBIAEo",
            "CzIVLmFwaS5nYW1lLlVzZXJBY2NvdW50SABSC3VzZXJBY2NvdW50iAEBQg8K",
            "DV91c2VyX2FjY291bnRCsAEKDGNvbS5hcGkuZ2FtZUIaQWNjb3VudENyZWF0",
            "ZVJlc3BvbnNlUHJvdG9QAVpDZ2l0aHViLmNvbS9nYW1lLWNvcmUvZ2Mtc2Vy",
            "dmVyL2FwaS9nYW1lL3ByZXNlbnRhdGlvbi9zZXJ2ZXIvYWNjb3VudKICA0FH",
            "WKoCCEFwaS5HYW1lygIIQXBpXEdhbWXiAhRBcGlcR2FtZVxHUEJNZXRhZGF0",
            "YeoCCUFwaTo6R2FtZWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Api.Game.UserAccountReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Api.Game.AccountCreateResponse), global::Api.Game.AccountCreateResponse.Parser, new[]{ "UserAccount" }, new[]{ "UserAccount" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  [global::System.Diagnostics.DebuggerDisplayAttribute("{ToString(),nq}")]
  public sealed partial class AccountCreateResponse : pb::IMessage<AccountCreateResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<AccountCreateResponse> _parser = new pb::MessageParser<AccountCreateResponse>(() => new AccountCreateResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<AccountCreateResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Api.Game.AccountCreateResponseReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public AccountCreateResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public AccountCreateResponse(AccountCreateResponse other) : this() {
      userAccount_ = other.userAccount_ != null ? other.userAccount_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public AccountCreateResponse Clone() {
      return new AccountCreateResponse(this);
    }

    /// <summary>Field number for the "user_account" field.</summary>
    public const int UserAccountFieldNumber = 1;
    private global::Api.Game.UserAccount userAccount_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Api.Game.UserAccount UserAccount {
      get { return userAccount_; }
      set {
        userAccount_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as AccountCreateResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(AccountCreateResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(UserAccount, other.UserAccount)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (userAccount_ != null) hash ^= UserAccount.GetHashCode();
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
      if (userAccount_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(UserAccount);
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
      if (userAccount_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(UserAccount);
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
      if (userAccount_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(UserAccount);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(AccountCreateResponse other) {
      if (other == null) {
        return;
      }
      if (other.userAccount_ != null) {
        if (userAccount_ == null) {
          UserAccount = new global::Api.Game.UserAccount();
        }
        UserAccount.MergeFrom(other.UserAccount);
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
          case 10: {
            if (userAccount_ == null) {
              UserAccount = new global::Api.Game.UserAccount();
            }
            input.ReadMessage(UserAccount);
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
          case 10: {
            if (userAccount_ == null) {
              UserAccount = new global::Api.Game.UserAccount();
            }
            input.ReadMessage(UserAccount);
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
