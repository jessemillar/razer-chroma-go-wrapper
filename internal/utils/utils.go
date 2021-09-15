package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jessemillar/razer-chroma-go-wrapper/pkg/razer"
	"github.com/spf13/viper"
)

func ConvertColor(r int, g int, b int) int {
	return ((b << 16) | (g << 8) | (r << 0))
}

func StructToBytes(theStruct interface{}) []byte {
	resultString, err := json.Marshal(theStruct)
	if err != nil {
		panic(err)
	}

	return resultString
}

func StringToInt(inputString string, defaultValue int) int {
	i, err := strconv.Atoi(inputString)
	if err != nil {
		return defaultValue
	}

	return i
}

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

func ReadConfigFile() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func FlashColor(color string, flashCount string, flashDuration string, flashInterval string) {
	defaultFlashCount := 5
	defaultFlashDuration := 1000
	defaultFlashInterval := 1500

	flashCountInt := StringToInt(flashCount, defaultFlashCount)
	flashDurationInt := StringToInt(flashDuration, defaultFlashDuration)
	flashIntervalInt := StringToInt(flashInterval, defaultFlashInterval)

	if flashCountInt == 0 {
		razer.SetColor(color)
		fmt.Println("Setting color to " + color)
	} else {
		// Use an anonymous func to allow a quick HTTP return to the client
		go func() {
			for i := 0; i < flashCountInt; i++ {
				fmt.Printf("Setting color to %s for %d\n", color, time.Duration(flashDurationInt)*time.Millisecond)
				razer.SetColor(color)
				time.Sleep(time.Duration(flashDurationInt) * time.Millisecond)
				fmt.Printf("Setting color to %s for %d\n", "black", time.Duration(flashIntervalInt)*time.Millisecond)
				razer.SetColor("000000")
				time.Sleep(time.Duration(flashIntervalInt) * time.Millisecond)
			}

			razer.SetColor()
		}()
	}
}
