package main

import (
	"FilesDIR/globals"
	"FilesDIR/log"
	"FilesDIR/task"
	"fmt"
	"time"
)

func main() {

	task.DrawStart()

	log.BlankDate.Println("*** Starting FilesDIR\n")
	timerStart := time.Now()

	s := task.Sch{
		SrcPath:     globals.SrcPathGen,
		PoolSize:    100000,
		NbFiles:     0,
		NbGoroutine: 0,
	}

	log.BlankDate.Printf(fmt.Sprintf("*** Starting search on: %s\n\n", s.SrcPath))
	task.RunSearch(&s)

	log.BlankDate.Println("\n*** Ending search\n")
	timerEnd := time.Since(timerStart)

	log.BlankDate.Println("*** Closing FilesDIR")
	task.DrawEnd(&s, timerEnd)
}
