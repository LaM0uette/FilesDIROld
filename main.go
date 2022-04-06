package main

import (
	"FilesDIR/task"
	"fmt"
	"time"
)

func main() {

	task.DrawStart()

	timerStart := time.Now()

	task.RunSearch(10)

	timerEnd := time.Since(timerStart)

	fmt.Println("FINI: Nb Fichiers: ", task.Id, " en ", timerEnd)
}
