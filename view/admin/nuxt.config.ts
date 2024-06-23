// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,
  target: "server",
  runtimeConfig: {
    public: {
      GcServerUrl: process.env.GC_SERVER_URL,
      GcViewUrl: process.env.GC_VIEW_URL,
    },
  },
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
});
