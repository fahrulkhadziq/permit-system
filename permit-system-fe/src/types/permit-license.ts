export interface PermitLicenseItem {
  id: string
  document_name: string
  description: string
  file_url: string
  file_size: number
  expired_at: string
  status: string
  uploaded_by: string
  unit: string

  master_document: {
    id: string
    name: string
  }

  related_prev_document_id?: string
}

export interface PaginationResponse<T> {
  page: number
  limit: number
  total_rows: number
  total_pages: number
  data: T[]
}

export interface CreatePermitLicensePayload {
  master_document_id: string
  document_name: string
  description: string
  expired_at: string

  related_prev_document_id?: string

  file: File
}