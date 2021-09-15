package razer

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func MakeRequest(method string, url string, body []byte) (string, error) {
	/*
		fmt.Println("URL:", url)
		fmt.Println("Method:", method)
		fmt.Println("Body:", string(body))
	*/

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	// fmt.Println("Response:", bodyString)

	return bodyString, nil
}

func PingHeartbeat() {
	// TODO Make a way to end this
	for range time.Tick(time.Second * 1) {
		// Only ping if we have a session ID
		if sessionID > 0 {
			_, err := makeRequest(http.MethodPut, getSessionURL()+"/heartbeat", nil)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func GetSessionURL() string {
	return baseURL + "/sid=" + strconv.Itoa(sessionID)
}

func CreateApp() {
	// TODO Make these values real values
	app := appCreationRequest{
		Title:       "Razer Chroma Go Wrapper",
		Description: "Poots",
		Author: appCreationRequestAuthor{
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

	resp, err := makeRequest(http.MethodPost, baseURL+"/razer/chromasdk", structToBytes(app))
	if err != nil {
		log.Fatalln(err)
	}

	var data appCreationResponse
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		panic(err.Error())
	}

	sessionID = data.SessionID
}

func CreateAndApplyEffect(color int) {
	effect := effectCreationRequest{
		Effect: "CHROMA_STATIC",
		Param: effectParam{
			Color: color,
		},
	}

	effectID := createEffect(effect)
	applyEffect(effectID)
}

func CreateEffect(effect effectCreationRequest) string {
	resp, err := makeRequest(http.MethodPost, getSessionURL()+"/chromalink", structToBytes(effect))
	if err != nil {
		log.Fatalln(err)
	}

	var data effectCreationResponse
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		panic(err.Error())
	}

	return data.ID
}

func ApplyEffect(effectID string) {
	requestBody := effectApplyRequest{
		ID: effectID,
	}

	_, err := makeRequest(http.MethodPut, getSessionURL()+"/effect", structToBytes(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
}
