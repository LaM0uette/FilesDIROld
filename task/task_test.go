package task

import (
	"testing"
)

func TestRunSearch(t *testing.T) {
	path := "F:\\"
	err := RunSearch(path, 10)
	if err != nil {
		t.Error(err)
	}
}
