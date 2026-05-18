package dto

type RegisterRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	UnitID   string `json:"unit_id" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	AccessToken  string            `json:"access_token"`
	RefreshToken string            `json:"refresh_token"`
	User         LoginUserResponse `json:"user"`
}

type LoginUserResponse struct {
	ID       string  `json:"id"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	UnitID   *string `json:"unit_id"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}
