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

	task.RunSearch(globals.SrcPath, 10)

	timerEnd := time.Since(timerStart)

	fmt.Println("FINI: Nb Fichiers: ", task.Id, " en ", timerEnd)
}
