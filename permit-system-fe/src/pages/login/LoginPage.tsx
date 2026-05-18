import { Button, Form, Input, Card, Typography } from "antd"
import { login } from "../../services/auth.service"

const { Title, Text } = Typography

const LoginPage = () => {

  const onFinish = async (values: any) => {
    try {
      await login(values)
      window.location.href = "/"
    } catch (err) {
      console.error(err)
      alert("Login failed")
    }
  }

  return (
    <div className="login-container">

      <Card className="login-card">

        <Title level={2} style={{ textAlign: "center", marginBottom: 6 }}>
          Permit System
        </Title>

        <Text style={{ display: "block", textAlign: "center", marginBottom: 24 }}>
          Sign in to continue
        </Text>

        <Form layout="vertical" onFinish={onFinish}>

          <Form.Item
            name="email"
            rules={[{ required: true, message: "Email required" }]}
          >
            <Input placeholder="Email" size="large" />
          </Form.Item>

          <Form.Item
            name="password"
            rules={[{ required: true, message: "Password required" }]}
          >
            <Input.Password placeholder="Password" size="large" />
          </Form.Item>

          <Button
            type="primary"
            htmlType="submit"
            block
            size="large"
          >
            Login
          </Button>

        </Form>

      </Card>

    </div>
  )
}

export default LoginPage