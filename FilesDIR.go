package main

import (
	"FilesDIR/FilesDIR"
	"log"
)

func main() {

	FilesDIR.DrawStart()

	err := FilesDIR.LoopDir("C:\\Users\\XD5965\\go\\src\\FilesDIR\\tests")
	if err != nil {
		log.Fatal(err)
	}
}
