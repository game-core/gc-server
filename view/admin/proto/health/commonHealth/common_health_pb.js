// ヘルスチェック

// @generated by protoc-gen-es v1.10.0
// @generated from file health/commonHealth/common_health.proto (package api.admin, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { CommonHealthType } from "./common_health_type_enum_pb.js";

/**
 * @generated from message api.admin.CommonHealth
 */
export const CommonHealth = /*@__PURE__*/ proto3.makeMessageType(
  "api.admin.CommonHealth",
  () => [
    { no: 1, name: "health_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "common_health_type", kind: "enum", T: proto3.getEnumType(CommonHealthType) },
  ],
);

