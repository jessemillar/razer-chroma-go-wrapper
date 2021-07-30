package main

import (
	"fmt"
	"log"
	"net/http"
)

const baseURL = "https://chromasdk.io:54236"

var sessionID string

type appCreationRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      struct {
		Name    string `json:"name"`
		Contact string `json:"contact"`
	} `json:"author"`
	DeviceSupported []string `json:"device_supported"`
	Category        string   `json:"category"`
}

type appCreationResponse struct {
	Sessionid int    `json:"sessionid"`
	URI       string `json:"uri"`
}

type effectApplyRequest struct {
	id string `json:"id"`
}

func main() {
	fmt.Println("vim-go")

	// TODO Create app
	// TODO Make a Goroutine to ping health
	// TODO Create lighting effect
	// TODO Apply lighting effect
	// TODO Test latency/request limits
}

func makeRequest(url string, payload string) {
	resp, err := http.Get(url)
	if err != nil {
		// TODO Handle errors instead of panic
		log.Fatalln(err)
	}
}

// Need to ping more often than every 15 seconds
func pingHeartbeat() {
	resp, err := http.Get(getSessionURL() + "/heartbeat")
	if err != nil {
		log.Fatalln(err)
	}
}

func getSessionURL() string {
	return baseURL + "/sid=" + sessionID
}

func createApp() {
	makeRequest(getSessionURL()+"/razer/chromasdk", nil)
}

func createEffect() {
	makeRequest(getSessionURL()+"/chromalink", nil)
}
