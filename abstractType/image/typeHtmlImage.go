package image

import iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"

type HtmlImage struct {
	Platform iotmaker_platform_IDraw.IHtml
	Img      interface{}
	Width    int
	Height   int
}

func (el *HtmlImage) Crete() {
	el.Platform.NewImage(el.Img, map[string]interface{}{"width": el.Width, "height": el.Height}, true)
}
