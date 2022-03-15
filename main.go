package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mel-rob/poc-deepmap-openapi/generated/api"
)

// Here's the server
func main() {
	e := echo.New()

	// This allows us to log errors.
	e.Use(middleware.Logger())

	// This prevents our server from crashing on an error
	e.Use(middleware.Recover())

	// This allows us to load the swagger spec
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8081"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// This builds routes based on the OpenAPI spec
	userAccountService := api.NewUserAccount()
	api.RegisterHandlers(e, userAccountService)

	e.Logger.Fatal(e.Start(":8081"))
}
