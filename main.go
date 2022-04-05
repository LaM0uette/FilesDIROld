package main

import (
	"FilesDIR/globals"
	"FilesDIR/task"
	"fmt"
	"log"
	"time"
)

func main() {
	timeStart := time.Now()

	task.DrawStart()

	err := task.LoopDir(globals.SrcPath)
	if err != nil {
		log.Print(err.Error())
	}

	timeEnd := time.Since(timeStart)
	fmt.Println(timeEnd)
}
