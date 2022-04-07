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

	FlgDevil := flag.Bool("devil", false, "Mode 'Démon' de l'application.")
	flag.Parse()

	f := task.Flags{
		FlgDevil: *FlgDevil,
	}

	task.DrawStart()

	log.BlankDate.Println("*** Starting FilesDIR\n")
	timerStart := time.Now()

	s := task.Sch{
		SrcPath:  globals.SrcPathGen,
		DstPath:  filepath.Join(globals.TempPathGen, "exports"),
		PoolSize: 10,
	}

	log.BlankDate.Printf(fmt.Sprintf("*** Starting search on: %s\n\n", s.SrcPath))
	task.RunSearch(&s, &f)

	log.BlankDate.Println("\n*** Ending search\n")
	timerEnd := time.Since(timerStart)

	log.BlankDate.Println("*** Closing FilesDIR")
	task.DrawEnd(&s, s.TimerSearch, timerEnd)
}
