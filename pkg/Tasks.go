package pkg

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"os"
	"strconv"
	"strings"
	"time"
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

func CleenTempFiles() {
	_ = os.RemoveAll(config.LogsPath)
	_ = os.RemoveAll(config.DumpsPath)
	_ = os.RemoveAll(config.ExportsPath)
}

func ExcelDateToDate(excelDate string) time.Time {

	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(excelDate)
	return excelTime.Add(time.Second * time.Duration(days*86400))
}

func (flg *SSearch) SetTimerEnd() {
	flg.Timer.AppEnd = time.Since(flg.Timer.AppStart)
}
