<template>
  <v-list>
    <ProfileList />
    <SearchList />
  </v-list>
</template>

<script lang="ts">
import ProfileList from "~/components/layout/list/ProfileList.vue";
import SearchList from "~/components/layout/list/SearchList.vue";

export default {
  name: "SidebarList",
  components: {
    ProfileList,
    SearchList,
  },
};
</script>

<style scoped></style>
