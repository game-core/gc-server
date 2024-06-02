// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: loginBonus/login_bonus_receive_response.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Api.Game {

  /// <summary>Holder for reflection information generated from loginBonus/login_bonus_receive_response.proto</summary>
  public static partial class LoginBonusReceiveResponseReflection {

    #region Descriptor
    /// <summary>File descriptor for loginBonus/login_bonus_receive_response.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static LoginBonusReceiveResponseReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Ci1sb2dpbkJvbnVzL2xvZ2luX2JvbnVzX3JlY2VpdmVfcmVzcG9uc2UucHJv",
            "dG8SCGFwaS5nYW1lGjBsb2dpbkJvbnVzL3VzZXJMb2dpbkJvbnVzL3VzZXJf",
            "bG9naW5fYm9udXMucHJvdG8ieQoZTG9naW5Cb251c1JlY2VpdmVSZXNwb25z",
            "ZRJHChB1c2VyX2xvZ2luX2JvbnVzGAEgASgLMhguYXBpLmdhbWUuVXNlckxv",
            "Z2luQm9udXNIAFIOdXNlckxvZ2luQm9udXOIAQFCEwoRX3VzZXJfbG9naW5f",
            "Ym9udXNCtwEKDGNvbS5hcGkuZ2FtZUIeTG9naW5Cb251c1JlY2VpdmVSZXNw",
            "b25zZVByb3RvUAFaRmdpdGh1Yi5jb20vZ2FtZS1jb3JlL2djLXNlcnZlci9h",
            "cGkvZ2FtZS9wcmVzZW50YXRpb24vc2VydmVyL2xvZ2luQm9udXOiAgNBR1iq",
            "AghBcGkuR2FtZcoCCEFwaVxHYW1l4gIUQXBpXEdhbWVcR1BCTWV0YWRhdGHq",
            "AglBcGk6OkdhbWViBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Api.Game.UserLoginBonusReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Api.Game.LoginBonusReceiveResponse), global::Api.Game.LoginBonusReceiveResponse.Parser, new[]{ "UserLoginBonus" }, new[]{ "UserLoginBonus" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  [global::System.Diagnostics.DebuggerDisplayAttribute("{ToString(),nq}")]
  public sealed partial class LoginBonusReceiveResponse : pb::IMessage<LoginBonusReceiveResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<LoginBonusReceiveResponse> _parser = new pb::MessageParser<LoginBonusReceiveResponse>(() => new LoginBonusReceiveResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<LoginBonusReceiveResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Api.Game.LoginBonusReceiveResponseReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public LoginBonusReceiveResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public LoginBonusReceiveResponse(LoginBonusReceiveResponse other) : this() {
      userLoginBonus_ = other.userLoginBonus_ != null ? other.userLoginBonus_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public LoginBonusReceiveResponse Clone() {
      return new LoginBonusReceiveResponse(this);
    }

    /// <summary>Field number for the "user_login_bonus" field.</summary>
    public const int UserLoginBonusFieldNumber = 1;
    private global::Api.Game.UserLoginBonus userLoginBonus_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Api.Game.UserLoginBonus UserLoginBonus {
      get { return userLoginBonus_; }
      set {
        userLoginBonus_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as LoginBonusReceiveResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(LoginBonusReceiveResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(UserLoginBonus, other.UserLoginBonus)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (userLoginBonus_ != null) hash ^= UserLoginBonus.GetHashCode();
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
      if (userLoginBonus_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(UserLoginBonus);
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
      if (userLoginBonus_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(UserLoginBonus);
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
      if (userLoginBonus_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(UserLoginBonus);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(LoginBonusReceiveResponse other) {
      if (other == null) {
        return;
      }
      if (other.userLoginBonus_ != null) {
        if (userLoginBonus_ == null) {
          UserLoginBonus = new global::Api.Game.UserLoginBonus();
        }
        UserLoginBonus.MergeFrom(other.UserLoginBonus);
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
            if (userLoginBonus_ == null) {
              UserLoginBonus = new global::Api.Game.UserLoginBonus();
            }
            input.ReadMessage(UserLoginBonus);
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
            if (userLoginBonus_ == null) {
              UserLoginBonus = new global::Api.Game.UserLoginBonus();
            }
            input.ReadMessage(UserLoginBonus);
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