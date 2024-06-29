import { ApiClient } from "~/pkg/infrastructure/api_client";
import type { HealthCheckRequest } from "~/pkg/domain/model/health/health_check_request.gen";
import type { HealthCheckResponse } from "~/pkg/domain/model/health/health_check_response.gen";

/**
 * ヘルスチェック
 */
export class HealthService {
  private apiClient: ApiClient;

  constructor(apiClient: ApiClient) {
    this.apiClient = apiClient;
  }

  /**
   * check ヘルスチェックを行う
   * @param {HealthCheckRequest} req ヘルスチェックリクエスト
   * @returns {Promise<HealthCheckResponse>} - ヘルスチェックレスポンス
   */
  async check(req: HealthCheckRequest): Promise<HealthCheckResponse> {
    const apiClient = new ApiClient();
    const config = useRuntimeConfig();

    return await apiClient.post(
      config.public.GcViewUrl + "/api/admin/health/check",
      req,
      ""
    );
  }
}
