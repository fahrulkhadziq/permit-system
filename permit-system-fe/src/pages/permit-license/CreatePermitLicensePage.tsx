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
} from "react-router-dom"

import {
  createPermitLicense,
} from "../../services/permit-license.service"

import {
  getMasterDocuments,
} from "../../services/master-document.service"
import type { MasterDocument } from "../../types/master-document"

const { Title } = Typography

const CreatePermitLicensePage =
  () => {

    const navigate = useNavigate()

    const [form] = Form.useForm()

    const [loading, setLoading] =
      useState(false)

    const [masterDocuments,
  setMasterDocuments] =
  useState<MasterDocument[]>([])

    const [file, setFile] =
      useState<File>()

    const fetchMasterDocuments =
      async () => {

        try {

          const response =
            await getMasterDocuments(
                1,
                ""
            )

          setMasterDocuments(
            response.data
          )

        } catch (err) {

          console.error(err)
        }
      }

    useEffect(() => {

      fetchMasterDocuments()

    }, [])

    const onFinish =
      async (values: any) => {

        try {

          if (!file) {

            message.error(
              "File required"
            )

            return
          }

          setLoading(true)

          await createPermitLicense({
            master_document_id:
              values.master_document_id,

            document_name:
              values.document_name,

            description:
              values.description,

            expired_at:
              dayjs(
                values.expired_at
              ).format(
                "YYYY-MM-DD"
              ),

            file,
          })

          message.success(
            "Document uploaded"
          )

          navigate("/my-activity")

        } catch (err) {

          console.error(err)

          message.error(
            "Upload failed"
          )

        } finally {

          setLoading(false)
        }
      }

    return (
      <div>

        <Title level={3}>
          Upload Document
        </Title>

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
              >

                {masterDocuments.map(
                  (item: any) => (

                    <Select.Option
                      key={item.id}
                      value={item.id}
                    >
                      {item.name}
                    </Select.Option>
                  )
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
                  file
                ) => {

                  const isPdf =
                    file.type ===
                    "application/pdf"

                  if (!isPdf) {

                    message.error(
                      "Only PDF allowed"
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
                      "Max 25MB"
                    )

                    return Upload.LIST_IGNORE
                  }

                  setFile(
                    file as File
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
                  Upload PDF
                </Button>

              </Upload>

            </Form.Item>

            <Button
              type="primary"
              htmlType="submit"
              loading={loading}
            >
              Submit
            </Button>

          </Form>

        </Card>

      </div>
    )
  }

export default
CreatePermitLicensePage