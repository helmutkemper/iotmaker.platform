package basicBox

import (
	"fmt"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	"image/color"
)

type BasicBox struct {
	Platform   iotmaker_platform_IDraw.IDraw
	ScratchPad iotmaker_platform_IDraw.IDraw
	Id         string
	Dimensions Dimensions
	Ink        Ink

	imageData [][]color.RGBA

	prepareShadowFilter   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilter func(iotmaker_platform_IDraw.ICanvasGradient)

	x1 int
	x2 int
	x3 int
	x4 int

	y1 int
	y2 int
	y3 int
	y4 int
}

func (el *BasicBox) PrepareShadowFilter() {
	if el.prepareShadowFilter != nil {
		el.prepareShadowFilter(el.Platform)
	}
}

func (el *BasicBox) PrepareGradientFilter() {
	if el.prepareGradientFilter != nil {
		el.prepareGradientFilter(el.Platform)
	}
}

func (el *BasicBox) configShadowPlatformAndFilter() {
	if el.Ink.Shadow == nil {
		return
	}

	el.SetShadowFilter(el.Ink.Shadow.PrepareFilter)
}

func (el *BasicBox) configGradientPlatformAndFilter() {
	if el.Ink.Gradient == nil {
		return
	}

	el.SetGradientAndMountColorListFilter(el.Ink.Gradient.PrepareFilter)
}

func (el *BasicBox) SetGradientAndMountColorListFilter(f func(iotmaker_platform_IDraw.ICanvasGradient)) {
	el.prepareGradientFilter = f
}

func (el *BasicBox) SetShadowFilter(f func(iotmaker_platform_IDraw.ICanvasShadow)) {
	el.prepareShadowFilter = f
}

func (el *BasicBox) calculateCoordinates() {
	//  draw_1:
	//              border        border
	//             x1  x2         x3 x4
	//           l     a          b     c
	//        y1    +--|----------|--+    y1
	//  border      |                |      border
	//        y2 k ---              --- d y2
	//              |                |
	//              |                |
	//        y3 j ---              --- e y3
	//  border      |                |      border
	//        y4    +--|----------|--+    y4
	//           i     h          g     f
	//             x1  x2         x3 x4
	//              border        border

	// correct the line width size
	el.Dimensions.X -= el.Ink.LineWidth / 2
	el.Dimensions.Y -= el.Ink.LineWidth / 2
	el.Dimensions.Width -= el.Ink.LineWidth
	el.Dimensions.Height -= el.Ink.LineWidth

	// set coordinates from de box in draw_1
	el.x1 = el.Dimensions.X
	el.x2 = el.x1 + el.Dimensions.Border
	el.x3 = el.x2 + el.Dimensions.Width - 2*el.Dimensions.Border
	el.x4 = el.x3 + el.Dimensions.Border

	el.y1 = el.Dimensions.Y
	el.y2 = el.y1 + el.Dimensions.Border
	el.y3 = el.y2 + el.Dimensions.Height - 2*el.Dimensions.Border
	el.y4 = el.y3 + el.Dimensions.Border
}

func (el *BasicBox) prepareImageData() {
	el.imageData = el.Platform.GetImageData(el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)
}

func (el *BasicBox) GetAlphaChannel(x, y int) int16 {
	fmt.Printf("%v\n", el.imageData)

	if el.Dimensions.X > x {
		return -1
	}

	if el.Dimensions.Y > y {
		return -2
	}

	if x > el.Dimensions.X+el.Dimensions.Width || y > el.Dimensions.Y+el.Dimensions.Height {
		return -3
	}

	x -= el.Dimensions.X
	y -= el.Dimensions.Y

	return int16(el.imageData[x][y].A)
}

func (el *BasicBox) clearRectangle() {
	x := el.Dimensions.X - el.Ink.LineWidth/2
	y := el.Dimensions.Y - el.Ink.LineWidth/2
	width := el.Dimensions.Width + el.Ink.LineWidth
	height := el.Dimensions.Height + el.Ink.LineWidth

	el.Platform.ClearRect(x, y, width, height)
}

func (el *BasicBox) drawInvisible(platform iotmaker_platform_IDraw.IDraw) {
	platform.LineWidth(el.Ink.LineWidth)
	platform.BeginPath()
	platform.MoveTo(el.x2, el.y1)                                    // a
	platform.LineTo(el.x3, el.y1)                                    // a->b
	platform.ArcTo(el.x4, el.y1, el.x4, el.y2, el.Dimensions.Border) // c->d
	platform.LineTo(el.x4, el.y3)                                    // d->e
	platform.ArcTo(el.x4, el.y4, el.x3, el.y4, el.Dimensions.Border) // f->g
	platform.LineTo(el.x2, el.y4)                                    // g->h
	platform.ArcTo(el.x1, el.y4, el.x1, el.y3, el.Dimensions.Border) // i->j
	platform.LineTo(el.x1, el.y2)                                    // j->k
	platform.ArcTo(el.x1, el.y1, el.x2, el.y1, el.Dimensions.Border) // i->j
	platform.ClosePath(el.x2, el.y1)                                 // a
}

func (el *BasicBox) Create() {
	el.calculateCoordinates()

	//platform := new(iotmaker_platform_IDraw.IDraw)
	//platform := iotmaker_platform_webbrowser.Stage{}
	//platform := iotmaker_platform_webbrowser.Canvas{}

	//el.Tmp.SelfElement = platform.Element.SelfElement

	el.drawInvisible(el.ScratchPad)
	el.ScratchPad.Stroke()
	el.imageData = el.ScratchPad.GetImageData(el.Dimensions.X, el.Dimensions.Y, el.Dimensions.Width, el.Dimensions.Height)

	el.drawInvisible(el.Platform)
	el.PrepareGradientFilter()
	el.PrepareShadowFilter()

	el.Platform.Stroke()
}
