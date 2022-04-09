package construct

import (
	"time"
)

type Flags struct {
	FlgMode      string
	FlgWord      string
	FlgExt       string
	FlgPoolSize  int
	FlgPath      string
	FlgMaj       bool
	FlgXl        bool
	FlgDevil     bool
	FlgSuper     bool
	FlgBlackList bool
}

type Sch struct {
	SrcPath      string
	DstPath      string
	NbFiles      int
	NbFilesTotal int
	NbGoroutine  int

	Mode      string
	Word      string
	BlackList []string
	Ext       string
	Maj       bool

	TimerSearch time.Duration
	ReqFinal    string
}

type exportData struct {
	Id       int    `json:"id"`
	File     string `json:"Fichier"`
	Date     string `json:"Date"`
	PathFile string `json:"Lien_Fichier"`
	Path     string `json:"Lien"`
}
