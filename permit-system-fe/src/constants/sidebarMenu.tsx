import {
  DashboardOutlined,
  FileTextOutlined,
  CheckCircleOutlined,
  TeamOutlined,
  ApartmentOutlined,
} from "@ant-design/icons"

export interface SidebarMenuItem {
  key: string
  label: string
  path: string
  roles: string[]
  icon: React.ReactNode
}

export const sidebarMenus:
  SidebarMenuItem[] = [

  {
    key: "dashboard",
    label: "Dashboard",
    path: "/",
    roles: [
      "USER_UNIT",
      "HEAD_UNIT",
      "DIRECTOR",
    ],
    icon: <DashboardOutlined />,
  },

  {
    key: "my-activity",
    label: "My Activity",
    path: "/my-activity",
    roles: [
      "USER_UNIT",
      "HEAD_UNIT",
      "DIRECTOR",
    ],
    icon: <FileTextOutlined />,
  },

  {
    key: "approval",
    label: "Approval",
    path: "/approval",
    roles: [
      "HEAD_UNIT",
      "DIRECTOR",
    ],
    icon: <CheckCircleOutlined />,
  },

  {
    key: "master-document",
    label: "Master Document",
    path: "/master-document",
    roles: ["DIRECTOR"],
    icon: <ApartmentOutlined />,
  },

  {
    key: "unit-management",
    label: "Unit Management",
    path: "/unit",
    roles: ["DIRECTOR"],
    icon: <TeamOutlined />,
  },
]