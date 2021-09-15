package handlers

import (
	"net/http"

	"github.com/icza/gox/imagex/colorx"
	"github.com/labstack/echo/v4"
)

func ChangeColor(c echo.Context) error {
	color := c.Param("color")
	parsedColor, _ := colorx.ParseHexColor("#" + color)
	createAndApplyEffect(convertColor(int(parsedColor.R), int(parsedColor.G), int(parsedColor.B)))
	return c.String(http.StatusOK, color)
}
