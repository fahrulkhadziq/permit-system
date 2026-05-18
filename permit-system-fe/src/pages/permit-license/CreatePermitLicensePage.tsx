import {
  Button,
  Card,
  DatePicker,
  Form,
  Input,
  Select,
  Upload,
  message,
  Typography,
  Alert,
  Space,
} from "antd"

import {
  UploadOutlined,
} from "@ant-design/icons"

import dayjs from "dayjs"

import {
  useEffect,
  useState,
} from "react"

import {
  useNavigate,
  useParams,
  useSearchParams,
} from "react-router-dom"

import {
  createPermitLicense,
  getPermitLicenseDetail,
  updatePermitLicense,
} from "../../services/permit-license.service"

import {
  getMasterDocuments,
} from "../../services/master-document.service"

import type {
  MasterDocument,
} from "../../types/master-document"

const { Title, Text } =
  Typography

const CreatePermitLicensePage =
  () => {

    const navigate =
      useNavigate()

    const { id } =
      useParams()

    const isEdit =
      !!id

    const [form] =
      Form.useForm()

    const [loading, setLoading] =
      useState(false)

    const [
      masterDocuments,
      setMasterDocuments,
    ] =
      useState<MasterDocument[]>([])

    const [file, setFile] =
      useState<File>()

    const [
      previousDocument,
      setPreviousDocument,
    ] = useState<any>(null)

    const [searchParams] =
    useSearchParams()

    const referenceId =
    searchParams.get("reference")

    const fetchReference =
        async () => {

            if (!referenceId) return

            try {

            const response =
                await getPermitLicenseDetail(
                referenceId
                )

            form.setFieldsValue({
                master_document_id:
                response.master_document.id,
            })

            } catch (err) {

            console.error(err)
            }
        }

    const fetchMasterDocuments =
      async () => {

        try {

          const response =
            await getMasterDocuments(
              1,
              "",
            )

          setMasterDocuments(
            response.data,
          )

        } catch (err) {

          console.error(err)
        }
      }

    const fetchDetail =
      async () => {

        if (!id) return

        try {

          const response =
            await getPermitLicenseDetail(
              id,
            )

          form.setFieldsValue({
            master_document_id:
              response.master_document.id,

            document_name:
              response.document_name,

            description:
              response.description,

            expired_at:
              dayjs(
                response.expired_at,
              ),
          })

          // ===== PREVIOUS DOC =====
          setPreviousDocument({
            id: response.id,
            document_name:
              response.document_name,
            expired_at:
              response.expired_at,
          })

        } catch (err) {

          console.error(err)

          message.error(
            "Failed load detail",
          )
        }
      }

    useEffect(() => {

      fetchMasterDocuments()

      fetchDetail()

    fetchReference()

    }, [id, referenceId])

    const onFinish =
      async (values: any) => {

        try {

          if (!file) {

            message.error(
              isEdit
                ? "Please upload revised PDF"
                : "File required",
            )

            return
          }

          setLoading(true)

          const payload = {

            master_document_id:
              referenceId
                ? form.getFieldValue(
                    "master_document_id"
                    )
                : values.master_document_id,

            document_name:
              values.document_name,

            description:
              values.description,

            expired_at:
              dayjs(
                values.expired_at,
              ).format(
                "YYYY-MM-DD",
              ),

            file,

            // ===== LINK HISTORY =====
            related_prev_document_id:
              isEdit
                ? previousDocument?.id
                : referenceId || undefined,
          }

          if (isEdit) {

            await updatePermitLicense(
              id,
              payload,
            )

            message.success(
              "Document revised successfully",
            )

          } else {

            await createPermitLicense(
              payload,
            )

            message.success(
              "Document uploaded successfully",
            )
          }

          navigate(
            "/my-activity",
          )

        } catch (err) {

          console.error(err)

          message.error(
            isEdit
              ? "Revise failed"
              : "Upload failed",
          )

        } finally {

          setLoading(false)
        }
      }

    return (
      <div>

        <Title level={3}>
          {
            isEdit
              ? "Revise Document"
              : "Upload Document"
          }
        </Title>

        {/* ===== REVISION INFO ===== */}

        {isEdit &&
          previousDocument && (

            <Alert
              type="warning"
              showIcon
              style={{
                marginBottom: 16,
              }}
              message="Revision Mode"
              description={

                <Space
                  direction="vertical"
                >

                  <Text>
                    This revised document
                    will be linked to:
                  </Text>

                  <Text strong>
                    {
                      previousDocument.document_name
                    }
                  </Text>

                  <Text type="secondary">
                    Expired:
                    {" "}
                    {dayjs(
                      previousDocument.expired_at,
                    ).format(
                      "DD MMM YYYY",
                    )}
                  </Text>

                </Space>
              }
            />
          )}

        <Card>

          <Form
            layout="vertical"
            form={form}
            onFinish={onFinish}
          >

            <Form.Item
              label="Master Document"
              name="master_document_id"
              rules={[
                {
                  required: true,
                },
              ]}
            >

              <Select
                placeholder="Select master document"
                disabled={!!referenceId}
                >

                {masterDocuments.map(
                  (
                    item,
                  ) => (

                    <Select.Option
                      key={item.id}
                      value={item.id}
                    >
                      {item.name}
                    </Select.Option>
                  ),
                )}

              </Select>

            </Form.Item>

            <Form.Item
              label="Document Name"
              name="document_name"
              rules={[
                {
                  required: true,
                },
              ]}
            >

              <Input />

            </Form.Item>

            <Form.Item
              label="Description"
              name="description"
            >

              <Input.TextArea
                rows={4}
              />

            </Form.Item>

            <Form.Item
              label="Expired At"
              name="expired_at"
              rules={[
                {
                  required: true,
                },
              ]}
            >

              <DatePicker
                style={{
                  width: "100%",
                }}
              />

            </Form.Item>

            <Form.Item
              label="PDF File"
              required
            >

              <Upload
                beforeUpload={(
                  file,
                ) => {

                  const isPdf =
                    file.type ===
                    "application/pdf"

                  if (!isPdf) {

                    message.error(
                      "Only PDF allowed",
                    )

                    return Upload.LIST_IGNORE
                  }

                  const isLt25Mb =
                    file.size /
                      1024 /
                      1024 <
                    25

                  if (!isLt25Mb) {

                    message.error(
                      "Max 25MB",
                    )

                    return Upload.LIST_IGNORE
                  }

                  setFile(
                    file as File,
                  )

                  return false
                }}
                maxCount={1}
              >

                <Button
                  icon={
                    <UploadOutlined />
                  }
                >
                  {
                    isEdit
                      ? "Upload Revised PDF"
                      : "Upload PDF"
                  }
                </Button>

              </Upload>

            </Form.Item>

            <Button
              type="primary"
              htmlType="submit"
              loading={loading}
            >
              {
                isEdit
                  ? "Revise Document"
                  : "Submit"
              }
            </Button>

          </Form>

        </Card>

      </div>
    )
  }

export default
CreatePermitLicensePage