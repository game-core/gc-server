// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: exchange/exchange_update_response.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Api.Game {

  /// <summary>Holder for reflection information generated from exchange/exchange_update_response.proto</summary>
  public static partial class ExchangeUpdateResponseReflection {

    #region Descriptor
    /// <summary>File descriptor for exchange/exchange_update_response.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ExchangeUpdateResponseReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CidleGNoYW5nZS9leGNoYW5nZV91cGRhdGVfcmVzcG9uc2UucHJvdG8SCGFw",
            "aS5nYW1lGilleGNoYW5nZS91c2VyRXhjaGFuZ2UvdXNlcl9leGNoYW5nZS5w",
            "cm90bxoyZXhjaGFuZ2UvdXNlckV4Y2hhbmdlSXRlbS91c2VyX2V4Y2hhbmdl",
            "X2l0ZW0ucHJvdG8iuAEKFkV4Y2hhbmdlVXBkYXRlUmVzcG9uc2USQAoNdXNl",
            "cl9leGNoYW5nZRgBIAEoCzIWLmFwaS5nYW1lLlVzZXJFeGNoYW5nZUgAUgx1",
            "c2VyRXhjaGFuZ2WIAQESSgoTdXNlcl9leGNoYW5nZV9pdGVtcxgCIAMoCzIa",
            "LmFwaS5nYW1lLlVzZXJFeGNoYW5nZUl0ZW1SEXVzZXJFeGNoYW5nZUl0ZW1z",
            "QhAKDl91c2VyX2V4Y2hhbmdlQrEBCgxjb20uYXBpLmdhbWVCG0V4Y2hhbmdl",
            "VXBkYXRlUmVzcG9uc2VQcm90b1ABWkNnaXRodWIuY29tL2dhbWUtY29yZS9n",
            "Yy1zZXJ2ZXIvYXBpL2dhbWUvcHJlc2VudGF0aW9uL3Byb3RvL2V4Y2hhbmdl",
            "ogIDQUdYqgIIQXBpLkdhbWXKAghBcGlcR2FtZeICFEFwaVxHYW1lXEdQQk1l",
            "dGFkYXRh6gIJQXBpOjpHYW1lYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Api.Game.UserExchangeReflection.Descriptor, global::Api.Game.UserExchangeItemReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Api.Game.ExchangeUpdateResponse), global::Api.Game.ExchangeUpdateResponse.Parser, new[]{ "UserExchange", "UserExchangeItems" }, new[]{ "UserExchange" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  [global::System.Diagnostics.DebuggerDisplayAttribute("{ToString(),nq}")]
  public sealed partial class ExchangeUpdateResponse : pb::IMessage<ExchangeUpdateResponse>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<ExchangeUpdateResponse> _parser = new pb::MessageParser<ExchangeUpdateResponse>(() => new ExchangeUpdateResponse());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<ExchangeUpdateResponse> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Api.Game.ExchangeUpdateResponseReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public ExchangeUpdateResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public ExchangeUpdateResponse(ExchangeUpdateResponse other) : this() {
      userExchange_ = other.userExchange_ != null ? other.userExchange_.Clone() : null;
      userExchangeItems_ = other.userExchangeItems_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public ExchangeUpdateResponse Clone() {
      return new ExchangeUpdateResponse(this);
    }

    /// <summary>Field number for the "user_exchange" field.</summary>
    public const int UserExchangeFieldNumber = 1;
    private global::Api.Game.UserExchange userExchange_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Api.Game.UserExchange UserExchange {
      get { return userExchange_; }
      set {
        userExchange_ = value;
      }
    }

    /// <summary>Field number for the "user_exchange_items" field.</summary>
    public const int UserExchangeItemsFieldNumber = 2;
    private static readonly pb::FieldCodec<global::Api.Game.UserExchangeItem> _repeated_userExchangeItems_codec
        = pb::FieldCodec.ForMessage(18, global::Api.Game.UserExchangeItem.Parser);
    private readonly pbc::RepeatedField<global::Api.Game.UserExchangeItem> userExchangeItems_ = new pbc::RepeatedField<global::Api.Game.UserExchangeItem>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<global::Api.Game.UserExchangeItem> UserExchangeItems {
      get { return userExchangeItems_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as ExchangeUpdateResponse);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(ExchangeUpdateResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(UserExchange, other.UserExchange)) return false;
      if(!userExchangeItems_.Equals(other.userExchangeItems_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (userExchange_ != null) hash ^= UserExchange.GetHashCode();
      hash ^= userExchangeItems_.GetHashCode();
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
      if (userExchange_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(UserExchange);
      }
      userExchangeItems_.WriteTo(output, _repeated_userExchangeItems_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (userExchange_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(UserExchange);
      }
      userExchangeItems_.WriteTo(ref output, _repeated_userExchangeItems_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (userExchange_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(UserExchange);
      }
      size += userExchangeItems_.CalculateSize(_repeated_userExchangeItems_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(ExchangeUpdateResponse other) {
      if (other == null) {
        return;
      }
      if (other.userExchange_ != null) {
        if (userExchange_ == null) {
          UserExchange = new global::Api.Game.UserExchange();
        }
        UserExchange.MergeFrom(other.UserExchange);
      }
      userExchangeItems_.Add(other.userExchangeItems_);
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
            if (userExchange_ == null) {
              UserExchange = new global::Api.Game.UserExchange();
            }
            input.ReadMessage(UserExchange);
            break;
          }
          case 18: {
            userExchangeItems_.AddEntriesFrom(input, _repeated_userExchangeItems_codec);
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
            if (userExchange_ == null) {
              UserExchange = new global::Api.Game.UserExchange();
            }
            input.ReadMessage(UserExchange);
            break;
          }
          case 18: {
            userExchangeItems_.AddEntriesFrom(ref input, _repeated_userExchangeItems_codec);
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
