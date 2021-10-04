package handlers

import (
	"net/http"

	"github.com/jessemillar/razer-chroma-http-wrapper/internal/utils"
	"github.com/jessemillar/razer-chroma-http-wrapper/pkg/razer"
	"github.com/labstack/echo/v4"
)

func SolidColor(c echo.Context) error {
	color := utils.TranslateCustomColor(c.Param("color"))

	razer.FlashColor(color, "0", "0", "0")

	return c.String(http.StatusOK, color)
}

func FlashColor(c echo.Context) error {
	color := utils.TranslateCustomColor(c.Param("color"))

	razer.FlashColor(color, c.QueryParam("count"), c.QueryParam("duration"), c.QueryParam("interval"))

	return c.String(http.StatusOK, color)
}
