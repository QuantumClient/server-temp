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

  modern: true,

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
    timeout: 1500,
    proxy: true
  },

  proxy: {
    '/api': process.env.API_URL
  },

  auth: {
    strategies: {
      local: {
        scheme: 'refresh',
        endpoints: {
          login: { url: '/api/v1/auth/login', method: 'put' },
          refresh: { url: '/api/v1/auth/refresh', method: 'post'},
          user: { url: '/api/v1/auth/token', method: 'get'},
          logout: false
        },
        token: {
          type: 'Bearer',
          name: 'Authorization',
          property: 'access_token',
          maxAge: 60 * 60 * 2
        },
        refreshToken: {
          property: 'refresh_token',
          data: 'refresh_token',
          maxAge: 60 * 60 * 24 * 90
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
    publicPath: '/cdn/',
    filenames: {
      app: ({ isDev, isModern }) => isDev ? `[name]${isModern ? '.modern' : ''}.js` : `js/[contenthash:7]${isModern ? '.modern' : ''}.js`,
      chunk: ({ isDev, isModern }) => isDev ? `[name]${isModern ? '.modern' : ''}.js` : `js/[contenthash:7]${isModern ? '.modern' : ''}.js`,
      css: ({ isDev }) => isDev ? '[name].css' : 'css/[contenthash:7].css',
      img: ({ isDev }) => isDev ? '[path][name].[ext]' : 'img/[name].[contenthash:7].[ext]',
      font: ({ isDev }) => isDev ? '[path][name].[ext]' : 'fonts/[name].[contenthash:7].[ext]',
      video: ({ isDev }) => isDev ? '[path][name].[ext]' : 'videos/[name].[contenthash:7].[ext]'
    }

  },

  toast: {
    position: 'top-center',

  }
}
