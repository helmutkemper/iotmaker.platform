package iotmaker_platform

import (
	pwb "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

type Stage struct {
	Canvas pwb.Canvas
}

func (el *Stage) LineTo(x, y float64) {
	el.Canvas.LineTo(x, y)
}
