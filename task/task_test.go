package task

import (
	"testing"
)

const scrTest = "C:\\Users\\doria\\go\\src\\FilesDIR\\tests"

func TestRunSearch(t *testing.T) {
	s := Sch{
		SrcPath:  scrTest,
		PoolSize: 10,
		NbFiles:  0,
	}

	RunSearch(&s)
	if s.NbFiles != 18 {
		t.Error("NbFiles is not 18", s.NbFiles)
	}
}

func BenchmarkRunSearch(b *testing.B) {
	s := Sch{
		SrcPath:  scrTest,
		PoolSize: 10,
		NbFiles:  0,
	}

	for i := 0; i < b.N; i++ {
		RunSearch(&s)
		if s.NbFiles != 18 {
			b.Error("NbFiles is not 18", s.NbFiles)
		}
	}
}
