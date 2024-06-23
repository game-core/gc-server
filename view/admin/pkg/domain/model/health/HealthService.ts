import { ApiClient } from "~/pkg/infrastructure/ApiClient";
import type { HealthCheckRequest } from "~/pkg/domain/model/health/HealthCheckRequest.gen";
import type { HealthCheckResponse } from "~/pkg/domain/model/health/HealthCheckResponse.gen";

export class HealthService {
  private apiClient: ApiClient;

  constructor(apiClient: ApiClient) {
    this.apiClient = apiClient;
  }

  async check(req: HealthCheckRequest): Promise<HealthCheckResponse> {
    const apiClient = new ApiClient();
    const config = useRuntimeConfig();

    return await apiClient.post(
      config.public.GcViewUrl + "/api/api.admin.Health/Check",
      req,
      ""
    );
  }
}
