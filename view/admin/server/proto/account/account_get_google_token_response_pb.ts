// アカウントGoogleToken取得レスポンス

// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file account/account_get_google_token_response.proto (package api.admin, syntax proto3)
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
import { Message, proto3 } from "@bufbuild/protobuf";
import { AdminAccountGoogleToken } from "./adminAccountGoogleToken/admin_account_google_token_pb.js";

/**
 * @generated from message api.admin.AccountGetGoogleTokenResponse
 */
export class AccountGetGoogleTokenResponse extends Message<AccountGetGoogleTokenResponse> {
  /**
   * @generated from field: optional api.admin.AdminAccountGoogleToken admin_account_google_token = 1;
   */
  adminAccountGoogleToken?: AdminAccountGoogleToken;

  constructor(data?: PartialMessage<AccountGetGoogleTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.admin.AccountGetGoogleTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: "admin_account_google_token",
      kind: "message",
      T: AdminAccountGoogleToken,
      opt: true,
    },
  ]);

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>
  ): AccountGetGoogleTokenResponse {
    return new AccountGetGoogleTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>
  ): AccountGetGoogleTokenResponse {
    return new AccountGetGoogleTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>
  ): AccountGetGoogleTokenResponse {
    return new AccountGetGoogleTokenResponse().fromJsonString(
      jsonString,
      options
    );
  }

  static equals(
    a:
      | AccountGetGoogleTokenResponse
      | PlainMessage<AccountGetGoogleTokenResponse>
      | undefined,
    b:
      | AccountGetGoogleTokenResponse
      | PlainMessage<AccountGetGoogleTokenResponse>
      | undefined
  ): boolean {
    return proto3.util.equals(AccountGetGoogleTokenResponse, a, b);
  }
}
