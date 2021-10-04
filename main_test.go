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
	if color != "#1A535C" {
		t.Log(color)
		t.Fail()
	}
}
