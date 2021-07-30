package main

import (
	"bytes"
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

func makeRequest(method string, url string, body string) string {
	fmt.Println("URL:>", url)

	// TODO Do I need to do anything special to handle not passing a body?
	var jsonStr = []byte(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(respBody))

	return body
}

func pingHeartbeat() {
	// TODO Make a way to end this
	for range time.Tick(time.Second * 1) {
		resp, err := makeRequest(http.MethodPut, getSessionURL()+"/heartbeat")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getSessionURL() string {
	return baseURL + "/sid=" + sessionID
}

func createApp() {
	app := appCreationRequest{
		Title:       "Razer Chroma Go Wrapper",
		Description: "Poots",
		Author: {
			Name:    "Jesse Millar",
			Contact: "jessemillar.com",
		},
		DeviceSupported: []string{
			"keyboard",
			"mouse",
			"headset",
			"mousepad",
			"keypad",
			"chromalink",
		},
		Category: "application",
	}

	resp, err := makeRequest(http.MethodPost, getSessionURL()+"/razer/chromasdk", app)
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
	resp, err := makeRequest(http.MethodPost, getSessionURL()+"/chromalink", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func applyEffect() {
	resp, err := makeRequest(http.MethodPut, getSessionURL()+"/chromalink", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
