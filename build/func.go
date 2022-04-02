package build

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Search struct {
	Mode      string
	Word      string
	Path      string
	Extension string
}

// CurrentDir : Return the actual directory
func CurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func (s *Search) SearchFiles() error {

	id := 0

	err := filepath.Walk(s.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() == false && strings.Contains(strings.ToLower(path), strings.ToLower(s.Word)) {

			fileStat, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Nom: %v | Id: %v\n", fileStat.Name(), id)
			id++
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
