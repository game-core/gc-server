<template>
  <NuxtLayout name="index">
    <v-container>
      <v-row>
        <v-col cols="12">
          <Login />
        </v-col>
      </v-row>
    </v-container>
  </NuxtLayout>
</template>

<script lang="ts" setup>
import Login from "~/components/account/Login.vue";
</script>
