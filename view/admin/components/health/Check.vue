<template>
  <v-card class="base_card" flat>
    <v-card-title>Health Check</v-card-title>
    <v-card-item>
      <button @click="sendPostRequest">Send HealthCheck</button>
      <div v-if="response">
        <p>Response: {{ response }}</p>
      </div>
    </v-card-item>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { ApiClient } from "~/pkg/infrastructure/api_client";
import { HealthService } from "~/pkg/domain/model/health/health_service";
import type { HealthCheckRequest } from "~/pkg/domain/model/health/health_check_request.gen";
import type { HealthCheckResponse } from "~/pkg/domain/model/health/health_check_response.gen";

const response = ref<HealthCheckResponse | null>(null);

const apiClient = new ApiClient();
const healthService = new HealthService(apiClient);

const sendPostRequest = async (): Promise<void> => {
  try {
    const req: HealthCheckRequest = {
      healthId: 1,
    };

    response.value = await healthService.check(req);
  } catch (error) {
    console.error("Health check request failed:", error);
  }
};
</script>

<style scoped>
.base_card {
  background: #ffffff;
}
</style>
