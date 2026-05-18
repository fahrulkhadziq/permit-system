package model

type ApprovalHistory struct {
	BaseModel

	PermitLicenseID string `gorm:"index"`
	ApproverID      string `gorm:"index"`
	StatusID        string `gorm:"index"`

	Notes string

	PermitLicense PermitLicense  `gorm:"foreignKey:PermitLicenseID"`
	Approver      User           `gorm:"foreignKey:ApproverID"`
	Status        ApprovalStatus `gorm:"foreignKey:StatusID"`
}
