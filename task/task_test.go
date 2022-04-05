package task

import (
	"FilesDIR/globals"
	"testing"
)

func TestLoopDir(t *testing.T) {
	path := globals.SrcPath
	err := LoopDir(path)
	if err != nil {
		t.Error(err)
	}
}
