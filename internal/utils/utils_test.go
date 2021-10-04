package utils

import (
	"testing"
)

func TestCustomColor(t *testing.T) {
	ReadConfigFile()
	color := TranslateCustomColor("pink")
	if color != "#f3b8b6" {
		t.Fail()
	}
}

func TestCustomColorAlias(t *testing.T) {
	ReadConfigFile()
	color := TranslateCustomColor("success")
	if color != "#4ECDC4" {
		t.Fail()
	}
}
