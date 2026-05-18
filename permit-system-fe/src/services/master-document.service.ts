import api from "../lib/axios"

import type {
  CreateMasterDocumentPayload,
  MasterDocument,
} from "../types/master-document"

import type {
  PaginationResponse,
} from "../types/permit-license"

export const getMasterDocuments =
  async (
    page = 1,
    search = ""
  ): Promise<
    PaginationResponse<
      MasterDocument
    >
  > => {

    const response = await api.get(
      "/master-document",
      {
        params: {
          page,
          search,
        },
      }
    )

    return response.data
  }

export const createMasterDocument =
  async (
    payload:
      CreateMasterDocumentPayload
  ) => {

    const response = await api.post(
      "/master-document",
      payload
    )

    return response.data
  }

export const updateMasterDocument =
  async (
    id: string,
    payload:
      CreateMasterDocumentPayload
  ) => {

    const response = await api.put(
      `/master-document/${id}`,
      payload
    )

    return response.data
  }

export const deleteMasterDocument =
  async (id: string) => {

    const response = await api.delete(
      `/master-document/${id}`
    )

    return response.data
  }