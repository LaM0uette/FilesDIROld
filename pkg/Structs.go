package pkg

import "time"

type STimer struct {
	AppStart time.Time
	AppEnd   time.Duration

	SearchStart time.Time
	SearchEnd   time.Duration
}

type SCounter struct {
	NbrFiles    uint64
	NbrAllFiles uint64

	NbrFolder uint64
}

type SProcess struct {
	NbrThreads    int
	NbrGoroutines int
}

type SSearch struct {

	// Flags
	Cls      bool
	Compiler bool

	//..
	Mode      string
	Word      string
	Ext       string
	PoolSize  int
	Maj       bool
	Devil     bool
	Silent    bool
	WordsList bool
	BlackList bool
	WhiteList bool

	// SSearch
	SrcPath string
	DstPath string
	ReqUse  string

	// Data
	ListWords     []string
	ListBlackList []string
	ListWhiteList []string
	Timer         *STimer
	Counter       *SCounter
	Process       *SProcess
}
