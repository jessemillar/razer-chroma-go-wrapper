package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jessemillar/razer-chroma-go-wrapper/internal/handlers"
	"github.com/jessemillar/razer-chroma-go-wrapper/pkg/razer"
	"github.com/labstack/echo/v4"
)

var sessionID int

func main() {
	fmt.Println("Launching...")

	razer.CreateApp()

	go razer.PingHeartbeat()

	fmt.Println("Got session", sessionID)

	// TODO Find a better way to do this instead of sleeping
	fmt.Println("Waiting...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done waiting")

	/*
		for range time.Tick(time.Millisecond * 100) {
			parsedColor, _ := colorx.ParseHexColor("#34ebd8")
			createAndApplyEffect(convertColor(int(parsedColor.R), int(parsedColor.G), int(parsedColor.B)))
		}
	*/

	fmt.Println("Starting server...")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// TODO Make a doc page for the default handler
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/color/:color", handlers.ChangeColor)
	e.Logger.Fatal(e.Start(":1323"))
}
