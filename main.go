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

	FlgMode := flag.String("mode", "%", "Mode de recherche")
	FlgWord := flag.String("word", "", "Non de fichier")
	FlgExt := flag.String("ext", "*", "Ext de fichier")

	FlgMaj := flag.Bool("maj", false, "Autorise les majuscules")
	FlgXl := flag.Bool("xl", false, "Lance l'export Excel à la fin")
	FlgDevil := flag.Bool("devil", false, "Mode 'Démon' de l'application")
	FlgSuper := flag.Bool("s", false, "Mode 'Super', évite toutes les choses inutiles")
	FlgBlackList := flag.Bool("b", false, "Ajout d'une blacklist de dossier")
	flag.Parse()

	f := task.Flags{
		FlgMode:      *FlgMode,
		FlgWord:      *FlgWord,
		FlgExt:       *FlgExt,
		FlgMaj:       *FlgMaj,
		FlgXl:        *FlgXl,
		FlgDevil:     *FlgDevil,
		FlgSuper:     *FlgSuper,
		FlgBlackList: *FlgBlackList,
	}

	if !f.FlgSuper {
		log.Blank.Print(task.DrawStart())
		fmt.Print(task.DrawStart())
	}

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
}
