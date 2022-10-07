//go:generate goversioninfo -icon=FilesDIR.ico -manifest=FilesDIR.exe.manifest
package main

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"FilesDIR/pkg"
	"FilesDIR/rgb"
	"bufio"
	"flag"
	"os"
	"time"
)

func main() {

	flg := GetFlags()

	flg.DrawStart()

	if flg.Cls {
		pkg.CleenTempFiles()
		flg.SetTimerEnd()
		flg.DrawCls()
	} else if flg.Compiler {
		flg.CompileFichesAppuis()
		flg.SetTimerEnd()
		flg.DrawFichesAppuisCompiled()
	} else {
		flg.RunSearch()
		flg.SetTimerEnd()
		flg.DrawBilanSearch()
	}

	flg.DrawEnd()

	rgb.GreenB.Print("Appuyer sur Entrée pour quitter...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		loger.Crash("Crash :", err)
	}
}

func GetFlags() *pkg.SSearch {

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
	FlgSilent := flag.Bool("s", false, "Mode 'Silent', évite toutes les choses inutiles")
	FlgWordsList := flag.Bool("wl", false, "Ajout d'une wordslist")
	FlgBlackList := flag.Bool("b", false, "Ajout d'une blacklist de dossier")
	FlgWhiteList := flag.Bool("w", false, "Ajout d'une whitelist de dossier")

	// Parse all Flags
	flag.Parse()

	return &pkg.SSearch{
		Cls:       *FlgCls,
		Compiler:  *FlgCompiler,
		Mode:      *FlgMode,
		Word:      *FlgWord,
		Ext:       *FlgExt,
		PoolSize:  *FlgPoolSize,
		Maj:       *FlgMaj,
		Devil:     *FlgDevil,
		Silent:    *FlgSilent,
		WordsList: *FlgWordsList,
		BlackList: *FlgBlackList,
		WhiteList: *FlgWhiteList,
		SrcPath:   pkg.GetCurrentDir(),
		DstPath:   config.DstPath,
		Timer:     &pkg.STimer{AppStart: time.Now()},
		Counter:   &pkg.SCounter{},
		Process:   &pkg.SProcess{},
	}
}
