// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,
  runtimeConfig: {
    public: {
      GcServerUrl: process.env.GC_SERVER_URL,
      GcViewUrl: process.env.GC_VIEW_URL,
    },
  },
  modules: ["@pinia/nuxt", "@pinia-plugin-persistedstate/nuxt"],
  components: [
    {
      path: "~/components/",
      pathPrefix: false,
    },
  ],
  css: ["vuetify/lib/styles/main.sass", "mdi/css/materialdesignicons.css"],
  build: {
    transpile: ["vuetify"],
  },
  vite: {
    define: {
      "process.env.DEBUG": false,
    },
  },
  vuetify: {
    customVariables: ["~/assets/styles/main.sass"],
    treeShake: true,
  },
});
