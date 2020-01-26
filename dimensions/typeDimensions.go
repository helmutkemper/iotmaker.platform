package dimensions

type Dimensions struct {
	X                     float64
	Y                     float64
	Width                 float64
	Height                float64
	Border                float64
	originalX             float64
	originalY             float64
	originalWidth         float64
	originalHeight        float64
	originalBorder        float64
	lastInteractionX      float64
	lastInteractionY      float64
	lastInteractionWidth  float64
	lastInteractionHeight float64
	lastInteractionBorder float64
	hasModified           bool
}

func (el Dimensions) GetModified() bool {
	return el.hasModified
}

func (el *Dimensions) RevertToLastValue() {
	el.X = el.lastInteractionX
	el.Y = el.lastInteractionY
	el.Width = el.lastInteractionWidth
	el.Height = el.lastInteractionHeight
	el.Border = el.lastInteractionBorder
	el.hasModified = true
}

func (el *Dimensions) RevertToOriginalValue() {
	el.X = el.originalX
	el.Y = el.originalY
	el.Width = el.originalWidth
	el.Height = el.originalHeight
	el.Border = el.originalBorder
	el.hasModified = true
}

func (el *Dimensions) Set(x, y, width, height, border float64) {

	if el.originalX == 0 && el.originalY == 0 && el.originalWidth == 0 && el.originalHeight == 0 && el.originalBorder == 0 {
		el.originalX = x
		el.originalY = y
		el.originalWidth = width
		el.originalHeight = height
		el.originalBorder = border
	}

	if el.lastInteractionX != x && el.lastInteractionY != y && el.lastInteractionWidth != width && el.lastInteractionHeight != height && el.lastInteractionBorder != border {
		el.lastInteractionX = x
		el.lastInteractionY = y
		el.lastInteractionWidth = width
		el.lastInteractionHeight = height
		el.lastInteractionBorder = border
		el.hasModified = true
	} else {
		el.hasModified = false
	}

	el.X = x
	el.Y = y
	el.Width = width
	el.Height = height
	el.Border = border
}
