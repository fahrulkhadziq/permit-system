import api from "../lib/axios"

import type {
  PaginationResponse,
  PermitLicenseItem,
  CreatePermitLicensePayload,
  UpdatePermitLicensePayload,
} from "../types/permit-license"

import type {
  PermitLicenseDetailResponse,
} from "../types/permit-license-detail"

export interface PermitLicenseQuery {
  page?: number
  limit?: number
  search?: string
  status?: string
  sort?: string
  order?: string
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

  export const updatePermitLicense =
  async (
    id: string,
    payload:
      UpdatePermitLicensePayload
  ) => {

    const formData =
      new FormData()

    if (
      payload.master_document_id
    ) {

      formData.append(
        "master_document_id",
        payload.master_document_id
      )
    }

    if (
      payload.document_name
    ) {

      formData.append(
        "document_name",
        payload.document_name
      )
    }

    if (
      payload.description
    ) {

      formData.append(
        "description",
        payload.description
      )
    }

    if (
      payload.expired_at
    ) {

      formData.append(
        "expired_at",
        payload.expired_at
      )
    }

    if (
      payload.related_prev_document_id
    ) {

      formData.append(
        "related_prev_document_id",
        payload.related_prev_document_id
      )
    }

    if (
      payload.is_extend !==
      undefined
    ) {

      formData.append(
        "is_extend",
        String(
          payload.is_extend
        )
      )
    }

    if (payload.file) {

      formData.append(
        "file",
        payload.file
      )
    }

    const response =
      await api.put(
        `/permit-license/${id}`,
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