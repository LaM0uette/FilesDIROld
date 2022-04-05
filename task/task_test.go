package task

import "testing"

func TestLoopDir(t *testing.T) {
	path := "C:\\Users\\XD5965\\go\\src\\task\\tests"

	err := LoopDir(path)
	if err != nil {
		t.Error(err)
	}
}
