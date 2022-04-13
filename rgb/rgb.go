package rgb

import (
	"github.com/gookit/color"
)

var (
	Green    = color.NewRGBStyle(color.RGB(44, 168, 65))
	Yellow   = color.NewRGBStyle(color.RGB(196, 168, 27)).SetOpts(color.Opts{color.OpUnderscore, color.Bold})
	BgYellow = color.NewRGBStyle(color.RGB(240, 240, 240), color.RGB(196, 168, 27))

	Gray = color.NewRGBStyle(color.RGB(60, 60, 60))
)
