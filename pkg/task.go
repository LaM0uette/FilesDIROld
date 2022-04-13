package pkg

import (
	"os"
	"strings"
)

func GetCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		//TODO: mettre un loger
		os.Exit(1)
	}
	return pwd
}

func StrToLower(s string) string {
	return strings.ToLower(s)
}
