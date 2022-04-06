package main

import (
	"FilesDIR/globals"
	"FilesDIR/log"
	"FilesDIR/task"
	"time"
)

func main() {

	task.DrawStart()
	log.Info.Println("Starting FilesDIR\n")

	log.Info.Println("Test Info\n")
	log.Warning.Println("Test Warning\n")
	log.Error.Println("Test Error\n")

	timerStart := time.Now()

	s := task.Sch{
		SrcPath:  globals.SrcPathGen,
		PoolSize: 10,
		NbFiles:  0,
	}

	log.Info.Println("Starting search\n")
	task.RunSearch(&s)
	log.Info.Println("Ending search\n")

	timerEnd := time.Since(timerStart)

	task.DrawEnd(&s, timerEnd)
	log.Info.Println("Ending FilesDIR\n")
}
