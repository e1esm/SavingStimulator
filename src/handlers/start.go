package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(handler Handler, port int) error {
	e := echo.New()

	g := e.Group("api/v1/", middleware.Recover())

	addEndpoints(g, handler.Urls())

	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		return fmt.Errorf("error while setting up router: %w", err)
	}
	return nil
}

func addEndpoints(g *echo.Group, endpoints map[string]map[string]echo.HandlerFunc) []echo.MiddlewareFunc {
	handlers := make([]echo.MiddlewareFunc, len(endpoints))
	for url, endpoint := range endpoints {
		for method, f := range endpoint {
			g.Add(method, url, f)
		}
	}

	return handlers
}
