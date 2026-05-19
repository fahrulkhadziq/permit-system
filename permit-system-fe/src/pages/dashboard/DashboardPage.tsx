import {
  Card,
  Col,
  Modal,
  Row,
  Spin,
  Table,
  Typography,
} from "antd"

import {
  Pie,
} from "@ant-design/plots"

import {
  useEffect,
  useState,
} from "react"

import {
  getDashboardStatistics,
} from "../../services/dashboard.service"

import {
  getPermitLicenses,
} from "../../services/permit-license.service"

import type {
  DashboardResponse,
} from "../../types/dashboard"

const { Title } = Typography

const DashboardPage = () => {

  const [loading, setLoading] =
    useState(true)

  const [data, setData] =
    useState<DashboardResponse>()

  const [modalOpen, setModalOpen] =
    useState(false)

  const [modalTitle, setModalTitle] =
    useState("")

  const [documents, setDocuments] =
    useState<any[]>([])

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

  const openDocumentModal =
    async (
      title: string,
      status?: string,
    ) => {

      try {

        const response =
          await getPermitLicenses({
            status,
          })

        setDocuments(
          response.data,
        )

        setModalTitle(title)

        setModalOpen(true)

      } catch (err) {

        console.error(err)
      }
    }

  if (loading) {

    return <Spin />
  }

  const chartData = [
    {
      type: "Approved",
      value:
        data?.approved_documents || 0,
    },
    {
      type: "Rejected",
      value:
        data?.rejected_documents || 0,
    },
    {
      type: "Pending",
      value:
        data?.pending_approvals || 0,
    },
    {
      type: "Expired",
      value:
        data?.expired_documents || 0,
    },
  ]

  return (
    <div>

      <Title level={2}>
        Dashboard
      </Title>

      <Row gutter={[16, 16]}>

        <Col span={6}>
          <Card
            hoverable
            onClick={() =>
              openDocumentModal(
                "All Documents",
              )
            }
          >
            <Title level={5}>
              Total Documents
            </Title>

            <Title level={2}>
              {data?.total_documents}
            </Title>
          </Card>
        </Col>

        <Col span={6}>
          <Card
            hoverable
            onClick={() =>
              openDocumentModal(
                "Approved Documents",
                "APPROVED",
              )
            }
          >
            <Title level={5}>
              Approved
            </Title>

            <Title
              level={2}
              style={{
                color: "#52c41a",
              }}
            >
              {
                data?.approved_documents
              }
            </Title>
          </Card>
        </Col>

        <Col span={6}>
          <Card
            hoverable
            onClick={() =>
              openDocumentModal(
                "Rejected Documents",
                "REJECTED",
              )
            }
          >
            <Title level={5}>
              Rejected
            </Title>

            <Title
              level={2}
              style={{
                color: "#ff4d4f",
              }}
            >
              {
                data?.rejected_documents
              }
            </Title>
          </Card>
        </Col>

        <Col span={6}>
          <Card
            hoverable
            onClick={() =>
              openDocumentModal(
                "Pending Approvals",
                "WAITING_APPROVAL",
              )
            }
          >
            <Title level={5}>
              Pending
            </Title>

            <Title
              level={2}
              style={{
                color: "#faad14",
              }}
            >
              {
                data?.pending_approvals
              }
            </Title>
          </Card>
        </Col>

        <Col span={12}>

          <Card title="Document Statistics">

            <Pie
              data={chartData}
              angleField="value"
              colorField="type"
              radius={0.9}
              label={{
                type: "outer",
                content:
                  "{name} {percentage}",
              }}
              interactions={[
                {
                  type:
                    "element-active",
                },
              ]}
            />

          </Card>

        </Col>

        <Col span={12}>

          <Card title="Quick Insights">

            <p>
              Active Documents:
              {" "}
              <b>
                {
                  data?.active_documents
                }
              </b>
            </p>

            <p>
              Expired Documents:
              {" "}
              <b>
                {
                  data?.expired_documents
                }
              </b>
            </p>

            <p>
              Not Extended:
              {" "}
              <b>
                {
                  data
                    ?.not_extended_documents
                }
              </b>
            </p>

          </Card>

        </Col>

      </Row>

      <Modal
        open={modalOpen}
        footer={null}
        width={900}
        title={modalTitle}
        onCancel={() =>
          setModalOpen(false)
        }
      >

        <Table
          rowKey="id"
          dataSource={documents}
          pagination={false}
          columns={[
            {
              title: "Document",
              dataIndex:
                "document_name",
            },
            {
              title: "Status",
              dataIndex: "status",
            },
            {
              title: "Expired At",
              dataIndex:
                "expired_at",
            },
          ]}
        />

      </Modal>

    </div>
  )
}

export default DashboardPage
