package mouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
)

type EventCoordinate struct {
	X             int
	Y             int
	ToBeProcessed bool
}

var Move = EventCoordinate{
	X:             -1000000,
	Y:             -1000000,
	ToBeProcessed: false,
}

var Click = EventCoordinate{
	X:             -1000000,
	Y:             -1000000,
	ToBeProcessed: false,
}

var DoubleClick = EventCoordinate{
	X:             -1000000,
	Y:             -1000000,
	ToBeProcessed: false,
}

var Press = EventCoordinate{
	X:             -1000000,
	Y:             -1000000,
	ToBeProcessed: false,
}

var Release = EventCoordinate{
	X:             -1000000,
	Y:             -1000000,
	ToBeProcessed: false,
}

func init() {
	go func() {
		for {
			select {
			case coordinate := <-mouse.BrowserMouseToPlatformMouseCoordinate:
				Move.X = coordinate.X
				Move.Y = coordinate.Y
				Move.ToBeProcessed = true

			case coordinate := <-mouse.BrowserMouseClickToPlatformMouseClickEvent:
				Click.X = coordinate.X
				Click.Y = coordinate.Y
				Click.ToBeProcessed = true

			case coordinate := <-mouse.BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent:
				DoubleClick.X = coordinate.X
				DoubleClick.Y = coordinate.Y
				DoubleClick.ToBeProcessed = true

			case coordinate := <-mouse.BrowserMouseDownToPlatformMouseDownEvent:
				Press.X = coordinate.X
				Press.Y = coordinate.Y
				Press.ToBeProcessed = true

			case coordinate := <-mouse.BrowserMouseUpToPlatformMouseUpEvent:
				Release.X = coordinate.X
				Release.Y = coordinate.Y
				Release.ToBeProcessed = true
			}
		}
	}()
}
