package dimensions

import (
	"errors"
	"fmt"
	"reflect"
)

type Corner int

const (
	KCornerA Corner = iota
	KCornerB
	KCornerC
	KCornerD
)

type Link struct {
	ContainerA  *Dimensions
	CornerFromA Corner
	ContainerB  *Dimensions
	CornerFromB Corner
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
func (el *Link) LinkIsPresentAtTheBottom(containerA, containerB *Dimensions) {
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
func (el *Link) LinkIsPresentAtTheTop(containerA, containerB *Dimensions) {
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
func (el *Link) LinkIsPresentAtTheTopAndBottom(containerA, containerB *Dimensions) {
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
func (el *Link) LinkIsPresentAtTheRight(containerA, containerB *Dimensions) {
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
func (el *Link) LinkIsPresentAtTheLeft(containerA, containerB *Dimensions) {
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
func (el *Link) LinkIsPresentAtTheRightAndLeft(containerA, containerB *Dimensions) {
	AX := containerA.X
	AWidth := containerA.Width
	CWidth := containerB.Width
	containerB.X = AX + AWidth/2 - CWidth/2
}

func (el *Link) SimpleLinkAssembly(linkA, linkB, linkC, linkD *Link) error {

	//                      A
	//             +--------O-------+
	//             |        |aPass  |
	//             | +------O-----+ |
	//           D O-O dPass      O-O B bPass
	//             | +------O-----+ |
	//             |        |cPass  |
	//             +--------O-------+
	//                      C
	aPass := !reflect.DeepEqual(*linkA, Link{})
	bPass := !reflect.DeepEqual(*linkB, Link{})
	cPass := !reflect.DeepEqual(*linkC, Link{})
	dPass := !reflect.DeepEqual(*linkD, Link{})

	if aPass == false && cPass == false {
		return errors.New("this container can't be assembled. links a and c are not connected")
	} else if aPass == false && cPass == true {
		linkA.LinkIsPresentAtTheBottom(linkC.ContainerA, linkC.ContainerB)
	} else if aPass == true && cPass == false {
		linkA.LinkIsPresentAtTheTop(linkA.ContainerA, linkA.ContainerB)
	} else if aPass == true && cPass == true {
		linkA.LinkIsPresentAtTheTopAndBottom(linkA.ContainerA, linkA.ContainerB)
	}

	if dPass == false && bPass == false {
		return errors.New("this container can't be assembled. links b and d are not connected")
	} else if dPass == false && bPass == true {
		linkA.LinkIsPresentAtTheRight(linkB.ContainerA, linkB.ContainerB)
	} else if dPass == true && bPass == false {
		linkA.LinkIsPresentAtTheLeft(linkD.ContainerA, linkD.ContainerB)
	} else if dPass == true && bPass == true {
		linkA.LinkIsPresentAtTheRightAndLeft(linkD.ContainerA, linkD.ContainerB)
	}

	return nil
}

//                A
//     +----------O---------+
//     |                    |
//     |   +------O-----+   |
//   D O   O            O   O B
//     |   +------O-----+   |
//     |                    |
//     +----------O---------+
//                C
//
//                      A
//             +--------O-------+
//             |                |
//           D O                O B
//             |                |
//             +--------O-------+
//                      C
//
//  +-father------------O-------------------+
//  |                   |                   |
//  |     +-containerA--O-------------+     |
//  |     |                           |     |
//  |  +--O                           O--+  |
//  |  |  |                           |  |  |
//  |  |  +-------------O-------------+  |  |
//  O--+                |                +--O
//  |  |  +-containerB--O-------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O                           O--+  |
//  |     |                           |     |
//  |     +-------------O-------------+     |
//  |                   |                   |
//  +-------------------O-------------------+
//
type Dimensions struct {
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
func NewLink(containerA, containerB *Dimensions, cornerFromA, cornerFromB Corner) (error, *Link) {
	switch cornerFromA {
	case KCornerA:
		fallthrough
	case KCornerC:

		if cornerFromB == KCornerB || cornerFromB == KCornerD {
			return errors.New("corners b and d must be linked to corners a and c"), &Link{}
		}

		return nil, &Link{
			ContainerA:  containerA,
			CornerFromA: cornerFromA,
			ContainerB:  containerB,
			CornerFromB: cornerFromB,
		}

	case KCornerB:
		fallthrough
	case KCornerD:

		if cornerFromB == KCornerA || cornerFromB == KCornerC {
			return errors.New("corners a and c must be linked to corners a and c"), &Link{}
		}

		return nil, &Link{
			ContainerA:  containerA,
			CornerFromA: cornerFromA,
			ContainerB:  containerB,
			CornerFromB: cornerFromB,
		}
	}

	return errors.New("corners a and c must be linked to corners a and c, and corners b and d must be linked to corners b and d"), &Link{}
}

func Test() []Dimensions {
	var err error

	var linkFatherAtoContainerA = &Link{}
	var linkFatherBtoContainerB = &Link{}
	var linkFatherCtoContainerC = &Link{}
	var linkFatherDtoContainerD = &Link{}

	father := NewContainer(600, 300)
	containerA := NewContainerWidthXY(10, 10, 300, 100)
	containerB := NewContainerWidthXY(10, 120, 300, 100)
	err, linkFatherAtoContainerA = NewLink(father, containerA, KCornerA, KCornerA)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	err, linkFatherBtoContainerB = NewLink(father, containerA, KCornerB, KCornerB)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	err, linkFatherCtoContainerC = NewLink(father, containerA, KCornerC, KCornerC)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	err, linkFatherDtoContainerD = NewLink(father, containerA, KCornerD, KCornerD)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err = linkFatherAtoContainerA.SimpleLinkAssembly(
		linkFatherAtoContainerA,
		linkFatherBtoContainerB,
		linkFatherCtoContainerC,
		linkFatherDtoContainerD,
	)
	if err != nil {
		fmt.Printf("SimpleLinkAssembly().error: %v\n", err.Error())
	}

	return []Dimensions{*father, *containerA, *containerB}
}
