<template>
  <NuxtLayout name="index">
    <v-container>
      <v-row>
        <v-col cols="12">
          <Redirect />
        </v-col>
      </v-row>
    </v-container>
  </NuxtLayout>
</template>

<script lang="ts" setup>
import Redirect from "~/components/account/Redirect.vue";
</script>
