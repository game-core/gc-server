<template>
  <NuxtLayout name="redirect">
    <v-container>
      <v-row>
        <v-col cols="12">
          <Redirect/>
        </v-col>
      </v-row>
    </v-container>
  </NuxtLayout>
</template>

<script lang="ts">
import Redirect from "~/components/redirect/Redirect.vue";

export default {
  components: {
    Redirect,
  },
};
</script>
