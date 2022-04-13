package pkg

import (
	"testing"
)

func TestGetCurrentDir(t *testing.T) {
	dir := GetCurrentDir()

	if dir != "C:\\Users\\XD5965\\go\\src\\FilesDIR\\pkg" {
		t.Error("Is not the current dir ! ", dir)
	}
}

func TestStrToLower(t *testing.T) {
	word := "ImATest"

	if StrToLower(word) != "imatest" {
		t.Error("Fail during strings conversion !", word)
	}
}
