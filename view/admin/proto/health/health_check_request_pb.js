// ヘルスチェックリクエスト

// @generated by protoc-gen-es v1.10.0
// @generated from file health/health_check_request.proto (package api.admin, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message api.admin.HealthCheckRequest
 */
export const HealthCheckRequest = /*@__PURE__*/ proto3.makeMessageType(
  "api.admin.HealthCheckRequest",
  () => [
    { no: 1, name: "health_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

