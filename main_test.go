package main

import (
	"testing"

	"github.com/jessemillar/razer-chroma-http-wrapper/internal/utils"
)

func TestCustomColor(t *testing.T) {
	utils.ReadConfigFile()
	color := utils.TranslateCustomColor("pink")
	if color != "#f3b8b6" {
		t.Fail()
	}
}

func TestCustomColorAlias(t *testing.T) {
	utils.ReadConfigFile()
	color := utils.TranslateCustomColor("success")
	if color != "#4ECDC4" {
		t.Fail()
	}
}
