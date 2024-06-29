import { ApiClient } from "~/pkg/infrastructure/api_client";
import type { AccountGetGoogleTokenRequest } from "~/pkg/domain/model/account/account_get_google_token_request.gen";
import type { AccountGetGoogleTokenResponse } from "~/pkg/domain/model/account/account_get_google_token_response.gen";
import type { AccountGetGoogleUrlRequest } from "~/pkg/domain/model/account/account_get_google_url_request.gen";
import type { AccountGetGoogleUrlResponse } from "~/pkg/domain/model/account/account_get_google_url_response.gen";

export class AccountService {
  private apiClient: ApiClient;

  constructor(apiClient: ApiClient) {
    this.apiClient = apiClient;
  }

  /**
   * getGoogleUrl GoogleURLを取得する
   * @param {AccountGetGoogleTokenRequest} req GoogleURL取得リクエスト
   * @returns {Promise<AccountGetGoogleTokenResponse>} - GoogleURL取得レスポンス
   */
  async getGoogleUrl(
    req: AccountGetGoogleUrlRequest
  ): Promise<AccountGetGoogleUrlResponse> {
    const apiClient = new ApiClient();
    const config = useRuntimeConfig();

    return await apiClient.post(
      config.public.GcViewUrl + "/api/admin/account/get_google_url",
      req,
      ""
    );
  }

  /**
   * getGoogleToken Googleトークンを取得する
   * @param {AccountGetGoogleTokenRequest} req Googleトークン取得リクエスト
   * @returns {Promise<AccountGetGoogleTokenResponse>} - Googleトークン取得レスポンス
   */
  async getGoogleToken(
    req: AccountGetGoogleTokenRequest
  ): Promise<AccountGetGoogleTokenResponse> {
    const apiClient = new ApiClient();
    const config = useRuntimeConfig();

    return await apiClient.post(
      config.public.GcViewUrl + "/api/admin/account/get_google_token",
      req,
      ""
    );
  }
}
