import { Middleware } from '@nuxt/types'

const adminMiddleware: Middleware = (context) => {
  if (!context.$auth.loggedIn) {
    return context.redirect({ path: '/auth/login', query: { to: context.route.fullPath } })
  }
  if (!context.$auth.user.admin) {
    return context.redirect('/dashboard')
  }
}

export default adminMiddleware
