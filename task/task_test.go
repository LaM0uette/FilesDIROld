package task

import (
	"FilesDIR/construct"
	"fmt"
	"strconv"
	"testing"
)

const (
	scrTest = "C:\\Users\\doria\\go\\src\\FilesDIR\\tests"
	DstPath = "C:\\Users\\doria\\go\\src\\FilesDIR\\export"
)

func TestRunSearch(t *testing.T) {
	s := Search{
		SrcPath: scrTest,
		DstPath: DstPath,
	}

	tabs := [][]string{
		//Mode | Word | Ext | PoolSize | Maj | Xl | Devil | Super | BlackList | Result
		{"%", "", "*", "10", "false", "true", "false", "false", "false", "26"},
		{"%", "Devis", "*", "10", "false", "true", "false", "false", "false", "6"},
		{"%", "Devis", "*", "10", "true", "true", "false", "false", "false", "6"},
		{"=", "Xl", "xlsx", "10", "true", "true", "false", "false", "false", "1"},
		{"%", "x", "txt", "10", "false", "true", "false", "false", "false", "3"},
		{"%", "", "txt", "10", "false", "true", "true", "false", "false", "21"},
		{"%", "", "txt", "10", "false", "true", "false", "false", "false", "21"},
	}

	for i, tab := range tabs {

		fmt.Printf(`
==================         TEST N°%v         ==================
DATA:   Mode=%s  Word=%s  Ext=%s  PoolSize=%s  Maj=%s  Xl=%s  Devil=%s  Super=%s  BlackList=%s

`, i+1, tab[0], tab[1], tab[2], tab[3], tab[4], tab[5], tab[6], tab[7], tab[8])

		VMode := tab[0]
		VWord := tab[1]
		VExt := tab[2]
		VPoolSize, _ := strconv.Atoi(tab[3])
		VMaj, _ := strconv.ParseBool(tab[4])
		VXl, _ := strconv.ParseBool(tab[5])
		VDevil, _ := strconv.ParseBool(tab[6])
		VSuper, _ := strconv.ParseBool(tab[7])
		VBlackList, _ := strconv.ParseBool(tab[8])
		VResult, _ := strconv.Atoi(tab[9])

		f := construct.Flags{
			FlgMode:      VMode,
			FlgWord:      VWord,
			FlgExt:       VExt,
			FlgPoolSize:  VPoolSize,
			FlgMaj:       VMaj,
			FlgXl:        VXl,
			FlgDevil:     VDevil,
			FlgSuper:     VSuper,
			FlgBlackList: VBlackList,
		}

		s.NbFiles = 0
		s.RunSearch(&f)

		if s.NbFiles != VResult {
			t.Error(fmt.Sprintf("the number of files found is incorrect: %v found but %v file was expected", s.NbFiles, VResult))
		}

		fmt.Println()
		fmt.Println()
	}
}
