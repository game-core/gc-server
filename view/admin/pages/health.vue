<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <Check />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Check from "~/components/health/Check.vue";

export default {
  components: {
    Check,
  },
};
</script>
