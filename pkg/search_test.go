package pkg

import (
	"FilesDIR/config"
	"FilesDIR/rgb"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestRunSearch(t *testing.T) {
	tabs := [][]string{
		//Mode | Word | Ext | PoolSize | Maj | Devil | Super | BlackList | WhiteList | Result
		{"%", "", "*", "10", "false", "false", "false", "false", "false", "61"},
		{"%", "comac", "*", "10", "false", "false", "false", "false", "false", "24"},
		{"%", "comac", "*", "10", "true", "false", "false", "false", "false", "6"},
		{"%", "", "*", "10", "false", "false", "false", "true", "false", "25"},
	}

	s := &Search{
		Cls:      false,
		Compiler: false,

		SrcPath: "C:\\Users\\XD5965\\go\\src\\FilesDIR\\test",
		DstPath: config.DstPath,
		Timer:   &Timer{},
		Counter: &Counter{},
	}

	for i, tab := range tabs {

		fmt.Printf(`
==================         TEST N°%v         ==================
DATA:   Mode=%s  Word=%s  Ext=%s  PoolSize=%s  Maj=%s  Devil=%s  Super=%s  BlackList=%s  WhiteList=%s
`, i+1, tab[0], tab[1], tab[2], tab[3], tab[4], tab[5], tab[6], tab[7], tab[8])

		s.Mode = tab[0]
		s.Word = tab[1]
		s.Ext = tab[2]
		s.PoolSize, _ = strconv.Atoi(tab[3])
		s.Maj, _ = strconv.ParseBool(tab[4])
		s.Devil, _ = strconv.ParseBool(tab[5])
		s.Silent, _ = strconv.ParseBool(tab[6])
		s.BlackList, _ = strconv.ParseBool(tab[7])
		s.WhiteList, _ = strconv.ParseBool(tab[8])
		s.Counter.NbrFiles = 0
		s.Counter.NbrAllFiles = 0
		Result, _ := strconv.Atoi(tab[9])

		s.RunSearch()
		time.Sleep(1 * time.Second)

		if s.Counter.NbrFiles != uint64(Result) {
			t.Error(rgb.RedBg.Sprintf("the number of files found is incorrect: %v found but %v file was expected", s.Counter.NbrFiles, Result))
		}
	}
}
