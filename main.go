package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	SessionID int    `json:"sessionid"`
	URI       string `json:"uri"`
}

type effectCreateRequest struct {
	Effect string `json:"effect"`
	Param  struct {
		Color int `json:"color"`
	} `json:"param"`
}

type effectApplyRequest struct {
	id string `json:"id"`
}

func main() {
	createApp()

	go pingHeartbeat()

	fmt.Println(sessionID)

	// TODO Create lighting effect
	// TODO Apply lighting effect
	// TODO Test latency/request limits
}

func pingHeartbeat() {
	// TODO Make a way to end this
	for range time.Tick(time.Second * 1) {
		resp, err := http.Put(getSessionURL() + "/heartbeat")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getSessionURL() string {
	return baseURL + "/sid=" + sessionID
}

func createApp() {
	app := appCreationRequest{}

	resp, err := http.Post(getSessionURL()+"/razer/chromasdk", app)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var data appCreationResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	sessionID = data.SessionID
}

func createEffect() {
	resp, err := http.Post(getSessionURL()+"/chromalink", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func applyEffect() {
	// TODO Handle Put
	resp, err := http.Put(getSessionURL()+"/chromalink", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
