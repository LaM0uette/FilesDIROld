package main

import (
	"FilesDIR/globals"
	"FilesDIR/task"
	"fmt"
	"time"
)

func main() {

	task.DrawStart()

	timerStart := time.Now()

	s := task.Sch{
		SrcPath:  globals.SrcPath,
		PoolSize: 10,
		NbFiles:  0,
	}

	task.RunSearch(&s)

	timerEnd := time.Since(timerStart)

	fmt.Println("FINI: Nb Fichiers: ", s.NbFiles, " en ", timerEnd)
}
