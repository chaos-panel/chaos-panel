// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: [
    '@nuxt/devtools',
    '@nuxt/test-utils',
    '@nuxt/image',
    '@nuxt/eslint',
    '@nuxt/icon',
    '@nuxt/ui',
    '@nuxt/content',
    '@nuxtjs/tailwindcss',
    '@nuxt/fonts',
    '@nuxt/scripts',
    '@pinia/nuxt',
    'pinia-plugin-persistedstate',
    '@nuxtjs/i18n',
    '@formkit/auto-animate',
    '@nuxtjs/color-mode',
    '@nuxtjs/device',
    'radix-vue',
    '@primevue/nuxt-module',
    '@vueuse/nuxt',
    'nuxt-particles'
  ]
})