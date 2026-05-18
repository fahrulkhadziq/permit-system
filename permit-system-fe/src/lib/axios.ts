import axios from "axios"

const api = axios.create({
  baseURL: "http://localhost:8080/api",
})

api.interceptors.request.use((config) => {

  const token =
    localStorage.getItem("access_token")

  if (token) {
    config.headers.Authorization =
      `Bearer ${token}`
  }

  return config
})

api.interceptors.response.use(
  (response) => response,

  async (error) => {

    const originalRequest = error.config

    if (
      error.response?.status === 401 &&
      !originalRequest._retry
    ) {

      originalRequest._retry = true

      try {

        const refreshToken =
          localStorage.getItem("refresh_token")

        const response = await axios.post(
          "http://localhost:8080/api/auth/refresh",
          {
            refresh_token: refreshToken,
          }
        )

        const newAccessToken =
          response.data.access_token

        localStorage.setItem(
          "access_token",
          newAccessToken
        )

        originalRequest.headers.Authorization =
          `Bearer ${newAccessToken}`

        return api(originalRequest)

      } catch (err) {

        localStorage.clear()

        window.location.href = "/login"

        return Promise.reject(err)
      }
    }

    return Promise.reject(error)
  }
)

export default api