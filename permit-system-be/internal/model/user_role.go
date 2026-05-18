package model

type UserRole struct {
	BaseModel

	UserID string
	RoleID string

	User User
	Role Role
}
