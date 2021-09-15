package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
