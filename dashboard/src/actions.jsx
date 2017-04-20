export const SIDE_BAR_TOGGLE = 'side-bar.toggle'
export const USERS_SIGN_IN = "users.sign-in"
export const USERS_SIGN_OUT = "users.sign-out"

export const signIn = (token) => {
  return {
    type: USERS_SIGN_IN,
    token
  }
}

export const signOut = () => {
  return {
    type: USERS_SIGN_OUT
  }
}

export const toggleSideBar = () => {
  return {
    type: SIDE_BAR_TOGGLE    
  }
}
