package main

import (
	"FilesDIR/task"
	"log"
)

func main() {

	task.DrawStart()

	path := "C:\\Users\\doria\\go\\src\\task\\tests"
	err := task.LoopDir(path)
	if err != nil {
		log.Fatal(err)
	}
}
