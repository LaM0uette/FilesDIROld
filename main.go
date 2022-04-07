package main

import (
	_ "FilesDIR/__init__"
	"FilesDIR/globals"
	"FilesDIR/log"
	"FilesDIR/task"
	"flag"
	"fmt"
	"path/filepath"
	"time"
)

func main() {

	FlgDevil := flag.Bool("devil", false, "Mode 'Démon' de l'application")
	FlgMode := flag.String("mode", "%", "Mode de recherche")
	FlgWord := flag.String("word", "", "Non de fichier")
	FlgExt := flag.String("ext", "*", "Ext de fichier")
	FlgMaj := flag.Bool("maj", false, "Autorise les majuscules")
	FlgXlsx := flag.Bool("xl", false, "Lance l'export Excel à la fin")
	flag.Parse()

	f := task.Flags{
		FlgDevil: *FlgDevil,
		FlgMode:  *FlgMode,
		FlgWord:  *FlgWord,
		FlgExt:   *FlgExt,
		FlgMaj:   *FlgMaj,
		FlgXl:    *FlgXlsx,
	}

	log.Blank.Print(task.DrawStart())
	fmt.Print(task.DrawStart())

	timerStart := time.Now()

	s := task.Sch{
		SrcPath:  globals.SrcPathGen,
		DstPath:  filepath.Join(globals.TempPathGen, "exports"),
		PoolSize: 10,
	}

	task.RunSearch(&s, &f)

	timerEnd := time.Since(timerStart)

	log.Blank.Print(task.DrawEnd(&s, s.TimerSearch, timerEnd))
	fmt.Print(task.DrawEnd(&s, s.TimerSearch, timerEnd))

	//log.Blank.Print(DrawSaveExcel.Pattern)
	//DrawSaveExcel.Draw()
}
