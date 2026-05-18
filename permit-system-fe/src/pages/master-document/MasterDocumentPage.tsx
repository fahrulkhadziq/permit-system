import {
  Button,
  Card,
  Form,
  Input,
  Modal,
  Popconfirm,
  Space,
  Table,
  Typography,
  message,
} from "antd"

import {
  useEffect,
  useState,
} from "react"

import {
  createMasterDocument,
  deleteMasterDocument,
  getMasterDocuments,
  updateMasterDocument,
} from "../../services/master-document.service"

import type {
  MasterDocument,
} from "../../types/master-document"

import type {
  PaginationResponse,
} from "../../types/permit-license"

const { Title } = Typography

const MasterDocumentPage =
  () => {

    const [loading, setLoading] =
      useState(false)

    const [search, setSearch] =
      useState("")

    const [page, setPage] =
      useState(1)

    const [form] = Form.useForm()

    const [editingId, setEditingId] =
      useState("")

    const [open, setOpen] =
      useState(false)

    const [data, setData] =
      useState<
        PaginationResponse<
          MasterDocument
        >
      >()

    const fetchData =
      async () => {

        try {

          setLoading(true)

          const response =
            await getMasterDocuments(
              page,
              search
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

    const handleSubmit =
      async (
        values: any
      ) => {

        try {

          if (editingId) {

            await updateMasterDocument(
              editingId,
              values
            )

            message.success(
              "Master document updated"
            )

          } else {

            await createMasterDocument(
              values
            )

            message.success(
              "Master document created"
            )
          }

          setOpen(false)

          form.resetFields()

          setEditingId("")

          fetchData()

        } catch (err) {

          console.error(err)

          message.error(
            "Operation failed"
          )
        }
      }

    const handleDelete =
      async (id: string) => {

        try {

          await deleteMasterDocument(
            id
          )

          message.success(
            "Deleted successfully"
          )

          fetchData()

        } catch (err) {

          console.error(err)

          message.error(
            "Delete failed"
          )
        }
      }

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
            Master Document
          </Title>

          <Button
            type="primary"
            onClick={() => {

              setEditingId("")

              form.resetFields()

              setOpen(true)
            }}
          >
            Add Master Document
          </Button>

        </Space>

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
                title: "Code",
                dataIndex: "code",
              },

              {
                title: "Name",
                dataIndex: "name",
              },

              {
                title:
                  "Description",
                dataIndex:
                  "description",
              },

              {
                title: "Action",

                render: (_, record) => (

                  <Space>

                    <Button
                      onClick={() => {

                        setEditingId(
                          record.id
                        )

                        form.setFieldsValue(
                          record
                        )

                        setOpen(true)
                      }}
                    >
                      Edit
                    </Button>

                    <Popconfirm
                      title="Delete?"
                      onConfirm={() =>
                        handleDelete(
                          record.id
                        )
                      }
                    >

                      <Button danger>
                        Delete
                      </Button>

                    </Popconfirm>

                  </Space>
                ),
              },
            ]}
          />

        </Card>

        <Modal
          open={open}
          title={
            editingId
              ? "Edit Master Document"
              : "Create Master Document"
          }
          onCancel={() =>
            setOpen(false)
          }
          footer={null}
        >

          <Form
            layout="vertical"
            form={form}
            onFinish={handleSubmit}
          >

            <Form.Item
              label="Code"
              name="code"
              rules={[
                {
                  required: true,
                },
              ]}
            >

              <Input />

            </Form.Item>

            <Form.Item
              label="Name"
              name="name"
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

            <Button
              type="primary"
              htmlType="submit"
              block
            >
              Submit
            </Button>

          </Form>

        </Modal>

      </div>
    )
  }

export default
MasterDocumentPage