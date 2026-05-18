export interface MasterDocument {
  id: string
  code: string
  name: string
  description: string
}

export interface CreateMasterDocumentPayload {
  code: string
  name: string
  description: string
}