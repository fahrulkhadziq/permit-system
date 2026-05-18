import api from "../lib/axios"

export interface ApprovalPayload {
  notes: string
}

export const approveDocument =
  async (
    permitId: string,
    payload: ApprovalPayload
  ) => {

    const endpoint = `/approval/${permitId}/approve`

    const response =
      await api.post(
        endpoint,
        payload
      )

    return response.data
  }

export const rejectDocument =
  async (
    permitId: string,
    payload: ApprovalPayload
  ) => {

    const endpoint =`/approval/${permitId}/reject`

    const response =
      await api.post(
        endpoint,
        payload
      )

    return response.data
  }