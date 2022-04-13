package pkg

import (
	"os"
)

func GetCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		//TODO: mettre un loger
		os.Exit(1)
	}
	return pwd
}
