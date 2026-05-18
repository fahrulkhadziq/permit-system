import api from "../lib/axios"

import type {
  Unit,
  UnitPayload,
} from "../types/unit"

import type {
  PaginationResponse,
} from "../types/permit-license"

export const getUnits =
  async (
    page = 1,
    search = ""
  ): Promise<
    PaginationResponse<Unit>
  > => {

    const response =
      await api.get(
        "/unit",
        {
          params: {
            page,
            search,
          },
        }
      )

    return response.data
  }

export const createUnit =
  async (
    payload: UnitPayload
  ) => {

    const response =
      await api.post(
        "/unit",
        payload
      )

    return response.data
  }

export const updateUnit =
  async (
    id: string,
    payload: UnitPayload
  ) => {

    const response =
      await api.put(
        `/unit/${id}`,
        payload
      )

    return response.data
  }

export const deleteUnit =
  async (id: string) => {

    const response =
      await api.delete(
        `/unit/${id}`
      )

    return response.data
  }