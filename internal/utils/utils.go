package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

// TranslateCustomColor accepts a string and checks if that string is a custom color defined in our config
func TranslateCustomColor(color string) string {
	customColors := viper.GetStringMapString("custom_colors")
	fmt.Println(customColors)
	if val, ok := customColors[color]; ok {
		return val
	}

	colorAliases := viper.GetStringMapString("color_aliases")
	if alias, ok := colorAliases[color]; ok {
		if val, ok := customColors[alias]; ok {
			return val
		}
	}

	return color
}

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

func ReadConfigFile() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	viper.SetDefault("server_port", "1323")
	viper.SetDefault("default_color", "#bada55")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
