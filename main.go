package main

import (
	"FilesDIR/globals"
	"FilesDIR/task"
	"log"
)

func main() {

	task.DrawStart()

	err := task.LoopDir(globals.SrcPath)
	if err != nil {
		log.Fatal(err)
	}
}
