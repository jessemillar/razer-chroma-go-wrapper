package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jessemillar/razer-chroma-go-wrapper/internal/utils"
	"github.com/jessemillar/razer-chroma-go-wrapper/pkg/razer"
	"github.com/labstack/echo/v4"
)

func SolidColor(c echo.Context) error {
	color := c.Param("color")

	flashColor(color, "0", "0", "0")

	return c.String(http.StatusOK, color)
}

func FlashColor(c echo.Context) error {
	color := c.Param("color")

	flashColor(color, c.QueryParam("count"), c.QueryParam("duration"), c.QueryParam("interval"))

	return c.String(http.StatusOK, color)
}

func flashColor(color string, flashCount string, flashDuration string, flashInterval string) {
	defaultFlashCount := 5
	defaultFlashDuration := float64(1)
	defaultFlashInterval := 1.5

	flashCountInt := utils.StringToInt(flashCount, defaultFlashCount)
	flashDurationFloat := utils.StringToFloat(flashDuration, defaultFlashDuration)
	flashIntervalFloat := utils.StringToFloat(flashInterval, defaultFlashInterval)

	if flashCountInt == 0 {
		razer.SetColor(color)
		fmt.Println("Setting color to " + color)
	} else {
		// Use an anonymous func to allow a quick HTTP return to the client
		go func() {
			for i := 0; i < flashCountInt; i++ {
				fmt.Printf("Setting color to %s for %f", color, flashDurationFloat)
				razer.SetColor(color)
				time.Sleep(time.Duration(flashDurationFloat) * time.Second)
				fmt.Printf("Setting color to %s for %f", "black", flashDurationFloat)
				razer.SetColor("000000")
				time.Sleep(time.Duration(flashIntervalFloat) * time.Second)
			}
		}()
	}
}
