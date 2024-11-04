package restapi

import (
	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo, handler *handler) {
	group := e.Group("api/v1")
	group.GET("/menu", handler.GetMenuList)
}