package mouse

import (
	"errors"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
)

type SendToChannel struct {
	List []chan Coordinate
}

func (el *SendToChannel) Init() {
	el.List = make([]chan Coordinate, 0)
}

func (el *SendToChannel) Add(value chan Coordinate) int {

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
	X       int
	Y       int
	HasInit bool
}

func (el *Coordinate) SetEventCoordinate(x, y int) {
	el.X = x
	el.Y = y
	el.HasInit = true
}

func (el *Coordinate) GetEventCoordinate() (int, int) {
	return el.X, el.Y
}

var Move SendToChannel
var Click SendToChannel
var DoubleClick SendToChannel
var Up SendToChannel
var Down SendToChannel

// en: A channel can be initialized in two ways:
// make (chan type)
// make (chan type, length)
// When length is used to start a channel, the len() and cap() functions can be
// used to determine whether or not a channel can receive more data.
// As some events happen much faster than engine events, there is the risk that the
// channel is full, causing a fatal go crash, so the len() function is used to clear
// the channel before new data can be written.
// if len( channel ) == 1 {
//   <- channel // clear channel
// }
//
// pt_br: Um canal pode ser inicializado de duas formas:
// make( chan type)
// make( chan type, length )
// Quando length é usado para iniciar um canal, as funções len() e cap() podem ser
// usadas para determinar se um canal pode ou não receber mais dados.
// Como alguns eventos acontecem muito mais rápido do que os eventos da engine, há o
// risco do canal está cheio, causando um travamento fatal do go, por isto, a função
// len() é usada para limpar o canal antes que um novo dado possa ser escrito.
// if len( channel ) == 1 {
//   <- channel // clear channel
// }
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

					if len(Move.List[k]) == 1 {
						<-Move.List[k]
					}

					Move.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y, HasInit: true}
				}

			case coordinate := <-mouse.BrowserMouseClickToPlatformMouseClickEvent:
				for k := range Click.List {

					if len(Click.List[k]) == 1 {
						<-Click.List[k]
					}

					Click.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y, HasInit: true}
				}

			case coordinate := <-mouse.BrowserMouseDoubleClickToPlatformMouseDoubleClickEvent:
				for k := range DoubleClick.List {

					if len(DoubleClick.List[k]) == 1 {
						<-DoubleClick.List[k]
					}

					DoubleClick.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y, HasInit: true}
				}

			case coordinate := <-mouse.BrowserMouseDownToPlatformMouseDownEvent:
				for k := range Down.List {

					if len(Down.List[k]) == 1 {
						<-Down.List[k]
					}

					Down.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y, HasInit: true}
				}

			case coordinate := <-mouse.BrowserMouseUpToPlatformMouseUpEvent:
				for k := range Up.List {

					if len(Up.List[k]) == 1 {
						<-Up.List[k]
					}

					Up.List[k] <- Coordinate{X: coordinate.X, Y: coordinate.Y, HasInit: true}
				}

			}
		}
	}()
}
