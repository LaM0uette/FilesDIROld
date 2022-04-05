package main

import (
	"FilesDIR/globals"
	"FilesDIR/task"
	"log"
)

func main() {

	task.DrawStart()

	path := globals.SrcPath
	err := task.LoopDir(path)
	if err != nil {
		log.Fatal(err)
	}
}
