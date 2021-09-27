import { NuxtConfig } from '@nuxt/types';

export default <NuxtConfig>{
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'Quantum',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Modern Minecraft Clients' },
      { hid: 'theme-color', name: 'theme-color', content: '#437373'},
      {
        hid: 'og:title',
        property: 'og:title',
        content: 'Quantum'
      },
      {
        hid: 'og:description',
        property: 'og:description',
        content: 'Modern Minecraft Clients'
      },
      {
        hid: 'og:image',
        property: 'og:image',
        content: '/favicon.png'
      },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.png' }
    ]
  },

  loading: {
    color: '#565AA6',
    height: '8px'
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    '~/assets/css/main.css'
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    '~/plugins/mixin'
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    '@nuxtjs/tailwindcss',
    '@nuxtjs/svg'
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/auth-next',
    '@nuxtjs/toast',
  ],

  axios: {
    // extra config e.g
    timeout: 1500,
    proxy: true
  },

  auth: {
    strategies: {
      local: {
        endpoints: {
          login: { url: '/api/v1/auth/login', method: 'put', propertyName: 'token' },
          user: { url: '/api/v1/auth/token', method: 'get'}
        },
        token: {
          type: 'Bearer',
          name: 'Authorization',
        },
        user: {
          property: false,
          autoFetch: true
        },
      }
    },
    redirect: {
      login: '/auth/login',
      logout: '/',
      home: '/'
    }
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    extractCSS: true,
    publicPath: '/qcn/'
  },

  toast: {
    position: 'top-center',

  }
}
