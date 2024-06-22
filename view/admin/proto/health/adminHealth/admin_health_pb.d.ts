// ヘルスチェック

// @generated by protoc-gen-es v1.10.0
// @generated from file health/adminHealth/admin_health.proto (package api.admin, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { AdminHealthType } from "./admin_health_type_enum_pb.js";

/**
 * @generated from message api.admin.AdminHealth
 */
export declare class AdminHealth extends Message<AdminHealth> {
  /**
   * @generated from field: int64 health_id = 1;
   */
  healthId: bigint;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: api.admin.AdminHealthType admin_health_type = 3;
   */
  adminHealthType: AdminHealthType;

  constructor(data?: PartialMessage<AdminHealth>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "api.admin.AdminHealth";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdminHealth;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdminHealth;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdminHealth;

  static equals(a: AdminHealth | PlainMessage<AdminHealth> | undefined, b: AdminHealth | PlainMessage<AdminHealth> | undefined): boolean;
}

