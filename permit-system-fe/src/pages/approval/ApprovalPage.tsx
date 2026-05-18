import {
  Button,
  Card,
  Input,
  Modal,
  Space,
  Table,
  Tag,
  Typography,
  message,
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

import {
  approveDocument,
  rejectDocument,
} from "../../services/approval.service"

import {
  getUser,
} from "../../utils/auth"

import {
  STATUS,
} from "../../constants/status"

import type {
  PaginationResponse,
  PermitLicenseItem,
} from "../../types/permit-license"

const { Title } = Typography

const ApprovalPage = () => {

  const navigate = useNavigate()

  const user = getUser()

  const canApprove =
    user?.role === "HEAD_UNIT" ||
    user?.role === "DIRECTOR"

  const [loading, setLoading] =
    useState(false)

  const [search, setSearch] =
    useState("")

  const [page, setPage] =
    useState(1)

  const [selectedId, setSelectedId] =
    useState("")

  const [notes, setNotes] =
    useState("")

  const [
    modalApproveOpen,
    setModalApproveOpen,
  ] = useState(false)

  const [
    modalRejectOpen,
    setModalRejectOpen,
  ] = useState(false)

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

        const params: any = {
          page,
          search,
        }

        if (
          user?.role ===
          "HEAD_UNIT"
        ) {

          params.status =
            STATUS
              .WAITING_HEAD_APPROVAL
        }

        if (
          user?.role ===
          "DIRECTOR"
        ) {

          params.status =
            STATUS
              .WAITING_DIRECTOR_APPROVAL
        }

        const response =
          await getPermitLicenses(
            params
          )

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

  const handleApprove =
    async () => {

      try {

        await approveDocument(
          selectedId,
          {
            notes,
          }
        )

        message.success(
          "Document approved"
        )

        setModalApproveOpen(
          false
        )

        setNotes("")

        fetchData()

      } catch (err) {

        console.error(err)

        message.error(
          "Approve failed"
        )
      }
    }

  const handleReject =
    async () => {

      try {

        await rejectDocument(
          selectedId,
          {
            notes,
          }
        )

        message.success(
          "Document rejected"
        )

        setModalRejectOpen(
          false
        )

        setNotes("")

        fetchData()

      } catch (err) {

        console.error(err)

        message.error(
          "Reject failed"
        )
      }
    }

  if (!canApprove) {

    return (
      <div>
        Access denied
      </div>
    )
  }

  return (
    <div>

      <Title level={3}>
        Approval
      </Title>

      <Card>

        <Space
          style={{
            marginBottom: 16,
          }}
        >

          <Input
            placeholder="Search..."
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
            onClick={fetchData}
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
              title: "Status",

              render: (_, record) => (
                <Tag color="orange">
                  {record.status}
                </Tag>
              ),
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

                  <Button
                    type="primary"
                    onClick={() => {

                      setSelectedId(
                        record.id
                      )

                      setModalApproveOpen(
                        true
                      )
                    }}
                  >
                    Approve
                  </Button>

                  <Button
                    danger
                    onClick={() => {

                      setSelectedId(
                        record.id
                      )

                      setModalRejectOpen(
                        true
                      )
                    }}
                  >
                    Reject
                  </Button>

                </Space>
              ),
            },
          ]}
        />

      </Card>

      <Modal
        open={modalApproveOpen}
        title="Approve Document"
        onCancel={() =>
          setModalApproveOpen(
            false
          )
        }
        onOk={handleApprove}
      >

        <Input.TextArea
          rows={4}
          placeholder="Notes"
          value={notes}
          onChange={(e) =>
            setNotes(
              e.target.value
            )
          }
        />

      </Modal>

      <Modal
        open={modalRejectOpen}
        title="Reject Document"
        onCancel={() =>
          setModalRejectOpen(
            false
          )
        }
        onOk={handleReject}
      >

        <Input.TextArea
          rows={4}
          placeholder="Reject reason"
          value={notes}
          onChange={(e) =>
            setNotes(
              e.target.value
            )
          }
        />

      </Modal>

    </div>
  )
}

export default ApprovalPage