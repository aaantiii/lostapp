import LoginIndex from './login/Index'
import LoginSuccess from './login/Success'
import LoginFailed from './login/Failed'

export default {
  login: {
    Index: LoginIndex,
    Success: LoginSuccess,
    Failed: LoginFailed,
  },
} as const
