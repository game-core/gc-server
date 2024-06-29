<template>
  <v-card class="base_card" flat>
    <v-card-title>Redirect</v-card-title>
    <v-card-item>
      <div>アカウント情報取得中...</div>
    </v-card-item>
  </v-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { ApiClient } from "~/pkg/infrastructure/api_client";
import { AccountService } from "~/pkg/domain/model/account/account_service";
import type { AccountGetGoogleTokenRequest } from "~/pkg/domain/model/account/account_get_google_token_request.gen";
import type { AccountGetGoogleTokenResponse } from "~/pkg/domain/model/account/account_get_google_token_response.gen";
import { useAuthStore } from "~/store/account/auth";

const apiClient = new ApiClient();
const accountService = new AccountService(apiClient);

const accessToken = useCookie("accessToken", { secure: true, sameSite: true });
const refreshToken = useCookie("refreshToken", {
  secure: true,
  sameSite: true,
});
const authStore = useAuthStore();

const route = useRoute();
const router = useRouter();
const queryParams = route.query;
const res = ref<AccountGetGoogleTokenResponse | null>(null);

const sendRequest = async (): Promise<void> => {
  try {
    const code: string = queryParams.code as string;
    const req: AccountGetGoogleTokenRequest = { code: code };
    res.value = await accountService.getGoogleToken(req);

    accessToken.value = res.value.adminAccountGoogleToken.accessToken;
    refreshToken.value = res.value.adminAccountGoogleToken.refreshToken;
    authStore.setIsLogin(true);

    await router.push("/");
  } catch (error) {
    console.error("accountService.getGoogleLoginToken:", error);
  }
};

onMounted(() => {
  sendRequest();
});
</script>

<style scoped>
.base_card {
  background: #ffffff;
}
</style>
