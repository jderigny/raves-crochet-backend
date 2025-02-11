package main

import (
	"net/http"
	"os"

	"github.com/jderigny/raves-crochet-backend/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	mongoClient, mongoContext, mongoContextCancel := db.OpenConnection()
	defer db.CloseConnection(mongoClient, mongoContext, mongoContextCancel)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
