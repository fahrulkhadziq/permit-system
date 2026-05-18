import type { UserLoginResponse }
from "../types/auth"

export const getUser =
(): UserLoginResponse | null => {

  const user =
    localStorage.getItem("user")

  if (!user) return null

  return JSON.parse(user)
}

export const getAccessToken =
(): string | null => {

  return localStorage.getItem(
    "access_token"
  )
}

export const isAuthenticated =
(): boolean => {

  return !!localStorage.getItem(
    "access_token"
  )
}