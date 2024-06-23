<template>
  <v-app class="main">
    <v-navigation-drawer v-model="drawer" class="px-6 py-4 scrollbar navbar">
      <SidebarList />
    </v-navigation-drawer>

    <v-app-bar app flat>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <div class="center_icon">
        <v-btn icon class="center_icon_size">
          <v-icon>mdi-vuejs</v-icon>
        </v-btn>
      </div>
    </v-app-bar>

    <v-main>
      <slot />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import SidebarList from "~/components/layout/SidebarList.vue";

export default {
  components: {
    SidebarList,
  },
  data() {
    return {
      drawer: true,
      items: [
        { title: "Timeline", icon: "mdi-forum-outline", to: "/" },
        { title: "Message", icon: "mdi-email-outline", to: "" },
      ],
    };
  },
  mounted() {},
  methods: {},
};
</script>

<style lang="scss" scoped>
.main {
  background: #eeebeb;
}

.center_icon {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.center_icon_size {
  font-size: 25px;
}

img.user_icon {
  width: 55px;
  height: 55px;
  object-fit: cover;
  border-radius: 50%;
  position: relative;
  left: 5px;
}

.scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.scrollbar ::-webkit-scrollbar {
  display: none;
}

.navbar {
  box-shadow: none !important;
}
</style>
