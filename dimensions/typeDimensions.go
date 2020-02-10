package dimensions

import (
	"errors"
	"fmt"
	"reflect"
)

type Corner int

const (
	KContainerBLinkOnTopToContainerALinkOnTop Corner = iota
	KContainerBLinkOnTopToContainerALinkOnBottom
	KContainerBLinkOnBottomToContainerALinkOnBottom
	KContainerBLinkOnLeftToContainerALinkOnLeft
	KContainerBLinkOnRightToContainerALinkOnRight
)

type LinkCollection struct {
	Top    *Link
	Left   *Link
	Right  *Link
	Bottom *Link
}

type Link struct {
	Container      *Dimensions
	CornerToFather Corner
}

// en: It calculates the vertical position of the container when there is only the
// link on the bottom are connected to
//
// pt_br: Calcula o posicionamento vertical do container quando só existe o link
// inferior conectado
//
//                A
//     +----------O---------+
//     |                    |
//     |   +------O-----+   |
//   D O   O            O   O B
//     |   +------O-----+   |
//     |          |         |
//     +----------O---------+
//                C
func (el *Dimensions) LinkIsPresentAtTheBottom(containerA, containerB *Dimensions) {
	AY := containerA.Y
	AHeight := containerA.Height
	BHeight := containerB.Height
	BBottom := containerB.SpaceBottom
	containerB.Y = AY + AHeight - BHeight + BBottom
}

// en: It calculates the vertical position of the container when there is only the
// link on the top are connected to
//
// pt_br: Calcula o posicionamento vertical do container quando só existe o link
// superior conectado
//
//                A
//     +----------O---------+
//     |          |         |
//     |   +------O-----+   |
//   D O   O            O   O B
//     |   +------O-----+   |
//     |                    |
//     +----------O---------+
//                C
func (el *Dimensions) LinkIsPresentAtTheTop(containerA, containerB *Dimensions) {
	AY := containerA.Y
	BTop := containerB.SpaceTop
	containerB.Y = AY + BTop
}

// en: Calculates the vertical position of the container when the links to the top
// and bottom are connected
//
// pt_br: Calcula o posicionamento vertical do container quando existe os links
// superior e inferior conectados
//
//                A
//     +----------O---------+
//     |          |         |
//     |   +------O-----+   |
//   D O   O            O   O B
//     |   +------O-----+   |
//     |          |         |
//     +----------O---------+
//                C
func (el *Dimensions) LinkIsPresentAtTheTopAndBottom(containerA, containerB *Dimensions) {
	AY := containerA.Y
	AHeight := containerA.Height
	CHeight := containerB.Height
	BTop := containerB.SpaceTop
	BY := AY + AHeight/2 - CHeight/2

	if BY < BTop {
		BY = BTop
	}

	containerB.Y = BY
}

// en: It calculates the horizontal position of the container when there is only the
// link on the right are connected to
//
// pt_br: Calcula o posicionamento horizontal do container quando só existe o link
// a direita conectado
//
//                A
//     +----------O---------+
//     |                    |
//     |   +------O-----+   |
//   D O   O            O---O B
//     |   +------O-----+   |
//     |                    |
//     +----------O---------+
//                C
func (el *Dimensions) LinkIsPresentAtTheRight(containerA, containerB *Dimensions) {
	AX := containerA.X
	AWidth := containerA.Width
	CWidth := containerB.Width
	containerB.X = AX + AWidth - CWidth
}

// en: It calculates the horizontal position of the container when there is only the
// link on the left are connected to
//
// pt_br: Calcula o posicionamento horizontal do container quando só existe o link
// a esquerda conectado
//
//                A
//     +----------O---------+
//     |                    |
//     |   +------O-----+   |
//   D O---O            O   O B
//     |   +------O-----+   |
//     |                    |
//     +----------O---------+
//                C
func (el *Dimensions) LinkIsPresentAtTheLeft(containerA, containerB *Dimensions) {
	AX := containerA.X
	containerB.X = AX
}

// en: Calculates the horizontal position of the container when the links to the
// right and the left are connected
//
// pt_br: Calcula o posicionamento horizontal do container quando existe os links
// direito e esquerdos conectados
//
//                A
//     +----------O---------+
//     |                    |
//     |   +------O-----+   |
//   D O---O            O---O B
//     |   +------O-----+   |
//     |                    |
//     +----------O---------+
//                C
func (el *Dimensions) LinkIsPresentAtTheRightAndLeft(containerA, containerB *Dimensions) {
	AX := containerA.X
	AWidth := containerA.Width
	CWidth := containerB.Width
	containerB.X = AX + AWidth/2 - CWidth/2
}

func (el *Link) LinkAssembly(father LinkCollection, list []LinkCollection) error {

	return nil
}

func (el *Link) SimpleLinkAssembly(father *Dimensions, container LinkCollection) error {

	//
	// +-------------O------------+
	// |             |topIsSet    |
	// | +-----------O----------+ |
	// O-O leftIsSet            O-O rightIsSet
	// | +-----------O----------+ |
	// |             |bottomIsSet |
	// +-------------O------------+
	topIsSet := !reflect.DeepEqual(*container.Top, Link{})
	leftIsSet := !reflect.DeepEqual(*container.Left, Link{})
	rightIsSet := !reflect.DeepEqual(*container.Right, Link{})
	bottomIsSet := !reflect.DeepEqual(*container.Bottom, Link{})

	if topIsSet == false && bottomIsSet == false {
		return errors.New("this container can't be assembled. links a and c are not connected")
	} else if topIsSet == false && bottomIsSet == true {
		father.LinkIsPresentAtTheBottom(father, container.Bottom.Container)
	} else if topIsSet == true && bottomIsSet == false {
		father.LinkIsPresentAtTheTop(father, container.Top.Container)
	} else if topIsSet == true && bottomIsSet == true {
		father.LinkIsPresentAtTheTopAndBottom(father, container.Bottom.Container)
	}

	if leftIsSet == false && rightIsSet == false {
		return errors.New("this container can't be assembled. links b and d are not connected")
	} else if leftIsSet == false && rightIsSet == true {
		father.LinkIsPresentAtTheRight(father, container.Right.Container)
	} else if leftIsSet == true && rightIsSet == false {
		father.LinkIsPresentAtTheLeft(father, container.Left.Container)
	} else if leftIsSet == true && rightIsSet == true {
		father.LinkIsPresentAtTheRightAndLeft(father, container.Right.Container)
	}

	return nil
}

//                Top
//      +----------O---------+
//      |                    |
//      |   +------O-----+   |
// Left O   O            O   O Right
//      |   +------O-----+   |
//      |                    |
//      +----------O---------+
//               Bottom
//
//                      A
//             +--------O-------+
//             |                |
//           D O                O B
//             |                |
//             +--------O-------+
//                      C
//
//  Molde:teste
//  +-father------------O-------------------+
//  |                   |                   |
//  |     +-containerA--O-------------+     |
//  |     |                           |     |
//  O-----O                           O-----O
//  |     |                           |     |
//  |     +-------------O-------------+     |
//  |                   |                   |
//  |     +-containerB--O-------------+     |
//  |     |                           |     |
//  O-----O                           O-----O
//  |     |                           |     |
//  |     +-------------O-------------+     |
//  |                   |                   |
//  +-------------------O-------------------+
//
//  Ponto de vista horizontal
//  Option: A - cada container é centralizado em relação ao pai
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O-----O                           O-----O
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  |     +-containerB----------------+     |
//  |     |                           |     |
//  O-----O                           O-----O
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: B - cada container é alinhado a direita em relação ao pai
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |     |                           O-----O
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  |     +-containerB----------------+     |
//  |     |                           |     |
//  |     |                           O-----O
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: C - cada container é alinhado a esquerda em relação ao pai
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O-----O                           |     |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  |     +-containerB----------------+     |
//  |     |                           |     |
//  O-----O                           |     |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: D - os containers a e b têm o seu tamanho ajustados pelo maior
//              o container a é centralizado pelo pai
//              o container b y é ajustado por a y
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O--+--O a link to father          O--+--O
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  |  |                                 |  |
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O b link to a               O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: E - o container a é ajustado a direita pelo pai
//              o container b y+w é ajustado por a y+w
//              w não é alterado nem em a nem em b
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |     | a link to father          O--+--O
//  |     |                           |  |  |
//  |     +---------------------------+  |  |
//  |                                    |  |
//  |     +-containerB----------------+  |  |
//  |     |                           |  |  |
//  |     | b link to a               O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: F - o container a é ajustado a direita pelo pai
//              o container b y+w é ajustado por a y+w
//              w é alterado pelo maior
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O a link to father          O--+--O
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  |  |                                 |  |
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O b link to a               O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: G
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O--+--O a link to father          |     |
//  |  |  |                           |     |
//  |  |  +---------------------------+     |
//  |  |                                    |
//  |  |  +-containerB----------------+     |
//  |  |  |                           |     |
//  |  +--O b link to a               |     |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: H
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O--+--O a link to father          O--+  |
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  |  |                                 |  |
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O b link to a               O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: I
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O                           O-----O
//  |  |  |                           |     |
//  |  |  +---------------------------+     |
//  |  |                                    |
//  |  |  +-containerB----------------+     |
//  |  |  |                           |     |
//  |  +--O                           O-----O
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: J
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O-----O                           O--+  |
//  |     |                           |  |  |
//  |     +---------------------------+  |  |
//  |                                    |  |
//  |     +-containerB----------------+  |  |
//  |     |                           |  |  |
//  O-----O                           O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Errado:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O                           O--+  |
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  X  |                                 |  X
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O                           O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Errado:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O                           O-----O
//  |  |  |                           |     |
//  |  |  +---------------------------+     |
//  X  |                                    |
//  |  |  +-containerB----------------+     |
//  |  |  |                           |     |
//  |  +--O                           X     |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  +-father--------------------------------O---------------------------------------+
//  |                                       |                                       |
//  |                   +-------------------+-------------------+                   |
//  |                   |                                       |                   |
//  |     +-containerA--O-------------+           +-containerB--O-------------+     |
//  |     |                           |           |                           |     |
//  |  +--O                           O-----------O                           O--+  |
//  |  |  |                           |           |                           |  |  |
//  |  |  +-------------O-------------+           +-------------O-------------+  |  |
//  O--+                |                                       |                +--O
//  |  |  +-containerC--O-------------+           +-containerD--O-------------+  |  |
//  |  |  |                           |           |                           |  |  |
//  |  +--O                           O-----------O                           O--+  |
//  |     |                           |           |                           |     |
//  |     +-------------O-------------+           +-------------O-------------+     |
//  |                   |                                       |                   |
//  |                   +-------------------+-------------------+                   |
//  |                                       |                                       |
//  +---------------------------------------O---------------------------------------+
//
type Dimensions struct {
	Level int

	X      int
	Y      int
	Width  int
	Height int
	Border int

	SpaceLeft   int
	SpaceRight  int
	SpaceTop    int
	SpaceBottom int
}

func NewContainer(width, height int) *Dimensions {
	return &Dimensions{
		X:      0,
		Y:      0,
		Width:  width,
		Height: height,

		SpaceLeft:   0,
		SpaceRight:  0,
		SpaceTop:    0,
		SpaceBottom: 0,
	}
}

func NewContainerWidthXY(x, y, width, height int) *Dimensions {
	return &Dimensions{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,

		SpaceLeft:   0,
		SpaceRight:  0,
		SpaceTop:    0,
		SpaceBottom: 0,
	}
}

func NewContainerWidthXYAndSpaces(x, y, width, height, left, right, top, bottom int) *Dimensions {
	return &Dimensions{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,

		SpaceLeft:   left,
		SpaceRight:  right,
		SpaceTop:    top,
		SpaceBottom: bottom,
	}
}

func NewContainerWithSpace(width, height, left, right, top, bottom int) *Dimensions {
	return &Dimensions{
		X:      0,
		Y:      0,
		Width:  width,
		Height: height,

		SpaceLeft:   left,
		SpaceRight:  right,
		SpaceTop:    top,
		SpaceBottom: bottom,
	}
}

// Warning: Father aways must be containerA
// Warning: The link should always be in the container children to be the container Father
func NewLink(container *Dimensions, cornerBToA Corner) *Link {
	return &Link{
		Container:      container,
		CornerToFather: cornerBToA,
	}
}

func Test() []Dimensions {
	var err error

	var linkFatherTopToContainerATop = &Link{}
	var linkFatherLeftToContainerALeft = &Link{}
	var linkFatherRightToContainerARight = &Link{}
	var linkFatherBottomToContainerABottom = &Link{}

	containerA := NewContainerWidthXY(300, 100, 300, 150)
	containerB := NewContainerWidthXY(300, 150, 600, 300)

	linkFatherTopToContainerATop = NewLink(containerB, KContainerBLinkOnTopToContainerALinkOnTop)
	linkFatherLeftToContainerALeft = NewLink(containerB, KContainerBLinkOnLeftToContainerALinkOnLeft)
	linkFatherRightToContainerARight = NewLink(containerB, KContainerBLinkOnRightToContainerALinkOnRight)
	linkFatherBottomToContainerABottom = NewLink(containerB, KContainerBLinkOnBottomToContainerALinkOnBottom)

	collection := LinkCollection{
		Top:    linkFatherTopToContainerATop,
		Left:   linkFatherLeftToContainerALeft,
		Right:  linkFatherRightToContainerARight,
		Bottom: linkFatherBottomToContainerABottom,
	}

	err = linkFatherTopToContainerATop.SimpleLinkAssembly(
		containerA,
		collection,
	)
	if err != nil {
		fmt.Printf("SimpleLinkAssembly().error: %v\n", err.Error())
	}

	return []Dimensions{*containerA, *containerB}
}
