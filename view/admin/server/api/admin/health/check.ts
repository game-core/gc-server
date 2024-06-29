import { createPromiseClient } from "@connectrpc/connect";
import { createGrpcTransport } from "@connectrpc/connect-node";
import { AdminHealthEnum } from "~/pkg/domain/model/health/adminHealth/admin_health_enum.gen";
import { CommonHealthEnum } from "~/pkg/domain/model/health/commonHealth/common_health_enum.gen";
import { MasterHealthEnum } from "~/pkg/domain/model/health/masterHealth/master_health_enum.gen";
import type { HealthCheckRequest } from "~/pkg/domain/model/health/health_check_request.gen";
import type { HealthCheckResponse } from "~/pkg/domain/model/health/health_check_response.gen";
import { Health } from "~/server/proto/health/health_handler_connect";
import { HealthCheckRequest as grpcHealthCheckRequest } from "~/server/proto/health/health_check_request_pb";
import { HealthCheckResponse as grpcHealthCheckResponse } from "~/server/proto/health/health_check_response_pb";

export default defineEventHandler(
  async (event: Promise<HealthCheckRequest>): Promise<HealthCheckResponse> => {
    const config = useRuntimeConfig();
    const client = createPromiseClient(
      Health,
      createGrpcTransport({
        httpVersion: "2",
        baseUrl: config.public.GcServerUrl,
      })
    );

    const body: HealthCheckRequest = await readBody(event);
    const req: grpcHealthCheckRequest = {
      healthId: body.healthId,
    };
    const res: grpcHealthCheckResponse = await client.check(req);

    return {
      adminHealth: {
        healthId: Number(res.adminHealth.healthId),
        name: res.adminHealth.name,
        adminHealthEnum: AdminHealthEnum[res.adminHealth.adminHealthEnum],
      },
      commonHealth: {
        healthId: Number(res.commonHealth.healthId),
        name: res.commonHealth.name,
        commonHealthEnum: CommonHealthEnum[res.commonHealth.commonHealthEnum],
      },
      masterHealth: {
        healthId: Number(res.masterHealth.healthId),
        name: res.masterHealth.name,
        masterHealthEnum: MasterHealthEnum[res.masterHealth.masterHealthEnum],
      },
    };
  }
);
