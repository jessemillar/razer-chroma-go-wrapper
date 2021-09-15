package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jessemillar/razer-chroma-go-wrapper/internal/handlers"
	"github.com/jessemillar/razer-chroma-go-wrapper/internal/utils"
	"github.com/jessemillar/razer-chroma-go-wrapper/pkg/razer"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Launching...")

	err := utils.ReadConfigFile()
	if err != nil {
		log.Fatal("Could not read config file:", err.Error())
	}

	razer.CreateApp()
	go razer.PingHeartbeat()

	// TODO Find a better way to do this instead of sleeping
	fmt.Println("Waiting for Razer session...")
	time.Sleep(3 * time.Second)

	razer.SetDefaultColor(viper.GetString("default_color"))
	razer.SetColor("")

	fmt.Println("Starting server...")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// TODO Make a doc page for the default handler
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/color/:color", handlers.SolidColor)
	e.GET("/flash/color/:color", handlers.FlashColor)
	e.Logger.Fatal(e.Start(":" + viper.GetString("server_port")))
}
