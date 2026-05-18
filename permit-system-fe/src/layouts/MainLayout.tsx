import {
  Layout,
  Menu,
  Button,
  Typography,
} from "antd"

import {
  Outlet,
  useLocation,
  useNavigate,
} from "react-router-dom"

import {
  sidebarMenus,
} from "../constants/sidebarMenu"

import {
  getUser,
} from "../utils/auth"

import {
  logout,
} from "../services/auth.service"

const {
  Sider,
  Header,
  Content,
} = Layout

const { Title } = Typography

const MainLayout = () => {

  const navigate = useNavigate()

  const location = useLocation()

  const user = getUser()

  const filteredMenus =
    sidebarMenus.filter((menu) =>
      menu.roles.includes(
        user?.role || ""
      )
    )

  const handleLogout =
    async () => {

      await logout()
    }

  return (
    <Layout
      style={{
        minHeight: "100vh",
      }}
    >

      <Sider>

        <div
          style={{
            padding: 16,
            color: "white",
            fontWeight: "bold",
          }}
        >
          Permit License
        </div>

        <Menu
          theme="dark"
          mode="inline"
          selectedKeys={[
            location.pathname,
          ]}
          items={filteredMenus.map(
            (menu) => ({
              key: menu.path,
              icon: menu.icon,
              label: menu.label,
              onClick: () =>
                navigate(menu.path),
            })
          )}
        />

      </Sider>

      <Layout>

        <Header
          style={{
            background: "#fff",
            display: "flex",
            justifyContent:
              "space-between",
            alignItems: "center",
            paddingInline: 24,
          }}
        >

          <Title
            level={4}
            style={{
              margin: 0,
            }}
          >
            Welcome,
            {" "}
            {user?.full_name}
          </Title>

          <Button
            danger
            onClick={handleLogout}
          >
            Logout
          </Button>

        </Header>

        <Content
          style={{
            padding: 24,
          }}
        >

          <Outlet />

        </Content>

      </Layout>

    </Layout>
  )
}

export default MainLayout