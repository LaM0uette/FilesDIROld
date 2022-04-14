package pkg

import (
	"FilesDIR/config"
	"fmt"
	"strconv"
	"testing"
)

func TestRunSearch(t *testing.T) {
	tabs := [][]string{
		//Mode | Word | Ext | PoolSize | Maj | Devil | Super | BlackList | WhiteList | Result
		{"%", "", "*", "10", "false", "false", "false", "false", "false", "61"},
		{"%", "comac", "*", "10", "false", "false", "false", "false", "false", "24"},
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
		VDevil, _ := strconv.ParseBool(tab[5])
		VSuper, _ := strconv.ParseBool(tab[6])
		VBlackList, _ := strconv.ParseBool(tab[7])
		VWhiteList, _ := strconv.ParseBool(tab[8])
		VResult, _ := strconv.Atoi(tab[9])

		s := &Search{
			Cls:       false,
			Compiler:  false,
			Mode:      VMode,
			Word:      VWord,
			Ext:       VExt,
			PoolSize:  VPoolSize,
			Maj:       VMaj,
			Devil:     VDevil,
			Silent:    VSuper,
			BlackList: VBlackList,
			WhiteList: VWhiteList,

			SrcPath: "C:\\Users\\XD5965\\go\\src\\FilesDIR\\test",
			DstPath: config.DstPath,
			Timer:   &Timer{},
			Counter: &Counter{},
		}

		s.RunSearch()

		if s.Counter.NbrFiles != uint64(VResult) {
			t.Error(fmt.Sprintf("the number of files found is incorrect: %v found but %v file was expected", s.Counter.NbrFiles, VResult))
		}

		fmt.Println()
		fmt.Println()
		fmt.Println(s.Counter.NbrFiles)
		fmt.Println()
		fmt.Println()
	}
}
