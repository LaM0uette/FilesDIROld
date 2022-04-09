package main

import (
	_ "FilesDIR/__init__"
	"FilesDIR/display"
	"FilesDIR/globals"
	"FilesDIR/log"
	"FilesDIR/task"
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	log.Bln("test")

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
	flag.Parse() // Parse all Flags

	f := task.Flags{
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
	}

	if !f.FlgSuper {

		log.Blank.Print(display.DrawStart())
		fmt.Print(display.DrawStart())

		time.Sleep(1 * time.Second)
	}

	timerStart := time.Now()

	s := task.Sch{
		SrcPath: *FlgPath,
		DstPath: filepath.Join(globals.TempPathGen, "exports"),
	}

	task.RunSearch(&s, &f)

	timerEnd := time.Since(timerStart)

	disp := display.DrawEnd(s.SrcPath, s.DstPath, s.ReqFinal, s.NbGoroutine, s.NbFiles, f.FlgPoolSize, s.TimerSearch, timerEnd)
	log.Blank.Print(disp)
	fmt.Print(disp)

	fmt.Print("Appuyer sur Entrée pour quitter...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		log.Crash.Println(err)
	}
}
