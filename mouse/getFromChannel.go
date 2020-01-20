package mouse

import (
	"errors"
	"fmt"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
)

type SendToChannel struct {
	List []chan Coordinate
}

func (el *SendToChannel) Init() {
	el.List = make([]chan Coordinate, 0)
}

func (el *SendToChannel) Add(value chan Coordinate) int {
	value = make(chan Coordinate)
	el.List = append(el.List, value)

	return len(el.List) - 1
}

func (el *SendToChannel) Remove(key int) error {

	if key < 0 {
		return errors.New("the key must be greater than or equal to zero")
	}

	if len(el.List) >= key {
		return errors.New("the key was not found")
	}

	el.List = append(el.List[:key], el.List[key+1:]...)

	return nil
}

type Coordinate struct {
	X    int
	Y    int
	Test bool
}

func (el *Coordinate) SetEventCoordinate(x, y int) {
	el.X = x
	el.Y = y
	el.Test = true
}

func (el *Coordinate) GetEventCoordinate() (int, int) {
	el.Test = false
	return el.X, el.Y
}

var Move SendToChannel
var Click SendToChannel
var DoubleClick SendToChannel
var Up SendToChannel
var Down SendToChannel

func init() {

	Move.Init()
	Click.Init()
	DoubleClick.Init()
	Up.Init()
	Down.Init()

	go func() {
		for {
			select {
			case coordinate := <-mouse.BrowserMouseToPlatformMouseCoordinate:
				for k := range Move.List {
					close(Move.List[k])
					Move.List[k] = make(chan Coordinate)
					Move.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y}
				}

			case coordinate := <-mouse.BrowserMouseClickToPlatformMouseClickEvent:
				for k := range Click.List {
					Click.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y}
				}

				fmt.Printf("mouse click\n")

			case coordinate := <-mouse.BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent:
				for k := range DoubleClick.List {
					DoubleClick.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y}
				}

				fmt.Printf("mouse double click\n")

			case coordinate := <-mouse.BrowserMouseDownToPlatformMouseDownEvent:
				for k := range Down.List {
					Down.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y}
				}

				fmt.Printf("mouse down\n")

			case coordinate := <-mouse.BrowserMouseUpToPlatformMouseUpEvent:
				for k := range Up.List {
					Up.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y}
				}

				fmt.Printf("mouse up\n")
			}
		}
	}()
}
