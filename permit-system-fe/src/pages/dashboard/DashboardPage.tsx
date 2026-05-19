import {
  Card,
  Col,
  Modal,
  Row,
  Space,
  Spin,
  Table,
  Tag,
  Typography,
  Button,
} from "antd"

import {
  Pie,
} from "@ant-design/plots"

import {
  useEffect,
  useState,
} from "react"

import {
  useNavigate,
} from "react-router-dom"

import {
  getDashboardStatistics,
} from "../../services/dashboard.service"

import {
  getPermitLicenses,
} from "../../services/permit-license.service"

import type {
  DashboardResponse,
} from "../../types/dashboard"

import type {
  PaginationResponse,
  PermitLicenseItem,
} from "../../types/permit-license"

const { Title } = Typography

const DashboardPage = () => {

  const navigate = useNavigate()

  const [loading, setLoading] =
    useState(true)

  const [data, setData] =
    useState<DashboardResponse>()

  const [
    modalOpen,
    setModalOpen,
  ] = useState(false)

  const [
    modalTitle,
    setModalTitle,
  ] = useState("")

  const [
    selectedStatus,
    setSelectedStatus,
  ] = useState("")

  const [
    documents,
    setDocuments,
  ] =
    useState<
      PaginationResponse<
        PermitLicenseItem
      >
    >()

  const [page, setPage] =
    useState(1)

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

  const fetchDocuments =
    async (
      status?: string,
      currentPage = 1,
    ) => {

      try {

        const response =
          await getPermitLicenses({
            page: currentPage,
            status,
          })

        setDocuments(response)

      } catch (err) {

        console.error(err)
      }
    }

  useEffect(() => {

    fetchDashboard()

  }, [])

  useEffect(() => {

    if (modalOpen) {

      fetchDocuments(
        selectedStatus,
        page,
      )
    }

  }, [
    modalOpen,
    selectedStatus,
    page,
  ])

  const openDocumentModal =
    (
      title: string,
      status?: string,
    ) => {

      setModalTitle(title)

      setSelectedStatus(
        status || "",
      )

      setPage(1)

      setModalOpen(true)
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

  const getStatusColor =
    (status: string) => {

      switch (status) {

        case "Approved":
          return "green"

        case "Rejected":
          return "red"

        case "Waiting Head Approval":
          return "orange"

        case "Waiting Director Approval":
          return "blue"

        default:
          return "default"
      }
    }

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
              {
                data?.total_documents
              }
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
              Approved Documents
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
              Rejected Documents
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
                "Pending Documents",
                "WAITING_APPROVAL",
              )
            }
          >

            <Title level={5}>
              Pending Approvals
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

        <Col span={14}>

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
              onReady={(plot) => {

                plot.on(
                  "element:click",
                  (args: any) => {

                    const type =
                      args.data.data.type

                    if (
                      type ===
                      "Approved"
                    ) {

                      openDocumentModal(
                        "Approved Documents",
                        "APPROVED",
                      )
                    }

                    if (
                      type ===
                      "Rejected"
                    ) {

                      openDocumentModal(
                        "Rejected Documents",
                        "REJECTED",
                      )
                    }

                    if (
                      type ===
                      "Pending"
                    ) {

                      openDocumentModal(
                        "Pending Documents",
                        "WAITING_APPROVAL",
                      )
                    }
                  },
                )
              }}
            />

          </Card>

        </Col>

        <Col span={10}>

          <Card title="Quick Insights">

            <Space
              direction="vertical"
            >

              <Card
                size="small"
              >
                Active Documents:
                {" "}
                <b>
                  {
                    data
                      ?.active_documents
                  }
                </b>
              </Card>

              <Card
                size="small"
              >
                Expired Documents:
                {" "}
                <b>
                  {
                    data
                      ?.expired_documents
                  }
                </b>
              </Card>

              <Card
                size="small"
              >
                Not Extended:
                {" "}
                <b>
                  {
                    data
                      ?.not_extended_documents
                  }
                </b>
              </Card>

            </Space>

          </Card>

        </Col>

      </Row>

      <Modal
        open={modalOpen}
        footer={null}
        width={1000}
        title={modalTitle}
        onCancel={() =>
          setModalOpen(false)
        }
      >

        <Table
          rowKey="id"
          dataSource={
            documents?.data
          }
          pagination={{
            current:
              documents?.page,

            pageSize:
              documents?.limit,

            total:
              documents?.total_rows,

            onChange: (
              page,
            ) =>
              setPage(page),
          }}
          columns={[
            {
              title: "Document",
              dataIndex:
                "document_name",
            },

            {
              title:
                "Master Document",

              render:
                (_, record) =>
                  record
                    .master_document
                    ?.name,
            },

            {
              title: "Status",

              render:
                (_, record) => (

                  <Tag
                    color={getStatusColor(
                      record.status,
                    )}
                  >
                    {
                      record.status
                    }
                  </Tag>
                ),
            },

            {
              title:
                "Expired At",

              dataIndex:
                "expired_at",
            },

            {
              title: "Action",

              render:
                (_, record) => (

                  <Button
                    type="primary"
                    onClick={() => {

                      setModalOpen(
                        false,
                      )

                      navigate(
                        `/permit-license/${record.id}`,
                      )
                    }}
                  >
                    Detail
                  </Button>
                ),
            },
          ]}
        />

      </Modal>

    </div>
  )
}

export default DashboardPage

