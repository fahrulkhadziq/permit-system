export interface UnitResponse {
  id: string
  name: string
}

export interface UserResponse {
  id: string
  full_name: string
  email: string
  unit: UnitResponse
}

export interface ApprovalStatusResponse {
  id: string
  code: string
  name: string
}

export interface MasterDocumentResponse {
  id: string
  code: string
  name: string
  description: string
}

export interface ApprovalHistoryResponse {
  id: string
  notes: string
  created_at: string

  approver: UserResponse

  status: ApprovalStatusResponse
}

export interface PermitLicenseDetailResponse {
  id: string

  document_name: string
  description: string

  file_url: string
  file_size: number

  expired_at: string
  approved_at?: string

  rejected_reason: string

  is_active: boolean

  master_document:
    MasterDocumentResponse

  user: UserResponse

  unit: UnitResponse

  current_status:
    ApprovalStatusResponse

  related_prev_document?: {
    id: string
    document_name: string
    expired_at: string
  }

  approval_histories:
    ApprovalHistoryResponse[]

  created_at: string
  updated_at: string
}