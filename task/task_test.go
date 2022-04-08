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
		//Mode | Word | Ext | Maj | Xl | Devil | Super | BlackList | Result
		{"%", "", "*", "false", "true", "false", "false", "false", "26"},
		{"%", "Devis", "*", "false", "true", "false", "false", "false", "6"},
		{"%", "Devis", "*", "true", "true", "false", "false", "false", "5"},
		{"=", "Xl", "xlsx", "true", "true", "false", "false", "false", "1"},
		{"%", "x", "txt", "false", "true", "false", "false", "false", "3"},
		{"%", "", "txt", "false", "true", "true", "false", "false", "21"},
		{"%", "", "txt", "false", "true", "false", "false", "false", "21"},
	}

	for i, tab := range tabs {

		fmt.Printf(`
==================         TEST N°%v         ==================
DATA:   Mode=%s  Word=%s  Ext=%s  Maj=%s  Xl=%s  Devil=%s  Super=%s  BlackList=%s

`, i+1, tab[0], tab[1], tab[2], tab[3], tab[4], tab[5], tab[6], tab[7])

		VMode := tab[0]
		VWord := tab[1]
		VExt := tab[2]
		VMaj, _ := strconv.ParseBool(tab[3])
		VXl, _ := strconv.ParseBool(tab[4])
		VDevil, _ := strconv.ParseBool(tab[5])
		VSuper, _ := strconv.ParseBool(tab[6])
		VBlackList, _ := strconv.ParseBool(tab[7])
		VResult, _ := strconv.Atoi(tab[8])

		f := Flags{
			FlgMode:      VMode,
			FlgWord:      VWord,
			FlgExt:       VExt,
			FlgMaj:       VMaj,
			FlgXl:        VXl,
			FlgDevil:     VDevil,
			FlgSuper:     VSuper,
			FlgBlackList: VBlackList,
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
