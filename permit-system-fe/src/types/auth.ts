export interface LoginRequest {
  email: string
  password: string
}

export interface UserLoginResponse {
  id: string
  full_name: string
  email: string
  role: string
  unit_id: string
}

export interface LoginResponse {
  access_token: string
  refresh_token: string
  user: UserLoginResponse
}

export interface RefreshTokenRequest {
  refresh_token: string
}