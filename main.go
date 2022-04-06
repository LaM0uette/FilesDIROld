package main

import (
	"FilesDIR/globals"
	"FilesDIR/task"
	"fmt"
	"log"
	"os"
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

	file, err := os.OpenFile(fmt.Sprintf("logs/SLog_%v.txt", timerStart.Format("20060102150405")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Println("Hello world!")

	task.RunSearch(&s)

	timerEnd := time.Since(timerStart)

	task.DrawEnd(&s, timerEnd)
}
