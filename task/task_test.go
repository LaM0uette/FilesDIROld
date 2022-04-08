package task

import (
	"fmt"
	"strconv"
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

	tabs := [][]string{
		//Devil  Mode Word Ext    Maj      Xl   Result
		{"false", "%", "", "*", "false", "true", "26"},
		{"false", "%", "Devis", "*", "false", "true", "6"},
		{"false", "%", "Devis", "*", "true", "true", "5"},
		{"false", "=", "Xl", "xlsx", "true", "true", "1"},
		{"false", "%", "x", "txt", "false", "true", "3"},
		{"false", "%", "", "txt", "false", "true", "21"},
		{"devil", "%", "", "txt", "false", "true", "21"},
	}

	for i, tab := range tabs {

		fmt.Printf(`
==================         TEST N°%v         ==================
DATA: Devil=%s  Mode=%s  Word=%s  Ext=%s  Maj=%s  Xl=%s

`, i+1, tab[0], tab[1], tab[2], tab[3], tab[4], tab[5])

		VDevil, _ := strconv.ParseBool(tab[0])
		VMaj, _ := strconv.ParseBool(tab[4])
		VXl, _ := strconv.ParseBool(tab[5])
		VResult, _ := strconv.Atoi(tab[6])

		f := Flags{
			FlgDevil: VDevil,
			FlgMode:  tab[1],
			FlgWord:  tab[2],
			FlgExt:   tab[3],
			FlgMaj:   VMaj,
			FlgXl:    VXl,
		}

		s.NbFiles = 0
		RunSearch(&s, &f)

		if s.NbFiles != VResult {
			t.Error(fmt.Sprintf("the number of files found is incorrect: %v found but %v file was expected", s.NbFiles, VResult))
		}

		fmt.Println()
		fmt.Println()
	}
}
