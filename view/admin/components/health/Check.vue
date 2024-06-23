<template>
  <div>
    <header class="app-header">
      <h1>Health</h1>
    </header>
    <button @click="sendPostRequest">Send HealthCheck</button>
    <div v-if="response">
      <p>Response: {{ response }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {ApiClient} from "~/pkg/infrastructure/ApiClient";
import {HealthService} from "~/pkg/domain/model/health/HealthService";
import type {HealthCheckRequest} from "~/pkg/domain/model/health/HealthCheckRequest.gen";
import type {HealthCheckResponse} from "~/pkg/domain/model/health/HealthCheckResponse.gen";

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
.app-header {
  text-align: center;
}
</style>
