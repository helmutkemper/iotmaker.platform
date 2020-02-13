package dimensions

import (
	"fmt"
)

func ExampleFilterContainerFather() {

	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		"linkAToC",
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

	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		"linkAToC",
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

	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		"linkAToC",
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
	ret := filter.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheRightAndTheLeftInRelationToAnother(containerA, listLink)
	if len(ret) == 3 {

		afPass := ret[0] == linkAToFather
		abPass := ret[1] == linkAToB
		acPass := ret[2] == linkAToC
		pass := afPass && abPass && acPass
		if pass == true {
			fmt.Println("passou")
		}
	}

	ret = filter.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheRightAndTheLeftInRelationToAnother(containerB, listLink)
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

	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)
	linkAToC := NewLink(
		"linkAToC",
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

func ExampleLinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 10, 340, 100, 100)
	containerE := NewContainerWidthXY("containerE", 10, 450, 100, 100)
	containerF := NewContainerWidthXY("containerF", 10, 560, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerF,
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkBToC := NewLink(
		"linkBToC",
		containerB,
		containerC,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkCToD := NewLink(
		"linkCToD",
		containerC,
		containerD,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkBToE := NewLink(
		"linkBToE",
		containerB,
		containerE,
		KCornerTopContainerBLinkOnTopToContainerALinkOnBottom,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkBToC, linkCToD, linkBToE}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(containerA, listLink)

	fmt.Printf("%v\n", ret)

	// Output:
	//
}
