<template>
  <v-card class="base_card" flat>
    <v-card-title>Redirect</v-card-title>
    <v-card-item>
      <div>アカウント情報取得中...</div>
      <div>情報: {{ res }}</div>
    </v-card-item>
  </v-card>
</template>

<script setup lang="ts">
import {ref, onMounted} from "vue";
import {useRoute} from 'vue-router';
import {ApiClient} from "~/pkg/infrastructure/api_client";
import {AccountService} from "~/pkg/domain/model/account/account_service";
import type {AccountGetGoogleTokenRequest} from "~/pkg/domain/model/account/account_get_google_token_request.gen";
import type {AccountGetGoogleTokenResponse} from "~/pkg/domain/model/account/account_get_google_token_response.gen";

const route = useRoute();
const queryParams = route.query;

const res = ref<AccountGetGoogleTokenResponse | null>(null);

const apiClient = new ApiClient();
const accountService = new AccountService(apiClient);

const sendPostRequest = async (): Promise<void> => {
  try {
    const code = queryParams.code as string;
    const req: AccountGetGoogleTokenRequest = {
      code: code,
    };
    res.value = await accountService.getGoogleLoginToken(req);
  } catch (error) {
    console.error("Health check request failed:", error);
  }
};

onMounted(() => {
  sendPostRequest();
});

</script>

<style scoped>
.base_card {
  background: #ffffff;
}
</style>
