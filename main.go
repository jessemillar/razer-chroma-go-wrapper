package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/jessemillar/razer-chroma-go-wrapper/internal/handlers"
	"github.com/jessemillar/razer-chroma-go-wrapper/internal/utils"
	"github.com/jessemillar/razer-chroma-go-wrapper/pkg/razer"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var echoServer *echo.Echo

func main() {
	fmt.Println("Launching...")

	err := utils.ReadConfigFile()
	if err != nil {
		log.Fatal("Could not read config file:", err.Error())
	}

	razer.CreateApp()
	// TODO Kill this as part of cleanup
	go razer.PingHeartbeat()

	// TODO Find a better way to do this instead of sleeping
	fmt.Println("Waiting for Razer session...")
	time.Sleep(3 * time.Second)

	razer.SetDefaultColor(viper.GetString("default_color"))
	razer.SetColor("")

	fmt.Println("Starting server...")

	echoServer := echo.New()
	echoServer.GET("/", func(c echo.Context) error {
		// TODO Make a doc page for the default handler
		return c.String(http.StatusOK, "Hello, World!")
	})
	echoServer.GET("/color/:color", handlers.SolidColor)
	echoServer.GET("/flash/color/:color", handlers.FlashColor)
	// TODO Kill this goroutine gracefully
	go func() { echoServer.Logger.Fatal(echoServer.Start(":" + viper.GetString("server_port"))) }()

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Razer Chroma HTTP Wrapper")
	systray.SetTooltip("Razer Chroma HTTP Wrapper")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	// TODO Set this to a custom icon
	mQuit.SetIcon(icon.Data)

	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func onExit() {
	// clean up here
	fmt.Println("Shutting down")

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := echoServer.Shutdown(ctx); err != nil {
		echoServer.Logger.Fatal(err)
	}
}
