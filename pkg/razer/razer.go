package razer

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/icza/gox/imagex/colorx"
	"github.com/jessemillar/razer-chroma-go-wrapper/internal/utils"
)

const baseURL = "https://chromasdk.io:54236"

var sessionID int

// TODO Unexport functions that can be private

func PingHeartbeat() {
	for range time.Tick(time.Second * 1) {
		// Only ping if we have a session ID
		if sessionID > 0 {
			_, err := utils.MakeRequest(http.MethodPut, GetSessionURL()+"/heartbeat", nil)
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

	resp, err := utils.MakeRequest(http.MethodPost, baseURL+"/razer/chromasdk", utils.StructToBytes(app))
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

func SetColor(color string) {
	parsedColor, _ := colorx.ParseHexColor("#" + strings.Replace(color, "#", "", -1)) // Hack to make sure there's only one pound sign
	CreateAndApplyEffect(utils.ConvertColor(int(parsedColor.R), int(parsedColor.G), int(parsedColor.B)))
}

func CreateAndApplyEffect(color int) {
	effect := effectCreationRequest{
		Effect: "CHROMA_STATIC",
		Param: effectParam{
			Color: color,
		},
	}

	effectID := CreateEffect(effect)
	ApplyEffect(effectID)
}

func CreateEffect(effect effectCreationRequest) string {
	resp, err := utils.MakeRequest(http.MethodPost, GetSessionURL()+"/chromalink", utils.StructToBytes(effect))
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

	_, err := utils.MakeRequest(http.MethodPut, GetSessionURL()+"/effect", utils.StructToBytes(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
}
