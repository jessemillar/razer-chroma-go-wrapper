package main

import (
	"testing"

	"github.com/jessemillar/razer-chroma-http-wrapper/internal/utils"
	"github.com/spf13/viper"
)

func TestCustomColor(t *testing.T) {
	utils.ReadConfigFile()
	t.Log(viper.GetString("default_color"))
	color := utils.TranslateCustomColor("pink")
	if color != "#f3b8b6" {
		t.Log(color)
		t.Fail()
	}
}
