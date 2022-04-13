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
