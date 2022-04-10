package main

import (
	_ "FilesDIR/__init__"
	"FilesDIR/construct"
	"FilesDIR/globals"
	"FilesDIR/task"
	"flag"
	"path/filepath"
	"time"
)

func main() {

	// Flag of search
	FlgMode := flag.String("mode", "%", "Mode de recherche")
	FlgWord := flag.String("word", "", "Non de fichier")
	FlgExt := flag.String("ext", "*", "Ext de fichier")
	FlgPoolSize := flag.Int("poolsize", 10, "Nombre de tâches en simultanées")
	FlgPath := flag.String("path", task.CurrentDir(), "Chemin de recherche")
	// Flag of criteral of search
	FlgMaj := flag.Bool("maj", false, "Autorise les majuscules")
	FlgXl := flag.Bool("xl", false, "Lance l'export Excel à la fin")
	// Flag of special mode
	FlgDevil := flag.Bool("devil", false, "Mode 'Démon' de l'application")
	FlgSuper := flag.Bool("s", false, "Mode 'Super', évite toutes les choses inutiles")
	FlgBlackList := flag.Bool("b", false, "Ajout d'une blacklist de dossier")
	FlgClear := flag.Bool("clear", false, "Nettoie les dossiers logs, dumps et exports")
	// Parse all Flags
	flag.Parse()

	f := construct.Flags{
		FlgMode:      *FlgMode,
		FlgWord:      *FlgWord,
		FlgExt:       *FlgExt,
		FlgPoolSize:  *FlgPoolSize,
		FlgPath:      *FlgPath,
		FlgMaj:       *FlgMaj,
		FlgXl:        *FlgXl,
		FlgDevil:     *FlgDevil,
		FlgSuper:     *FlgSuper,
		FlgBlackList: *FlgBlackList,
		FlgClear:     *FlgClear,
	}

	f.ClearTempFiles()

	s := task.Search{
		SrcPath: *FlgPath,
		DstPath: filepath.Join(globals.FolderExports),
	}

	f.DrawStart()

	timerStart := time.Now()

	s.RunSearch(&f)

	timerEnd := time.Since(timerStart)

	f.DrawEnd(s.SrcPath, s.DstPath, s.ReqFinal, s.NbGoroutine, s.NbFiles, s.TimerSearch, timerEnd)
}
