package button

import (
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/text"
)

type Button struct {
	BasicBox draw.BasicBox
	Label    text.Text
}
