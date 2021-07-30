package main

import (
	"fmt"
	"log"
	"net/http"
)

const baseURL = "https://chromasdk.io:54236"

var sessionID string

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
