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
	defaultFlashDuration := 1000
	defaultFlashInterval := 1500

	flashCountInt := utils.StringToInt(flashCount, defaultFlashCount)
	flashDurationInt := utils.StringToInt(flashDuration, defaultFlashDuration)
	flashIntervalInt := utils.StringToInt(flashInterval, defaultFlashInterval)

	if flashCountInt == 0 {
		razer.SetColor(color)
		fmt.Println("Setting color to " + color)
	} else {
		// Use an anonymous func to allow a quick HTTP return to the client
		go func() {
			for i := 0; i < flashCountInt; i++ {
				fmt.Printf("Setting color to %s for %d\n", color, time.Duration(flashDurationInt)*time.Millisecond)
				razer.SetColor(color)
				time.Sleep(time.Duration(flashDurationInt) * time.Millisecond)
				fmt.Printf("Setting color to %s for %d\n", "black", time.Duration(flashIntervalInt)*time.Millisecond)
				razer.SetColor("000000")
				time.Sleep(time.Duration(flashIntervalInt) * time.Millisecond)
			}
		}()
	}
}
