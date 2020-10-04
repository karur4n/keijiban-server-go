package main

import (
	"github.com/karur4n/kujibiki-server-go/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers.RegisterHandlers(e, &handlers.ServerInterfaceImpl{})

	e.Logger.Fatal(e.Start(":1323"))
}
