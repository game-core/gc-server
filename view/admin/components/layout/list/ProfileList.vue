<template>
  <v-list-item key="profile" value="profile" color="primary" rounded="xl" to="">
    <template v-slot:prepend>
      <v-icon icon="mdi-account"></v-icon>
    </template>
    <v-list-item-title>{{ "Profile" }}</v-list-item-title>
  </v-list-item>
</template>

<script lang="ts">
export default {
  name: "ProfileList",
};
</script>

<style scoped></style>
