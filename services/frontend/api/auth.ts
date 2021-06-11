import instance from './instance'

export default {
  login (body: string) {
    return instance.put(
      'auth/login',
      body
    )
  },
  signup (body: string) {
    return instance.put(
      'auth/register',
      body
    )
  },
  user (token: string) {
    return instance.get(
      'auth/token',
      { headers: {"Authorization" : `Bearer ${token}`}}
    )
  }
}
