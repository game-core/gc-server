import {ApiClient} from "~/pkg/infrastructure/ApiClient"
import type {HealthCheckRequest} from "~/pkg/domain/health/HealthCheckRequest";
import type {HealthCheckResponse} from "~/pkg/domain/health/HealthCheckResponse";

export class HealthService {
    private apiClient: ApiClient

    constructor(apiClient: ApiClient) {
        this.apiClient = apiClient
    }

    async check(req: HealthCheckRequest): Promise<HealthCheckResponse> {
        const apiClient = new ApiClient();
        const config = useRuntimeConfig();

        return await apiClient.post(config.public.GcViewUrl + "/api/api.admin.Health/Check", req, "")
    }
}
