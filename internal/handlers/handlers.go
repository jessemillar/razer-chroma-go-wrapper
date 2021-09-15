package handlers

import (
	"net/http"
	"time"

	"github.com/jessemillar/razer-chroma-go-wrapper/internal/utils"
	"github.com/jessemillar/razer-chroma-go-wrapper/pkg/razer"
	"github.com/labstack/echo/v4"
)

func SolidColor(c echo.Context) error {
	color := c.Param("color")

	flashColor(color, "0", "0")

	return c.String(http.StatusOK, color)
}

func FlashColor(c echo.Context) error {
	color := c.Param("color")

	flashColor(color, c.QueryParam("count"), c.QueryParam("duration"))

	return c.String(http.StatusOK, color)
}

func flashColor(color string, flashCount string, flashDuration string) {
	defaultFlashCount := 5
	defaultFlashDuration := float64(1)

	flashCountInt := utils.StringToInt(flashCount, defaultFlashCount)
	flashDurationFloat := utils.StringToFloat(flashDuration, defaultFlashDuration)

	for i := 0; i < flashCountInt; i++ {
		razer.SetColor("#" + color)
		time.Sleep(time.Duration(flashDurationFloat) * time.Second)
		razer.SetColor("#" + color)
		time.Sleep(time.Duration(flashDurationFloat) * time.Second)
	}
}
