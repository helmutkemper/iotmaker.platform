package dimensions

import (
	"fmt"
)

func ExampleFilterContainerFather() {

	father := NewContainer(300, 600)
	containerA := NewContainerWidthXY(10, 10, 100, 100)
	containerB := NewContainerWidthXY(10, 120, 100, 100)
	containerC := NewContainerWidthXY(10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		containerA,
		containerC,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	retDimensions := filter.FilterContainerFather(listLink)

	if len(retDimensions) == 1 && retDimensions[0] == father {
		fmt.Println("passou")
	}

	// output:
	// passou
}

//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           O------O
//  | +----O                           O----+ |
//  | | +--O                           O--+ | |
//  | | |  +---------------------------+  | | |
//  | | |                                 | | |
//  | | |  +-containerB----------------+  | | |
//  | | |  |                           |  | | |
//  | | +--O                           O--+ | |
//  | |    |                           |    | |
//  | |    +---------------------------+    | |
//  | |                                     | |
//  | |    +-containerC----------------+    | |
//  | |    |                           |    | |
//  | +----O                           O----+ |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           O------O
//  |      |                           O      |
//  |   +--O                           O--+   |
//  |   |  +---------------------------+  |   |
//  |   |                                 |   |
//  |   |  +-containerB----------------+  |   |
//  |   +--O                           O--+   |
//  |      |                           |      |
//  |   +--O                           O--+   |
//  |   |  +---------------------------+  |   |
//  |   |                                 |   |
//  |   |  +-containerC----------------+  |   |
//  |   |  |                           |  |   |
//  |   +--O                           O--+   |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           O------O
//  |      |                           |      |
//  |   +--O                           |      |
//  |   |  +---------------------------+      |
//  |   |                                     |
//  |   |  +-containerB----------------+      |
//  |   +--O                           |      |
//  |      |                           |      |
//  |   +--O                           |      |
//  |   |  +---------------------------+      |
//  |   |                                     |
//  |   |  +-containerC----------------+      |
//  |   |  |                           |      |
//  |   +--O                           |      |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           O------O
//  |      |                           O      |
//  |      |                           O--+   |
//  |      +---------------------------+  |   |
//  |                                     |   |
//  |      +-containerB----------------+  |   |
//  |      |                           O--+   |
//  |      |                           |      |
//  |      |                           O--+   |
//  |      +---------------------------+  |   |
//  |                                     |   |
//  |      +-containerC----------------+  |   |
//  |      |                           |  |   |
//  |      |                           O--+   |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  |      |                           O------O
//  | +----O                           O----+ |
//  | | +--O                           O--+ | |
//  | | |  +---------------------------+  | | |
//  | | |                                 | | |
//  | | |  +-containerB----------------+  | | |
//  | | |  |                           |  | | |
//  | | +--O                           O--+ | |
//  | |    |                           |    | |
//  | |    +---------------------------+    | |
//  | |                                     | |
//  | |    +-containerC----------------+    | |
//  | |    |                           |    | |
//  | +----O                           O----+ |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  |      |                           O------O
//  |      |                           O      |
//  |   +--O                           O--+   |
//  |   |  +---------------------------+  |   |
//  |   |                                 |   |
//  |   |  +-containerB----------------+  |   |
//  |   +--O                           O--+   |
//  |      |                           |      |
//  |   +--O                           O--+   |
//  |   |  +---------------------------+  |   |
//  |   |                                 |   |
//  |   |  +-containerC----------------+  |   |
//  |   |  |                           |  |   |
//  |   +--O                           O--+   |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  |      |                           O------O
//  |      |                           |      |
//  |   +--O                           |      |
//  |   |  +---------------------------+      |
//  |   |                                     |
//  |   |  +-containerB----------------+      |
//  |   +--O                           |      |
//  |      |                           |      |
//  |   +--O                           |      |
//  |   |  +---------------------------+      |
//  |   |                                     |
//  |   |  +-containerC----------------+      |
//  |   |  |                           |      |
//  |   +--O                           |      |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  |      |                           O------O
//  |      |                           O      |
//  |      |                           O--+   |
//  |      +---------------------------+  |   |
//  |                                     |   |
//  |      +-containerB----------------+  |   |
//  |      |                           O--+   |
//  |      |                           |      |
//  |      |                           O--+   |
//  |      +---------------------------+  |   |
//  |                                     |   |
//  |      +-containerC----------------+  |   |
//  |      |                           |  |   |
//  |      |                           O--+   |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           |      |
//  | +----O                           O----+ |
//  | | +--O                           O--+ | |
//  | | |  +---------------------------+  | | |
//  | | |                                 | | |
//  | | |  +-containerB----------------+  | | |
//  | | |  |                           |  | | |
//  | | +--O                           O--+ | |
//  | |    |                           |    | |
//  | |    +---------------------------+    | |
//  | |                                     | |
//  | |    +-containerC----------------+    | |
//  | |    |                           |    | |
//  | +----O                           O----+ |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           |      |
//  |      |                           O      |
//  |   +--O                           O--+   |
//  |   |  +---------------------------+  |   |
//  |   |                                 |   |
//  |   |  +-containerB----------------+  |   |
//  |   +--O                           O--+   |
//  |      |                           |      |
//  |   +--O                           O--+   |
//  |   |  +---------------------------+  |   |
//  |   |                                 |   |
//  |   |  +-containerC----------------+  |   |
//  |   |  |                           |  |   |
//  |   +--O                           O--+   |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           |      |
//  |      |                           |      |
//  |   +--O                           |      |
//  |   |  +---------------------------+      |
//  |   |                                     |
//  |   |  +-containerB----------------+      |
//  |   +--O                           |      |
//  |      |                           |      |
//  |   +--O                           |      |
//  |   |  +---------------------------+      |
//  |   |                                     |
//  |   |  +-containerC----------------+      |
//  |   |  |                           |      |
//  |   +--O                           |      |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
//
//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           |      |
//  |      |                           O      |
//  |      |                           O--+   |
//  |      +---------------------------+  |   |
//  |                                     |   |
//  |      +-containerB----------------+  |   |
//  |      |                           O--+   |
//  |      |                           |      |
//  |      |                           O--+   |
//  |      +---------------------------+  |   |
//  |                                     |   |
//  |      +-containerC----------------+  |   |
//  |      |                           |  |   |
//  |      |                           O--+   |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+

//  +-father----------------------------------+
//  |                                         |
//  |      +-containerA----------------+      |
//  O------O                           O------O
//  | +----O                           O----+ |
//  | | +--O                           O--+ | |
//  | | |  +---------------------------+  | | |
//  | | |                                 | | |
//  | | |  +-containerB----------------+  | | |
//  | | |  |                           |  | | |
//  | | +--O                           O--+ | |
//  | |    |                           |    | |
//  | |    +---------------------------+    | |
//  | |                                     | |
//  | |    +-containerC----------------+    | |
//  | |    |                           |    | |
//  | +----O                           O----+ |
//  |      |                           |      |
//  |      +---------------------------+      |
//  |                                         |
//  +-----------------------------------------+
func ExampleLinkAssemblyHorizontalFilterContainerIsCentralizedInRelationToTheFather() {

	father := NewContainer(300, 600)
	containerA := NewContainerWidthXY(10, 10, 100, 100)
	containerB := NewContainerWidthXY(10, 120, 100, 100)
	containerC := NewContainerWidthXY(10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		containerA,
		containerC,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalFilterContainerIsCentralizedInRelationToTheFather(listLink)

	if len(ret) == 1 && ret[0] == linkAToFather {
		fmt.Println("passou")
	}

	// output:
	// passou
}

func ExampleLinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftAndTheRightInRelationToAnotherContainer() {

	father := NewContainer(300, 600)
	containerA := NewContainerWidthXY(10, 10, 100, 100)
	containerB := NewContainerWidthXY(10, 120, 100, 100)
	containerC := NewContainerWidthXY(10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		containerA,
		containerC,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	//listDimensions := []*Dimensions{father, containerA, containerB, containerC}
	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftAndTheRightInRelationToAnotherContainer(containerA, listLink)
	if len(ret) == 3 {

		afPass := ret[0] == linkAToFather
		abPass := ret[1] == linkAToB
		acPass := ret[2] == linkAToC
		pass := afPass && abPass && acPass
		if pass == true {
			fmt.Println("passou")
		}
	}

	ret = filter.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftAndTheRightInRelationToAnotherContainer(containerB, listLink)
	if len(ret) == 1 {

		abPass := ret[0] == linkAToB
		if abPass == true {
			fmt.Println("passou")
		}
	}

	// output:
	// passou
	// passou
}

func ExampleLinkAssemblyCheckIfFatherExistsAndHasOnlyOne() {

	father := NewContainer(300, 600)
	containerA := NewContainerWidthXY(10, 10, 100, 100)
	containerB := NewContainerWidthXY(10, 120, 100, 100)
	containerC := NewContainerWidthXY(10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		containerA,
		containerC,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	ret := filter.LinkAssemblyCheckIfFatherExistsAndHasOnlyOne(listLink)

	if ret == true {
		fmt.Println("passou")
	}

	// output:
	// passou
}

func ExampleLinkAssemblyHorizontalCheckEachContainerIsCentralizedInRelationToTheFather_1() {

	father := NewContainer(300, 600)
	containerA := NewContainerWidthXY(10, 10, 100, 100)
	containerB := NewContainerWidthXY(10, 120, 100, 100)
	containerC := NewContainerWidthXY(10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		containerA,
		containerC,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalCheckEachContainerIsCentralizedInRelationToTheFather(listLink)

	if ret == false {
		fmt.Println("passou")
	}

	// output:
	// passou
}

func ExampleLinkAssemblyHorizontalCheckEachContainerIsCentralizedInRelationToTheFather_2() {

	father := NewContainer(300, 600)
	containerA := NewContainerWidthXY(10, 10, 100, 100)
	containerB := NewContainerWidthXY(10, 120, 100, 100)
	containerC := NewContainerWidthXY(10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkBToFather := NewLinkWithFather(
		father,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkCToFather := NewLinkWithFather(
		father,
		containerC,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkBToFather, linkCToFather}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalCheckEachContainerIsCentralizedInRelationToTheFather(listLink)

	if ret == true {
		fmt.Println("passou")
	}

	// output:
	// passou
}
