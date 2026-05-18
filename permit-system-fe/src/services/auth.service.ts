import api from "../lib/axios"

import type {
  LoginRequest,
  LoginResponse,
} from "../types/auth"

export const login = async (
  payload: LoginRequest
): Promise<LoginResponse> => {

  const response = await api.post(
    "/auth/login",
    payload
  )

  const data: LoginResponse =
    response.data

  localStorage.setItem(
    "access_token",
    data.access_token
  )

  localStorage.setItem(
    "refresh_token",
    data.refresh_token
  )

  localStorage.setItem(
    "user",
    JSON.stringify(data.user)
  )

  return data
}

export const logout = async () => {

  const refreshToken =
    localStorage.getItem("refresh_token")

  await api.post("/auth/logout", {
    refresh_token: refreshToken,
  })

  localStorage.removeItem(
    "access_token"
  )

  localStorage.removeItem(
    "refresh_token"
  )

  localStorage.removeItem("user")

  window.location.href = "/login"
}