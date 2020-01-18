package mouse

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
)

type EventCoordinate struct {
	X int
	Y int
}

var Move = EventCoordinate{X: -1000000, Y: -1000000}
var Click = EventCoordinate{X: -1000000, Y: -1000000}
var DoubleClick = EventCoordinate{X: -1000000, Y: -1000000}
var Press = EventCoordinate{X: -1000000, Y: -1000000}
var Release = EventCoordinate{X: -1000000, Y: -1000000}

func init() {
	go func() {
		for {
			select {
			case coordinate := <-mouse.BrowserMouseToPlatformMouseCoordinate:
				Move.X = coordinate.X
				Move.Y = coordinate.Y

			case coordinate := <-mouse.BrowserMouseClickToPlatformMouseClickEvent:
				Click.X = coordinate.X
				Click.Y = coordinate.Y

			case coordinate := <-mouse.BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent:
				DoubleClick.X = coordinate.X
				DoubleClick.Y = coordinate.Y

			case coordinate := <-mouse.BrowserMouseDownToPlatformMouseDownEvent:
				Press.X = coordinate.X
				Press.Y = coordinate.Y

			case coordinate := <-mouse.BrowserMouseUpToPlatformMouseUpEvent:
				Release.X = coordinate.X
				Release.Y = coordinate.Y
			}
		}
	}()
}
