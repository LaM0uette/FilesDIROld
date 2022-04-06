package main

import (
	"FilesDIR/task"
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()

	task.DrawStart()

	task.Run()

	fmt.Println("FINI: Nb Fichiers: ", task.Id)

	timeEnd := time.Since(timeStart)
	fmt.Println(timeEnd)
}
