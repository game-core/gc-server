// ヘルスチェックタイプ

// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file health/adminHealth/admin_health_enum.proto (package api.admin, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum api.admin.AdminHealthEnum
 */
export enum AdminHealthEnum {
  /**
   * @generated from enum value: AdminNone = 0;
   */
  AdminNone = 0,

  /**
   * @generated from enum value: AdminSuccess = 1;
   */
  AdminSuccess = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(AdminHealthEnum)
proto3.util.setEnumType(AdminHealthEnum, "api.admin.AdminHealthEnum", [
  { no: 0, name: "AdminNone" },
  { no: 1, name: "AdminSuccess" },
]);