//go:generate goversioninfo -icon=FilesDIR.ico -manifest=FilesDIR.exe.manifest
package main

import (
	_ "FilesDIR/__init__"
	"FilesDIR/construct"
	"FilesDIR/globals"
	"FilesDIR/loger"
	"FilesDIR/pkg"
	"FilesDIR/task"
	"bufio"
	"flag"
	"fmt"
	"os"
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
	// Flag of Packages
	FlgCls := flag.Bool("cls", false, "Nettoie les dossiers logs, dumps et exports")
	FlgCompiler := flag.Bool("c", false, "Lance le packager de compilation")
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
		FlgCls:       *FlgCls,
		FlgCompiler:  *FlgCompiler,
	}

	s := task.Search{
		SrcPath: *FlgPath,
		DstPath: filepath.Join(globals.FolderExports),
	}

	f.DrawStart()

	if *FlgCls {
		pkg.ClsTempFiles()
		construct.DrawEndCls()
	} else if *FlgCompiler {
		pkg.CompilerFicheAppuiFt(s.SrcPath)
	} else {

		timerStart := time.Now()
		s.RunSearch(&f)
		timerEnd := time.Since(timerStart)

		f.DrawEnd(s.SrcPath, s.DstPath, s.ReqFinal, s.NbGoroutine, s.NbFiles, s.TimerSearch, timerEnd)
	}

	fmt.Print("<cyan>Appuyer sur Entrée pour quitter...</>")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		loger.Crashln(err)
	}
}
