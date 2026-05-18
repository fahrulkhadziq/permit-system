import {
  Button,
  Card,
  Input,
  Space,
  Table,
  Tag,
  Typography,
} from "antd"

import {
  useEffect,
  useState,
} from "react"

import {
  useNavigate,
} from "react-router-dom"

import {
  getPermitLicenses,
} from "../../services/permit-license.service"

import type {
  PaginationResponse,
  PermitLicenseItem,
} from "../../types/permit-license"

import {
  getUser,
} from "../../utils/auth"

const { Title } = Typography

const MyActivityPage = () => {

  const navigate = useNavigate()

  const user = getUser()

  const [loading, setLoading] =
    useState(false)

  const [search, setSearch] =
    useState("")

  const [page, setPage] =
    useState(1)

  const [data, setData] =
    useState<
      PaginationResponse<
        PermitLicenseItem
      >
    >()

  const fetchData =
    async () => {

      try {

        setLoading(true)

        const response =
          await getPermitLicenses({
            page,
            search,
          })

        setData(response)

      } catch (err) {

        console.error(err)

      } finally {

        setLoading(false)
      }
    }

  useEffect(() => {

    fetchData()

  }, [page])

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

  const canUpload =
    user?.role === "USER_UNIT"

  return (
    <div>

      <Space
        style={{
          width: "100%",
          justifyContent:
            "space-between",
          marginBottom: 16,
        }}
      >

        <Title level={3}>
          My Activity
        </Title>

        {canUpload && (
          <Button
            type="primary"
            onClick={() =>
              navigate(
                "/my-activity/create"
              )
            }
          >
            Upload Document
          </Button>
        )}

      </Space>

      <Card>

        <Space
          style={{
            marginBottom: 16,
          }}
        >

          <Input
            placeholder="Search document..."
            value={search}
            onChange={(e) =>
              setSearch(
                e.target.value
              )
            }
            onPressEnter={() =>
              fetchData()
            }
          />

          <Button
            type="primary"
            onClick={() =>
              fetchData()
            }
          >
            Search
          </Button>

        </Space>

        <Table
          rowKey="id"
          loading={loading}
          dataSource={data?.data}
          pagination={{
            current: data?.page,
            pageSize: data?.limit,
            total: data?.total_rows,
            onChange: (page) =>
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

              render: (_, record) =>
                record
                  .master_document
                  ?.name,
            },

            {
              title: "Status",

              render: (_, record) => (
                <Tag
                  color={getStatusColor(
                    record.status
                  )}
                >
                  {record.status}
                </Tag>
              ),
            },

            {
              title: "Expired At",
              dataIndex:
                "expired_at",
            },

           {
            title: "Action",

            render: (_, record) => (

                <Space>

                <Button
                    onClick={() =>
                    navigate(
                        `/permit-license/${record.id}`
                    )
                    }
                >
                    Detail
                </Button>

                {
                    user?.role ===
                    "USER_UNIT" &&
                    record.status ===
                    "Rejected" && (

                    <Button
                        type="primary"
                        onClick={() =>
                        navigate(
                            `/my-activity/${record.id}/edit`
                        )
                        }
                    >
                        Revise
                    </Button>
                    )
                }

                </Space>
            ),
            },
          ]}
        />

      </Card>

    </div>
  )
}

export default MyActivityPage