import { ApiClient } from "~/pkg/infrastructure/api_client";
import type {AccountGetGoogleTokenRequest} from "~/pkg/domain/model/account/account_get_google_token_request.gen";
import type {AccountGetGoogleTokenResponse} from "~/pkg/domain/model/account/account_get_google_token_response.gen";

export class AccountService {
    private apiClient: ApiClient;

    constructor(apiClient: ApiClient) {
        this.apiClient = apiClient;
    }

    async getGoogleLoginToken(req: AccountGetGoogleTokenRequest): Promise<AccountGetGoogleTokenResponse> {
        const apiClient = new ApiClient();
        const config = useRuntimeConfig();

        return await apiClient.post(
            config.public.GcViewUrl + "/api/admin/account/get_google_login_token",
            req,
            ""
        );
    }
}
