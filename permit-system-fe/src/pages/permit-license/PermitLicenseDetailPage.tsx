import {
  Button,
  Card,
  Descriptions,
  Divider,
  List,
  Space,
  Tag,
  Typography,
} from "antd"

import {
  useEffect,
  useState,
} from "react"

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

    }, [])

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
                "APPROVED" && (

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

            {data
                ?.related_prev_document && (

                <Descriptions.Item
                    label="Previous Document"
                >

                    <Space
                    direction="vertical"
                    >

                    <Button
                        type="link"
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

                    <Text type="secondary">
                        Expired At:
                        {" "}
                        {
                        data
                            .related_prev_document
                            ?.expired_at
                        }
                    </Text>

                    </Space>

                </Descriptions.Item>
                )}

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

      </div>
    )
  }

export default
PermitLicenseDetailPage