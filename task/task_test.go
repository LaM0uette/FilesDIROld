package task

import (
	"testing"
)

const (
	scrTest = "C:\\Users\\XD5965\\go\\src\\FilesDIR\\tests"
	DstPath = "C:\\Users\\XD5965\\go\\src\\FilesDIR\\export"
)

func TestRunSearch(t *testing.T) {
	s := Sch{
		SrcPath:  scrTest,
		DstPath:  DstPath,
		PoolSize: 10,
	}

	RunSearch(&s)
	if s.NbFiles != 18 {
		t.Error("NbFiles is not 18", s.NbFiles)
	}
}
