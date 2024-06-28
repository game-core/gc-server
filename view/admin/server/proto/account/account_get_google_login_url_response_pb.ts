// アカウントGoogleログインURL取得レスポンス

// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file account/account_get_google_login_url_response.proto (package api.admin, syntax proto3)
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

/**
 * @generated from message api.admin.AccountGetGoogleLoginUrlResponse
 */
export class AccountGetGoogleLoginUrlResponse extends Message<AccountGetGoogleLoginUrlResponse> {
  /**
   * @generated from field: string url = 1;
   */
  url = "";

  constructor(data?: PartialMessage<AccountGetGoogleLoginUrlResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.admin.AccountGetGoogleLoginUrlResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>
  ): AccountGetGoogleLoginUrlResponse {
    return new AccountGetGoogleLoginUrlResponse().fromBinary(bytes, options);
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>
  ): AccountGetGoogleLoginUrlResponse {
    return new AccountGetGoogleLoginUrlResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>
  ): AccountGetGoogleLoginUrlResponse {
    return new AccountGetGoogleLoginUrlResponse().fromJsonString(
      jsonString,
      options
    );
  }

  static equals(
    a:
      | AccountGetGoogleLoginUrlResponse
      | PlainMessage<AccountGetGoogleLoginUrlResponse>
      | undefined,
    b:
      | AccountGetGoogleLoginUrlResponse
      | PlainMessage<AccountGetGoogleLoginUrlResponse>
      | undefined
  ): boolean {
    return proto3.util.equals(AccountGetGoogleLoginUrlResponse, a, b);
  }
}
