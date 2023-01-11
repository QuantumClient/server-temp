import Vue from 'vue'
import { Plugin } from '@nuxt/types'
import validator from "~/utils/validator";

Vue.prototype.$validator = validator;

declare module '@nuxt/types' {
  interface NuxtAppOptions {
    $download(data: any, name: string): void
  }
  interface Context {
    $download(data: any, name: string): void
  }
}

declare module 'vue/types/vue' {
  interface Vue {
    $download(data: any, name: string): void
  }
}

const mixin: Plugin = (context, inject) => {
  inject('download', (data: any, name: string) => {
    const jsonData = JSON.stringify(data)
    const url = window.URL.createObjectURL(new Blob( [ btoa(jsonData) ] , { type: 'application/octet-stream' }))
    const link = document.createElement('a')

    link.href = url
    link.setAttribute( 'href', url );
    link.setAttribute('download', name);

    document.body.appendChild(link)

    link.click()
    link.remove()
    window.URL.revokeObjectURL(url)
  });
}

export default mixin
