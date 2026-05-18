import api from "../lib/axios"

import type {
  PaginationResponse,
  PermitLicenseItem,
  CreatePermitLicensePayload,
} from "../types/permit-license"

import type {
  PermitLicenseDetailResponse,
} from "../types/permit-license-detail"

interface PermitLicenseQuery {
  page?: number
  search?: string
  approval?: boolean
}

export const getPermitLicenses =
  async (
    params?: PermitLicenseQuery
  ): Promise<
    PaginationResponse<PermitLicenseItem>
  > => {

    const response = await api.get(
      "/permit-license",
      {
        params,
      }
    )

    return response.data
  }

  export const getPermitLicenseDetail =
  async (
    id: string
  ): Promise<
    PermitLicenseDetailResponse
  > => {

    const response = await api.get(
      `/permit-license/${id}`
    )

    return response.data
  }

  export const createPermitLicense =
  async (
    payload:
      CreatePermitLicensePayload
  ) => {

    const formData =
      new FormData()

    formData.append(
      "master_document_id",
      payload.master_document_id
    )

    formData.append(
      "document_name",
      payload.document_name
    )

    formData.append(
      "description",
      payload.description
    )

    formData.append(
      "expired_at",
      payload.expired_at
    )

    if (
      payload.related_prev_document_id
    ) {

      formData.append(
        "related_prev_document_id",
        payload
          .related_prev_document_id
      )
    }

    formData.append(
      "file",
      payload.file
    )

    const response =
      await api.post(
        "/permit-license",
        formData,
        {
          headers: {
            "Content-Type":
              "multipart/form-data",
          },
        }
      )

    return response.data
  }