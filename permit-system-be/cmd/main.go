package main

import (
	"os"
	"permit-license/config"
	"permit-license/internal/cron"
	"permit-license/internal/database"
	"permit-license/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitTimezone()
	config.LoadEnv()
	config.ConnectDB()

	cron.StartCron()

	database.Migrate()
	database.Seed()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Static("/storage", "storage")

	routes.SetUpRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
