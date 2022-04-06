package task

import (
	"testing"
)

func TestRunSearch(t *testing.T) {
	path := "C:\\Users\\doria\\go\\src\\FilesDIR\\tests"
	RunSearch(path, 10)
	if Id != 18 {
		t.Error("Id is not 18", Id)
	}
}
