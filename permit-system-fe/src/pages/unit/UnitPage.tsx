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
  createUnit,
  deleteUnit,
  getUnits,
  updateUnit,
} from "../../services/unit.service"

import type {
  Unit,
} from "../../types/unit"

import type {
  PaginationResponse,
} from "../../types/permit-license"

const { Title } = Typography

const UnitPage = () => {

  const [loading, setLoading] =
    useState(false)

  const [search, setSearch] =
    useState("")

  const [page, setPage] =
    useState(1)

  const [open, setOpen] =
    useState(false)

  const [editingId, setEditingId] =
    useState("")

  const [form] = Form.useForm()

  const [data, setData] =
    useState<
      PaginationResponse<Unit>
    >()

  const fetchData =
    async () => {

      try {

        setLoading(true)

        const response =
          await getUnits(
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

          await updateUnit(
            editingId,
            values
          )

          message.success(
            "Unit updated"
          )

        } else {

          await createUnit(
            values
          )

          message.success(
            "Unit created"
          )
        }

        setOpen(false)

        form.resetFields()

        setEditingId("")

        fetchData()

      } catch (err: any) {

        console.error(err)

        message.error(
          err?.response?.data?.message ||
          "Operation failed"
        )
      }
    }

  const handleDelete =
    async (
      id: string
    ) => {

      try {

        await deleteUnit(id)

        message.success(
          "Unit deleted"
        )

        fetchData()

      } catch (err: any) {

        console.error(err)

        message.error(
          err?.response?.data?.message ||
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
          Unit Management
        </Title>

        <Button
          type="primary"
          onClick={() => {

            setEditingId("")

            form.resetFields()

            setOpen(true)
          }}
        >
          Add Unit
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
              title: "Name",
              dataIndex: "name",
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
                    title="Delete unit?"
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
            ? "Edit Unit"
            : "Create Unit"
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
            label="Unit Name"
            name="name"
            rules={[
              {
                required: true,
                message:
                  "Unit name required",
              },
            ]}
          >

            <Input />

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

export default UnitPage