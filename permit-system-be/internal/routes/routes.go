package routes

import (
	"permit-license/internal/constants"
	"permit-license/internal/handler"
	"permit-license/internal/middleware"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	authHandler := handler.AuthHandler{}
	permitHandler := handler.PermitLicenseHandler{}
	approvalHandler := handler.ApprovalHandler{}
	dashboardHandler := handler.DashboardHandler{}
	masterDocumentHandler := handler.MasterDocumentHandler{}
	unitHandler := handler.UnitHandler{}

	public := e.Group("/api/auth")
	public.POST("/register", authHandler.Register)
	public.POST("/login", authHandler.Login)

	protected := e.Group("/api")
	protected.Use(middleware.JWTMiddleware)

	protected.POST("/auth/refresh", authHandler.RefreshToken)
	protected.POST("/auth/logout", authHandler.Logout)

	protected.GET("/protected", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "you are authorized",
		})
	})

	protected.POST("/permit-license", permitHandler.Create)
	protected.GET("/permit-license/:id", permitHandler.FindByID)
	protected.GET("/permit-license", permitHandler.FindAll)
	protected.GET("/permit-license/:id/download", permitHandler.Download)
	// HEAD UNIT and DIRECTOR
	approval := protected.Group("/approval")

	approval.POST("/:id/approve",
		approvalHandler.Approve,
	)

	approval.POST("/:id/reject",
		approvalHandler.Reject,
	)

	// DASHBOARD
	protected.GET("/dashboard/statistics", dashboardHandler.GetStatistics)

	// MASTER DOCUMENT
	master := protected.Group("/master-document")
	master.POST("", masterDocumentHandler.Create, middleware.RoleMiddleware(
		constants.RoleDirector,
	))
	master.GET("", masterDocumentHandler.FindAll)
	master.GET("/:id", masterDocumentHandler.FindByID)
	master.PUT("/:id", masterDocumentHandler.Update, middleware.RoleMiddleware(
		constants.RoleDirector,
	))
	master.DELETE("/:id", masterDocumentHandler.Delete, middleware.RoleMiddleware(
		constants.RoleDirector,
	))

	// UNIT
	unit := protected.Group("/unit")
	unit.POST("", unitHandler.Create, middleware.RoleMiddleware(
		constants.RoleDirector,
	))
	unit.GET("", unitHandler.FindAll)
	unit.GET("/:id", unitHandler.FindByID)
	unit.PUT("/:id", unitHandler.Update, middleware.RoleMiddleware(
		constants.RoleDirector,
	))
	unit.DELETE("/:id", unitHandler.Delete, middleware.RoleMiddleware(
		constants.RoleDirector,
	))

}
