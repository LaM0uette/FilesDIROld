//go:generate goversioninfo -icon=FilesDIR.ico -manifest=FilesDIR.exe.manifest
package main

import (
	"FilesDIR/config"
	"FilesDIR/pkg"
	"flag"
	"fmt"
	"time"
)

func main() {

	timerStart := time.Now()

	pkg.DrawStart()

	// Flag of Packages
	FlgCls := flag.Bool("cls", false, "Nettoie les dossiers logs, dumps et exports")
	FlgCompiler := flag.Bool("c", false, "Lance le mode de compilation")
	// Flag of search
	FlgMode := flag.String("mode", "%", "Mode de recherche")
	FlgWord := flag.String("word", "", "Non de fichier")
	FlgExt := flag.String("ext", "*", "Ext de fichier")
	FlgPoolSize := flag.Int("poolsize", 10, "Nombre de tâches en simultanées")
	// Flag of criteral of search
	FlgMaj := flag.Bool("maj", false, "Autorise les majuscules")
	// Flag of special mode
	FlgDevil := flag.Bool("devil", false, "Mode 'Démon' de l'application")
	FlgSuper := flag.Bool("s", false, "Mode 'Super', évite toutes les choses inutiles")
	FlgBlackList := flag.Bool("b", false, "Ajout d'une blacklist de dossier")
	// Parse all Flags
	flag.Parse()

	s := &pkg.Search{
		Cls:       *FlgCls,
		Compiler:  *FlgCompiler,
		Mode:      *FlgMode,
		Word:      *FlgWord,
		Ext:       *FlgExt,
		PoolSize:  *FlgPoolSize,
		Maj:       *FlgMaj,
		Devil:     *FlgDevil,
		Super:     *FlgSuper,
		BlackList: *FlgBlackList,

		SrcPath: pkg.GetCurrentDir(),
		DstPath: config.DstPath,
	}

	if s.Cls {

	} else if s.Compiler {

	} else {
		s.RunSearch()
	}

	timerEnd := time.Since(timerStart)

	fmt.Println(timerEnd)

	pkg.DrawEnd()
}
