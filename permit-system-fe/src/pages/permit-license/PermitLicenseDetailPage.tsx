import {
  Button,
  Card,
  Descriptions,
  Divider,
  List,
  Space,
  Tag,
  Typography,
  Modal,
  Input,
  message,
} from "antd"

import {
  useEffect,
  useState,
} from "react"

import {
  LeftOutlined,
  RightOutlined,
} from "@ant-design/icons"

import {
  useNavigate,
  useParams,
} from "react-router-dom"

import {
  getPermitLicenseDetail,
} from "../../services/permit-license.service"

import type {
  PermitLicenseDetailResponse,
} from "../../types/permit-license-detail"

import {
  getUser,
} from "../../utils/auth"

import {
  approveDocument,
  rejectDocument,
} from "../../services/approval.service"

import {
  STATUS,
} from "../../constants/status"

const { Title, Text } =
  Typography

const PermitLicenseDetailPage =
  () => {

    const navigate = useNavigate()

    const user = getUser()

    const { id } = useParams()

    const [loading, setLoading] =
      useState(false)

    const [data, setData] =
      useState<
        PermitLicenseDetailResponse
      >()

    const [
      modalApproveOpen,
      setModalApproveOpen,
    ] = useState(false)

    const [
      modalRejectOpen,
      setModalRejectOpen,
    ] = useState(false)

    const [notes, setNotes] =
      useState("")

    const [actionLoading, setActionLoading] =
      useState(false)

    const fetchDetail =
      async () => {

        try {

          setLoading(true)

          const response =
            await getPermitLicenseDetail(
              id!
            )

          setData(response)

        } catch (err) {

          console.error(err)

        } finally {

          setLoading(false)
        }
      }

    useEffect(() => {

      fetchDetail()

    }, [id])

    const canApprove =
      (
        user?.role === "HEAD_UNIT" &&
        data?.current_status.code ===
          STATUS.WAITING_HEAD_APPROVAL
      ) ||
      (
        user?.role === "DIRECTOR" &&
        data?.current_status.code ===
          STATUS.WAITING_DIRECTOR_APPROVAL
      )

      const handleApprove =
        async () => {

          if (!data?.id) return

          try {

            setActionLoading(true)

            await approveDocument(
              data.id,
              {
                notes,
              },
            )

            message.success(
              "Document approved",
            )

            setModalApproveOpen(false)

            setNotes("")

            fetchDetail()

          } catch (err) {

            console.error(err)

            message.error(
              "Approve failed",
            )

          } finally {

            setActionLoading(false)
          }
        }

      const handleReject =
        async () => {

          if (!data?.id) return

          try {

            setActionLoading(true)

            await rejectDocument(
              data.id,
              {
                notes,
              },
            )

            message.success(
              "Document rejected",
            )

            setModalRejectOpen(false)

            setNotes("")

            fetchDetail()

          } catch (err) {

            console.error(err)

            message.error(
              "Reject failed",
            )

          } finally {

            setActionLoading(false)
          }
        }

    return (
      <div>

        <Space
            style={{
                marginBottom: 16,
                width: "100%",
                justifyContent:
                "space-between",
            }}
            >

            <Space>

                <Button
                onClick={() =>
                    navigate(-1)
                }
                >
                Back
                </Button>

                <Title
                level={3}
                style={{
                    margin: 0,
                }}
                >
                Document Detail
                </Title>

            </Space>

            {
                user?.role === "USER_UNIT" &&
                data?.current_status.code ===
                "APPROVED" && data?.is_extend == null && (

                <Button
                    type="primary"
                    onClick={() =>
                    navigate(
                        `/my-activity/create?reference=${data.id}`
                    )
                    }
                >
                    Extend Permit
                </Button>
                )
            }

            {
              (
                user?.role === "HEAD_UNIT" ||
                user?.role === "DIRECTOR"
              ) && (

                <Space>

                  <Button
                    type="primary"
                    disabled={!canApprove}
                    onClick={() =>
                      setModalApproveOpen(
                        true,
                      )
                    }
                  >
                    Approve
                  </Button>

                  <Button
                    danger
                    disabled={!canApprove}
                    onClick={() =>
                      setModalRejectOpen(
                        true,
                      )
                    }
                  >
                    Reject
                  </Button>

                </Space>
              )
            }

            </Space>

        <Card loading={loading}>

          <Descriptions
            bordered
            column={1}
          >

            <Descriptions.Item
              label="Document Name"
            >
              {data?.document_name}
            </Descriptions.Item>

            <Descriptions.Item
              label="Description"
            >
              {data?.description}
            </Descriptions.Item>

            <Descriptions.Item
              label="Master Document"
            >
              {
                data?.master_document
                  ?.name
              }
            </Descriptions.Item>

            <Descriptions.Item
              label="Status"
            >
              <Tag color="blue">
                {
                  data
                    ?.current_status
                    ?.name
                }
              </Tag>
            </Descriptions.Item>

            <Descriptions.Item
              label="Expired At"
            >
              {data?.expired_at}
            </Descriptions.Item>

            <Descriptions.Item
              label="Uploaded By"
            >
              {data?.user.full_name}
            </Descriptions.Item>

            <Descriptions.Item
              label="Unit"
            >
              {data?.unit.name}
            </Descriptions.Item>

            <Descriptions.Item
              label="File"
            >

              <Button
                type="primary"
                href={`http://localhost:8080${data?.file_url}`}
                target="_blank"
              >
                View PDF
              </Button>

            </Descriptions.Item>

            {
            data?.related_prev_document && (

                <Descriptions.Item
                label="Previous Document"
                >

                <Button
                    type="link"
                    icon={<LeftOutlined />}
                    style={{
                    padding: 0,
                    }}
                    onClick={() =>
                    navigate(
                        `/permit-license/${data.related_prev_document?.id}`
                    )
                    }
                >
                    {
                    data
                        .related_prev_document
                        ?.document_name
                    }
                </Button>

                </Descriptions.Item>
            )
            }

            {
            data?.related_next_document && (

                <Descriptions.Item
                label="Next Document"
                >

                <Button
                    type="link"
                    iconPosition="end"
                    icon={<RightOutlined />}
                    style={{
                    padding: 0,
                    }}
                    onClick={() =>
                    navigate(
                        `/permit-license/${data.related_next_document?.id}`
                    )
                    }
                >
                    {
                    data
                        .related_next_document
                        ?.document_name
                    }
                </Button>

                </Descriptions.Item>
            )
            }

            {data
              ?.rejected_reason && (

              <Descriptions.Item
                label="Rejected Reason"
              >

                <Text type="danger">
                  {
                    data
                      ?.rejected_reason
                  }
                </Text>

              </Descriptions.Item>
            )}

          </Descriptions>

          <Divider />

          <Title level={4}>
            Approval Histories
          </Title>

          <List
            dataSource={
              data?.approval_histories
            }
            renderItem={(item) => (

              <List.Item>

                <List.Item.Meta
                  title={
                    item.status.name
                  }
                  description={
                    <>
                      <div>
                        Approver:
                        {" "}
                        {
                          item.approver
                            .full_name
                        }
                      </div>

                      <div>
                        Notes:
                        {" "}
                        {item.notes}
                      </div>

                      <div>
                        Date:
                        {" "}
                        {
                          item.created_at
                        }
                      </div>
                    </>
                  }
                />

              </List.Item>
            )}
          />

        </Card>

        <Modal
          open={modalApproveOpen}
          title="Approve Document"
          confirmLoading={actionLoading}
          onCancel={() =>
            setModalApproveOpen(
              false,
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
                e.target.value,
              )
            }
          />

        </Modal>

        <Modal
          open={modalRejectOpen}
          title="Reject Document"
          confirmLoading={actionLoading}
          onCancel={() =>
            setModalRejectOpen(
              false,
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
                e.target.value,
              )
            }
          />

        </Modal>
      </div>
    )
  }

export default
PermitLicenseDetailPage