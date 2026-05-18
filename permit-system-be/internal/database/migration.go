package database

import (
	"fmt"
	"permit-license/config"
	"permit-license/internal/model"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&model.Unit{},
		&model.Role{},
		&model.User{},
		&model.UserRole{},
		&model.MasterDocument{},
		&model.ApprovalStatus{},
		&model.PermitLicense{},
		&model.ApprovalHistory{},
		&model.RefreshToken{},
	)
	if err != nil {
		panic("Failed to migrate database!")
	}

	fmt.Println("Database migrated successfully.")
}
