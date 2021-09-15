package handlers

import (
	"net/http"

	"github.com/jessemillar/razer-chroma-go-wrapper/internal/utils"
	"github.com/labstack/echo/v4"
)

func SolidColor(c echo.Context) error {
	color := c.Param("color")

	utils.FlashColor(color, "0", "0", "0")

	return c.String(http.StatusOK, color)
}

func FlashColor(c echo.Context) error {
	color := c.Param("color")

	utils.FlashColor(color, c.QueryParam("count"), c.QueryParam("duration"), c.QueryParam("interval"))

	return c.String(http.StatusOK, color)
}
