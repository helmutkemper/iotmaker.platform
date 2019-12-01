package basicBox

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
)

type BasicBox struct {
	Platform   iotmaker_platform_IDraw.IDraw
	Id         string
	Dimensions Dimensions
	Ink        Ink

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
	el.Dimensions.X += el.Ink.LineWidth / 2
	el.Dimensions.Y += el.Ink.LineWidth / 2
	el.Dimensions.Width -= el.Ink.LineWidth / 2
	el.Dimensions.Height -= el.Ink.LineWidth / 2

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

func (el *BasicBox) drawInvisible() {
	el.Platform.LineWidth(el.Ink.LineWidth)
	el.Platform.BeginPath()
	el.Platform.MoveTo(el.x2, el.y1)                                    // a
	el.Platform.LineTo(el.x3, el.y1)                                    // a->b
	el.Platform.ArcTo(el.x4, el.y1, el.x4, el.y2, el.Dimensions.Border) // c->d
	el.Platform.LineTo(el.x4, el.y3)                                    // d->e
	el.Platform.ArcTo(el.x4, el.y4, el.x3, el.y4, el.Dimensions.Border) // f->g
	el.Platform.LineTo(el.x2, el.y4)                                    // g->h
	el.Platform.ArcTo(el.x1, el.y4, el.x1, el.y3, el.Dimensions.Border) // i->j
	el.Platform.LineTo(el.x1, el.y2)                                    // j->k
	el.Platform.ArcTo(el.x1, el.y1, el.x2, el.y1, el.Dimensions.Border) // i->j
	el.Platform.ClosePath(el.x2, el.y1)                                 // a
}

func (el *BasicBox) Create() {
	el.calculateCoordinates()
	el.drawInvisible()

	el.PrepareGradientFilter()
	el.PrepareShadowFilter()

	el.Platform.Stroke()
}
