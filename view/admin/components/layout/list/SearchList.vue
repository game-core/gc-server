<template>
  <v-list-group color="primary" value="Search">
    <template v-slot:activator="{ props }">
      <v-list-item
          v-bind="props"
          title="Search"
          rounded="xl"
          prepend-icon="mdi-magnify"
      ></v-list-item>
    </template>

    <v-list-item
        key="creator"
        color="primary"
        title="creator"
        prepend-icon="mdi-account-group"
        to="/search/searchUser"
    >
    </v-list-item>

    <v-list-item
        key="image"
        color="primary"
        title="image"
        prepend-icon="mdi-image-multiple"
        to="/search/searchRoom"
    >
    </v-list-item>

  </v-list-group>
</template>

<script lang="ts">
export default {
  name: "SearchList",
};
</script>

<style scoped></style>
