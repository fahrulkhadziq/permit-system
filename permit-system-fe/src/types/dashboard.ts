export interface DashboardResponse {
  total_documents: number
  active_documents: number
  expired_documents: number
  not_extended_documents: number
  pending_approvals: number
  approved_documents: number
  rejected_documents: number
}