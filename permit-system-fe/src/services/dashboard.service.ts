import api from "../lib/axios"

import type {
  DashboardResponse,
} from "../types/dashboard"

export const getDashboardStatistics =
  async (): Promise<DashboardResponse> => {

    const response = await api.get(
      "/dashboard/statistics"
    )

    return response.data
  }