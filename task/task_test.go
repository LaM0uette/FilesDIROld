package task

import "testing"

func TestLoopDir(t *testing.T) {
	path := "C:\\Users\\doria\\go\\src\\FilesDIR\\tests"
	err := LoopDir(path)
	if err != nil {
		t.Error(err)
	}
}
