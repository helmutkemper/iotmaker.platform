package dimensions

import (
	"errors"
	"fmt"
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
		Width:  width,
		Height: height,
	}
}

func NewContainerWithSpace(width, height, left, right, top, bottom int) *Dimensions {
	return &Dimensions{
		Width:  width,
		Height: height,

		SpaceLeft:   left,
		SpaceRight:  right,
		SpaceTop:    top,
		SpaceBottom: bottom,
	}
}

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

func NewAssembly(linkA, linkB, linkC, linkD *Link) {

}

func test() {
	var err error

	var linkFatherAtoContainerA *Link
	var linkFatherBtoContainerB *Link
	var linkFatherCtoContainerC *Link
	var linkFatherDtoContainerD *Link

	father := NewContainer(600, 300)
	containerA := NewContainer(300, 100)
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

	NewAssembly(
		linkFatherAtoContainerA,
		linkFatherBtoContainerB,
		linkFatherCtoContainerC,
		linkFatherDtoContainerD,
	)
}
