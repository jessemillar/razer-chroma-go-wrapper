package main

import (
	"encoding/json"
)

func convertColor(r int, g int, b int) int {
	return ((b << 16) | (g << 8) | (r << 0))
}

func structToBytes(theStruct interface{}) []byte {
	resultString, err := json.Marshal(theStruct)
	if err != nil {
		panic(err)
	}

	return resultString
}
