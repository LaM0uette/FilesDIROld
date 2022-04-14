package pkg

import (
	"FilesDIR/loger"
	"os"
	"strings"
)

func GetCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		loger.Error("Error with current dir:", err)
		os.Exit(1)
	}
	return pwd
}

func StrToLower(s string) string {
	return strings.ToLower(s)
}
