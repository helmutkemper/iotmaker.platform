package dimensions

type Overflow int

type Dimensions struct {
	X                     int
	Y                     int
	Width                 int
	Height                int
	Border                int
	originalX             int
	originalY             int
	originalWidth         int
	originalHeight        int
	originalBorder        int
	lastInteractionX      int
	lastInteractionY      int
	lastInteractionWidth  int
	lastInteractionHeight int
	lastInteractionBorder int
	hasModified           bool

	Parent   *Dimensions
	Overflow Overflow
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

func (el *Dimensions) Set(x, y, width, height, border int) {

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
