// アカウントGoogleToken情報取得レスポンス

// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file account/account_get_google_token_info_response.proto (package api.admin, syntax proto3)
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
import { AdminAccountGoogleTokenInfo } from "./adminAccountGoogleTokenInfo/admin_account_google_token_info_pb.js";

/**
 * @generated from message api.admin.AccountGetGoogleTokenInfoResponse
 */
export class AccountGetGoogleTokenInfoResponse extends Message<AccountGetGoogleTokenInfoResponse> {
  /**
   * @generated from field: optional api.admin.AdminAccountGoogleTokenInfo admin_account_google_token_info = 1;
   */
  adminAccountGoogleTokenInfo?: AdminAccountGoogleTokenInfo;

  constructor(data?: PartialMessage<AccountGetGoogleTokenInfoResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.admin.AccountGetGoogleTokenInfoResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: "admin_account_google_token_info",
      kind: "message",
      T: AdminAccountGoogleTokenInfo,
      opt: true,
    },
  ]);

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>
  ): AccountGetGoogleTokenInfoResponse {
    return new AccountGetGoogleTokenInfoResponse().fromBinary(bytes, options);
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>
  ): AccountGetGoogleTokenInfoResponse {
    return new AccountGetGoogleTokenInfoResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>
  ): AccountGetGoogleTokenInfoResponse {
    return new AccountGetGoogleTokenInfoResponse().fromJsonString(
      jsonString,
      options
    );
  }

  static equals(
    a:
      | AccountGetGoogleTokenInfoResponse
      | PlainMessage<AccountGetGoogleTokenInfoResponse>
      | undefined,
    b:
      | AccountGetGoogleTokenInfoResponse
      | PlainMessage<AccountGetGoogleTokenInfoResponse>
      | undefined
  ): boolean {
    return proto3.util.equals(AccountGetGoogleTokenInfoResponse, a, b);
  }
}
