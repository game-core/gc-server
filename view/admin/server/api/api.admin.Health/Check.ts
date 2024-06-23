import { ApiClient } from "~/pkg/infrastructure/ApiClient";
import type { HealthCheckResponse } from "~/pkg/domain/model/health/HealthCheckResponse.gen";

export default defineEventHandler(
  async (event): Promise<HealthCheckResponse> => {
    const apiClient = new ApiClient();
    const config = useRuntimeConfig();
    const body = await readBody(event);

    return await apiClient.post(
      config.public.GcServerUrl + "/api.admin.Health/Check",
      body,
      ""
    );
  }
);
