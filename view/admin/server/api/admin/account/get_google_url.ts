import { createPromiseClient } from "@connectrpc/connect";
import { createGrpcTransport } from "@connectrpc/connect-node";
import type { AccountGetGoogleUrlRequest } from "~/pkg/domain/model/account/account_get_google_url_request.gen";
import type { AccountGetGoogleUrlResponse } from "~/pkg/domain/model/account/account_get_google_url_response.gen";
import { Account } from "~/server/proto/account/account_handler_connect";
import { AccountGetGoogleUrlRequest as grpcAccountGetGoogleUrlRequest } from "~/server/proto/account/account_get_google_url_request_pb";
import { AccountGetGoogleUrlResponse as grpcAccountGetGoogleUrlResponse } from "~/server/proto/account/account_get_google_url_response_pb";

export default defineEventHandler(
  async (
    event: Promise<AccountGetGoogleUrlRequest>
  ): Promise<AccountGetGoogleUrlResponse> => {
    const config = useRuntimeConfig();
    const client = createPromiseClient(
      Account,
      createGrpcTransport({
        httpVersion: "2",
        baseUrl: config.public.GcServerUrl,
      })
    );

    const req: grpcAccountGetGoogleUrlRequest = {};
    const res: grpcAccountGetGoogleUrlResponse = await client.getGoogleUrl(req);

    return {
      adminAccountGoogleUrl: {
        url: res.adminAccountGoogleUrl.url,
      },
    };
  }
);
