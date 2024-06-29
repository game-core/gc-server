import {createPromiseClient} from "@connectrpc/connect";
import {createGrpcTransport} from "@connectrpc/connect-node";
import type {AccountGetGoogleTokenRequest} from "~/pkg/domain/model/account/account_get_google_token_request.gen";
import type {AccountGetGoogleTokenResponse} from "~/pkg/domain/model/account/account_get_google_token_response.gen";
import {Account} from "~/server/proto/account/account_handler_connect";
import {AccountGetGoogleTokenRequest as grpcAccountGetGoogleTokenRequest} from "~/server/proto/account/account_get_google_token_request_pb";
import {AccountGetGoogleTokenResponse as grpcAccountGetGoogleTokenResponse} from "~/server/proto/account/account_get_google_token_response_pb";

export default defineEventHandler(
    async (event: Promise<AccountGetGoogleTokenRequest>): Promise<AccountGetGoogleTokenResponse> => {
        const config = useRuntimeConfig();
        const client = createPromiseClient(
            Account,
            createGrpcTransport({
                httpVersion: "2",
                baseUrl: config.public.GcServerUrl,
            })
        );

        const body: AccountGetGoogleTokenRequest = await readBody(event);
        const req: grpcAccountGetGoogleTokenRequest = {
            code: body.code
        };
        const res: grpcAccountGetGoogleTokenResponse = await client.getGoogleToken(req);

        return {
            adminAccountGoogleToken: {
                accessToken: res.adminAccountGoogleToken.accessToken,
                refreshToken: res.adminAccountGoogleToken.refreshToken,
                expiredAt: res.adminAccountGoogleToken.expiredAt,
            }
        };
    }
);
