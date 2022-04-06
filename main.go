package main

import (
	"FilesDIR/globals"
	"FilesDIR/log"
	"FilesDIR/task"
	"time"
)

func main() {

	task.DrawStart()
	log.Info.Println("Starting FilesDIR")

	log.Info.Println("Test Info")
	log.Warning.Println("Test Warning")
	log.Error.Println("Test Error")

	timerStart := time.Now()

	s := task.Sch{
		SrcPath:  globals.SrcPathGen,
		PoolSize: 10,
		NbFiles:  0,
	}

	log.Info.Println("Starting search")
	task.RunSearch(&s)
	log.Info.Println("Ending search")

	timerEnd := time.Since(timerStart)

	task.DrawEnd(&s, timerEnd)
	log.Info.Println("Ending FilesDIR")
}
