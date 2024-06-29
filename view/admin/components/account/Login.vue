<template></template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { ApiClient } from "~/pkg/infrastructure/api_client";
import { AccountService } from "~/pkg/domain/model/account/account_service";
import type { AccountGetGoogleUrlRequest } from "~/pkg/domain/model/account/account_get_google_url_request.gen";
import type { AccountGetGoogleUrlResponse } from "~/pkg/domain/model/account/account_get_google_url_response.gen";

const apiClient = new ApiClient();
const accountService = new AccountService(apiClient);

const route = useRoute();
const queryParams = route.query;
const res = ref<AccountGetGoogleUrlResponse | null>(null);

const sendRequest = async (): Promise<void> => {
  try {
    const code: string = queryParams.code as string;
    const req: AccountGetGoogleUrlRequest = { code: code };
    res.value = await accountService.getGoogleUrl(req);

    window.location.href = res.value.adminAccountGoogleUrl.url;
  } catch (error) {
    console.error("accountService.getGoogleUrl:", error);
  }
};

onMounted(() => {
  sendRequest();
});
</script>
