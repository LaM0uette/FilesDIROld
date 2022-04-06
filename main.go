package main

import (
	"FilesDIR/globals"
	"FilesDIR/task"
	"time"
)

func main() {

	task.DrawStart()

	timerStart := time.Now()

	s := task.Sch{
		SrcPath:  globals.SrcPathGen,
		PoolSize: 10,
		NbFiles:  0,
	}

	task.RunSearch(&s)

	timerEnd := time.Since(timerStart)

	task.DrawEnd(&s, timerEnd)
}
