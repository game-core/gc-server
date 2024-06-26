// アカウント

// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file account/account_handler.proto (package api.admin, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AccountGetGoogleUrlRequest } from "./account_get_google_url_request_pb.js";
import { AccountGetGoogleUrlResponse } from "./account_get_google_url_response_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { AccountGetGoogleTokenRequest } from "./account_get_google_token_request_pb.js";
import { AccountGetGoogleTokenResponse } from "./account_get_google_token_response_pb.js";
import { AccountRefreshGoogleTokenRequest } from "./account_refresh_google_token_request_pb.js";
import { AccountRefreshGoogleTokenResponse } from "./account_refresh_google_token_response_pb.js";

/**
 * @generated from service api.admin.Account
 */
export const Account = {
  typeName: "api.admin.Account",
  methods: {
    /**
     * @generated from rpc api.admin.Account.GetGoogleUrl
     */
    getGoogleUrl: {
      name: "GetGoogleUrl",
      I: AccountGetGoogleUrlRequest,
      O: AccountGetGoogleUrlResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.admin.Account.GetGoogleToken
     */
    getGoogleToken: {
      name: "GetGoogleToken",
      I: AccountGetGoogleTokenRequest,
      O: AccountGetGoogleTokenResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc api.admin.Account.RefreshGoogleToken
     */
    refreshGoogleToken: {
      name: "RefreshGoogleToken",
      I: AccountRefreshGoogleTokenRequest,
      O: AccountRefreshGoogleTokenResponse,
      kind: MethodKind.Unary,
    },
  },
} as const;
