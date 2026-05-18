package dto

type DashboardResponse struct {
	TotalDocuments       int64 `json:"total_documents"`
	ActiveDocuments      int64 `json:"active_documents"`
	ExpiredDocuments     int64 `json:"expired_documents"`
	NotExtendedDocuments int64 `json:"not_extended_documents"`
	PendingApprovals     int64 `json:"pending_approvals"`
	ApprovedDocuments    int64 `json:"approved_documents"`
	RejectedDocuments    int64 `json:"rejected_documents"`
}
