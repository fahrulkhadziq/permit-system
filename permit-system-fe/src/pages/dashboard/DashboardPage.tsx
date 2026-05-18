import {
  Card,
  Col,
  Row,
  Spin,
  Typography,
} from "antd"

import { useEffect, useState }
from "react"

import {
  getDashboardStatistics,
} from "../../services/dashboard.service"

import type {
  DashboardResponse,
} from "../../types/dashboard"

const { Title } = Typography

const DashboardPage = () => {

  const [loading, setLoading] =
    useState(true)

  const [data, setData] =
    useState<DashboardResponse>()

  const fetchDashboard =
    async () => {

      try {

        const response =
          await getDashboardStatistics()

        setData(response)

      } catch (err) {

        console.error(err)

      } finally {

        setLoading(false)
      }
    }

  useEffect(() => {

    fetchDashboard()

  }, [])

  if (loading) {
    return <Spin />
  }

  return (
    <div>

      <Title level={2}>
        Dashboard
      </Title>

      <Row gutter={[16, 16]}>

        <Col span={8}>
          <Card title="Total Documents">
            <h2>
              {data?.total_documents}
            </h2>
          </Card>
        </Col>

        <Col span={8}>
          <Card title="Active Documents">
            <h2>
              {data?.active_documents}
            </h2>
          </Card>
        </Col>

        <Col span={8}>
          <Card title="Expired Documents">
            <h2>
              {data?.expired_documents}
            </h2>
          </Card>
        </Col>

        <Col span={8}>
          <Card title="Not Extended">
            <h2>
              {data?.not_extended_documents}
            </h2>
          </Card>
        </Col>

        <Col span={8}>
          <Card title="Pending Approvals">
            <h2>
              {data?.pending_approvals}
            </h2>
          </Card>
        </Col>

        <Col span={8}>
          <Card title="Approved Documents">
            <h2>
              {data?.approved_documents}
            </h2>
          </Card>
        </Col>

        <Col span={8}>
          <Card title="Rejected Documents">
            <h2>
              {data?.rejected_documents}
            </h2>
          </Card>
        </Col>

      </Row>

    </div>
  )
}

export default DashboardPage