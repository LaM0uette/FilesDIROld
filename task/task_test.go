package task

import (
	"FilesDIR/globals"
	"fmt"
	"sync"
	"testing"
)

func TestLoopDir(t *testing.T) {
	path := globals.SrcPath
	err := LoopDir(path)
	if err != nil {
		t.Error(err)
	}
}

func TestLoopDirsFiles(t *testing.T) {
	path := globals.SrcPath
	var wg sync.WaitGroup

	err := LoopDirsFiles(path, &wg)
	if err != nil {
		t.Error(err)
	}

	wg.Wait()
	fmt.Println("FINI: Nb Fichiers: ", Id)
}

func TestLoopAlls(t *testing.T) {
	path := globals.SrcPath
	var wg sync.WaitGroup

	err := LoopAlls(path, &wg)
	if err != nil {
		t.Error(err)
	}

	wg.Wait()
	fmt.Println("FINI: Nb Fichiers: ", Id)
}
