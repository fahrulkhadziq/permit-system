package model

import "time"

type PermitLicense struct {
	BaseModel

	MasterDocumentID string `gorm:"index"`
	UploadedBy       string `gorm:"index"`
	UnitID           string `gorm:"index"`

	CurrentStatusID string `gorm:"index"`

	RelatedPrevDocumentID *string `gorm:"index"`
	RelatedNextDocumentID *string `gorm:"index"`

	DocumentName string
	Description  string

	FileURL  string
	FileSize int64

	ExpiredAt  time.Time `gorm:"index"`
	ApprovedAt *time.Time

	RejectedReason string

	IsActive bool `gorm:"default:true"`
	IsExtend bool `gorm:"default:false"`

	MasterDocument MasterDocument `gorm:"foreignKey:MasterDocumentID"`
	User           User           `gorm:"foreignKey:UploadedBy"`
	Unit           Unit           `gorm:"foreignKey:UnitID"`

	CurrentStatus ApprovalStatus `gorm:"foreignKey:CurrentStatusID"`

	RelatedPrevDocument *PermitLicense `gorm:"foreignKey:RelatedPrevDocumentID"`

	RelatedNextDocument *PermitLicense `gorm:"foreignKey:RelatedNextDocumentID"`

	ApprovalHistories []ApprovalHistory `gorm:"foreignKey:PermitLicenseID"`
}
