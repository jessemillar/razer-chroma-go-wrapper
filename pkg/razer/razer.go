package razer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/icza/gox/imagex/colorx"
	"github.com/jessemillar/razer-chroma-go-wrapper/internal/utils"
)

const baseURL = "https://chromasdk.io:54236"

var defaultColor string
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
	app := appCreationRequest{
		Title:       "Razer Chroma HTTP Wrapper",
		Description: "An HTTP server wrapper around the Razer Chroma REST API to enable easy scripting",
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

func SetDefaultColor(color string) {
	defaultColor = color
}

func SetColor(color string) {
	if len(color) == 0 {
		color = defaultColor
	}

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

func FlashColor(color string, flashCount string, flashDuration string, flashInterval string) {
	// Match interval to duration if only duration is supplied
	if len(flashDuration) > 0 && len(flashInterval) == 0 {
		flashInterval = flashDuration
	}

	defaultFlashCount := 3
	defaultFlashDuration := 500
	defaultFlashInterval := 500

	flashCountInt := utils.StringToInt(flashCount, defaultFlashCount)
	flashDurationInt := utils.StringToInt(flashDuration, defaultFlashDuration)
	flashIntervalInt := utils.StringToInt(flashInterval, defaultFlashInterval)

	if flashCountInt == 0 {
		SetColor(color)
		fmt.Println("Setting color to " + color)
	} else {
		// Use an anonymous func to allow a quick HTTP return to the client
		go func() {
			for i := 0; i < flashCountInt; i++ {
				fmt.Printf("Setting color to %s for %d\n", color, time.Duration(flashDurationInt)*time.Millisecond)
				SetColor(color)
				time.Sleep(time.Duration(flashDurationInt) * time.Millisecond)

				fmt.Printf("Setting color to %s for %d\n", "black", time.Duration(flashIntervalInt)*time.Millisecond)
				SetColor("000000")
				time.Sleep(time.Duration(flashIntervalInt) * time.Millisecond)
			}

			SetColor("")
		}()
	}
}
