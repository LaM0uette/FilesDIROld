package main

import (
	"FilesDIR/globals"
	"FilesDIR/task"
	"fmt"
	"log"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	timeStart := time.Now()

	var wg sync.WaitGroup

	task.DrawStart()

	debug.SetMaxThreads(5 * 1000)

	/*
		err := task.LoopDir(globals.SrcPath)
		if err != nil {
			log.Print(err.Error())
		}
	*/

	err := task.LoopDirsFiles(globals.SrcPath, &wg)
	if err != nil {
		log.Print(err.Error())
	}

	wg.Wait()

	fmt.Println("FINI: Nb Fichiers: ", task.Id)

	timeEnd := time.Since(timeStart)
	fmt.Println(timeEnd)
}
