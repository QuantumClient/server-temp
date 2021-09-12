import { Middleware } from '@nuxt/types'

const authMiddleware: Middleware = (context) => {
  if (!context.$auth.loggedIn) {
    return context.redirect({ path: '/auth/login', query: { to: context.route.fullPath } })
  }
}

export default authMiddleware
