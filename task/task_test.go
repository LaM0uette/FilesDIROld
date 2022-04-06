package task

import (
	"testing"
)

func TestRunSearch(t *testing.T) {
	s := Sch{
		SrcPath:  "C:\\Users\\doria\\go\\src\\FilesDIR\\tests",
		PoolSize: 10,
		NbFiles:  0,
	}

	RunSearch(&s)
	if s.NbFiles != 18 {
		t.Error("NbFiles is not 18", Id)
	}
}
