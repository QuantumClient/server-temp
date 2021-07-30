export default {
  username: (v: string) => {
    return (!v || (/^(?=.{3,50}$)[A-Za-z0-9]+(?:[-_.][A-Za-z0-9]+)*$/.test(v))) || 'Username should be between 3 and 50 characters and can only contains alphanumeric characters, underscore and dot.';
  },
  password: (v: string) => {
   return (!v || (v && v.length > 6 && v.length < 100)) || 'Password should be between 6 and 100 characters.'
  },
  passwordMatch: (pass: string, v: string) => {
    return (!v || v === pass) || 'Repeat Password does not match Password.'
  },
}
