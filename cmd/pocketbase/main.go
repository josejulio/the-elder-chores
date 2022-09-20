package main

import (
	"log"
	"net/http"

	_ "github.com/adventureboss/the-elder-chores/migrations"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/hello",
			Handler: func(c echo.Context) error {
				return c.String(200, "lubdub")
			},
		})

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/", "./the-elder-chores-user-ui/build")
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
