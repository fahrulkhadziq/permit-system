import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom"

import LoginPage
from "./pages/login/LoginPage"

import ProtectedRoute
from "./components/protectedRoute"
import DashboardPage from "./pages/dashboard/DashboardPage"

import MainLayout
from "./layouts/MainLayout"

import MyActivityPage
from "./pages/my-activity/MyActivityPage"

import PermitLicenseDetailPage
from "./pages/permit-license/PermitLicenseDetailPage"

import CreatePermitLicensePage
from "./pages/permit-license/CreatePermitLicensePage"

import ApprovalPage
from "./pages/approval/ApprovalPage"

import MasterDocumentPage
from "./pages/master-document/MasterDocumentPage"

import UnitPage
from "./pages/unit/UnitPage"

function App() {

  return (
    <BrowserRouter>

      <Routes>

        <Route
          path="/login"
          element={<LoginPage />}
        />

        <Route
          element={<ProtectedRoute />}
        >

          <Route
            element={
              <MainLayout />
          }
          >

            <Route
              path="/"
              element={<DashboardPage />}
            />

            <Route
              path="/my-activity"
              element={<MyActivityPage />}
            />

            <Route
              path="/permit-license/:id"
              element={<PermitLicenseDetailPage />}
            />

            <Route
              path="/my-activity/create"
              element={<CreatePermitLicensePage />}
            />

            <Route
              path="/approval"
              element={<ApprovalPage />}
            />

            <Route
              path="/master-document"
              element={<MasterDocumentPage />}
            />

            <Route
              path="/unit"
              element={<UnitPage />}
            />

          </Route>

        </Route>

      </Routes>

    </BrowserRouter>
  )
}

export default App