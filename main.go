package main

import (
	"fmt"
	"log"
	"net/http"
)

const baseURL = "https://chromasdk.io:54236"
const url = "https://chromasdk.io:54236"

var sessionID string

func main() {
	fmt.Println("vim-go")
}

func makeRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// TODO Handle errors instead of panic
		log.Fatalln(err)
	}
}

func pingHeartbeat() {
	resp, err := http.Get(getSessionURL + "/heartbeat")
	if err != nil {
		log.Fatalln(err)
	}
}

func getSessionURL() string {
	return baseURL + "/sid=" + sessionID
}
