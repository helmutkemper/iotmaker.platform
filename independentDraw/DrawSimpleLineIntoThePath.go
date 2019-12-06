package independentDraw

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"

func DrawSimpleLineIntoThePath(platform iotmaker_platform_IDraw.IDraw, x0, y0, x1, y1 int) {
	if platform == nil {
		return
	}

	platform.BeginPath()
	platform.MoveTo(x0, y0)
	platform.LineTo(x1, y1)
	platform.ClosePath(x1, y1)
}
