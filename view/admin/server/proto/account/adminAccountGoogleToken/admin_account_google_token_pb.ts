// 管理者アカウントのGoogleToken

// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file account/adminAccountGoogleToken/admin_account_google_token.proto (package api.admin, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
  BinaryReadOptions,
  FieldList,
  JsonReadOptions,
  JsonValue,
  PartialMessage,
  PlainMessage,
} from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message api.admin.AdminAccountGoogleToken
 */
export class AdminAccountGoogleToken extends Message<AdminAccountGoogleToken> {
  /**
   * @generated from field: string access_token = 1;
   */
  accessToken = "";

  /**
   * @generated from field: string refresh_token = 2;
   */
  refreshToken = "";

  /**
   * @generated from field: google.protobuf.Timestamp expired_at = 3;
   */
  expiredAt?: Timestamp;

  constructor(data?: PartialMessage<AdminAccountGoogleToken>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.admin.AdminAccountGoogleToken";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: "access_token",
      kind: "scalar",
      T: 9 /* ScalarType.STRING */,
    },
    {
      no: 2,
      name: "refresh_token",
      kind: "scalar",
      T: 9 /* ScalarType.STRING */,
    },
    { no: 3, name: "expired_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>
  ): AdminAccountGoogleToken {
    return new AdminAccountGoogleToken().fromBinary(bytes, options);
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>
  ): AdminAccountGoogleToken {
    return new AdminAccountGoogleToken().fromJson(jsonValue, options);
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>
  ): AdminAccountGoogleToken {
    return new AdminAccountGoogleToken().fromJsonString(jsonString, options);
  }

  static equals(
    a:
      | AdminAccountGoogleToken
      | PlainMessage<AdminAccountGoogleToken>
      | undefined,
    b:
      | AdminAccountGoogleToken
      | PlainMessage<AdminAccountGoogleToken>
      | undefined
  ): boolean {
    return proto3.util.equals(AdminAccountGoogleToken, a, b);
  }
}