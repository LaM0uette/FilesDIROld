package build

import (
	"fmt"
	"os"
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

func (s *Search) SearchFiles() {

	fmt.Println(s)

}
