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
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)
	linkAToC := NewLink(
		"linkAToC",
		containerA,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	retDimensions := filter.FilterContainerFather(listLink)

	if retDimensions != nil && retDimensions == father {
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
func ExampleLinkHorizontalFilterContainerIsCentralizedInRelationToTheFather() {

	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)
	linkAToC := NewLink(
		"linkAToC",
		containerA,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	ret := filter.LinkHorizontalFilterContainerIsCentralizedInRelationToTheFather(listLink)

	if len(ret) == 1 && ret[0] == linkAToFather {
		fmt.Println("passou")
	}

	// output:
	// passou
}

func ExampleLinkHorizontalFilterEachContainerIsAlignsFromTheLeftAndTheRightInRelationToAnotherContainer() {

	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)
	linkAToC := NewLink(
		"linkAToC",
		containerA,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	//listDimensions := []*Dimensions{father, containerA, containerB, containerC}
	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	ret := filter.LinkHorizontalFilterEachContainerIsAlignsFromTheRightAndTheLeftInRelationToAnother(containerA, listLink)
	if len(ret) == 3 {

		afPass := ret[0] == linkAToFather
		abPass := ret[1] == linkAToB
		acPass := ret[2] == linkAToC
		pass := afPass && abPass && acPass
		if pass == true {
			fmt.Println("passou")
		}
	}

	ret = filter.LinkHorizontalFilterEachContainerIsAlignsFromTheRightAndTheLeftInRelationToAnother(containerB, listLink)
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

func ExampleLinkCheckIfFatherExistsAndHasOnlyOne() {

	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 240, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)
	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)
	linkAToC := NewLink(
		"linkAToC",
		containerA,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	ret := filter.CheckIfFatherExistsAndHasOnlyOne(listLink)

	if ret == true {
		fmt.Println("passou")
	}

	// output:
	// passou
}

func ExampleLinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_1() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 10, 340, 100, 100)
	containerE := NewContainerWidthXY("containerE", 10, 450, 100, 100)
	containerF := NewContainerWidthXY("containerF", 10, 560, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkFToFather",
		containerF,
		father,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkBToC := NewLink(
		"linkBToC",
		containerB,
		containerC,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkCToD := NewLink(
		"linkCToD",
		containerC,
		containerD,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkBToE := NewLink(
		"linkBToE",
		containerB,
		containerE,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkBToC, linkCToD, linkBToE}

	filter := Filter{}
	ret := filter.LinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(containerA, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkAToB
	// linkBToC
	// linkBToE
	// linkCToD
}

func ExampleLinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_2() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 10, 340, 100, 100)
	containerE := NewContainerWidthXY("containerE", 10, 450, 100, 100)
	containerF := NewContainerWidthXY("containerF", 10, 560, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkFToFather",
		father,
		containerF,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkBToC := NewLink(
		"linkBToC",
		containerB,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkCToD := NewLink(
		"linkCToD",
		containerC,
		containerD,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkBToE := NewLink(
		"linkBToE",
		containerB,
		containerE,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkBToC, linkCToD, linkBToE}

	filter := Filter{}
	ret := filter.LinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(containerF, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkFToFather
}

func ExampleLinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_3() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 10, 340, 100, 100)
	containerE := NewContainerWidthXY("containerE", 10, 450, 100, 100)
	containerF := NewContainerWidthXY("containerF", 10, 560, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkFToFather",
		father,
		containerF,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkAToB := NewLink(
		"linkAToB",
		containerA,
		containerB,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkBToC := NewLink(
		"linkBToC",
		containerB,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkCToD := NewLink(
		"linkCToD",
		containerC,
		containerD,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkBToE := NewLink(
		"linkBToE",
		containerB,
		containerE,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkBToC, linkCToD, linkBToE}

	filter := Filter{}
	ret := filter.LinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(father, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkFToFather
}

func ExampleLinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_4() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 10, 340, 100, 100)
	containerE := NewContainerWidthXY("containerE", 10, 450, 100, 100)
	containerF := NewContainerWidthXY("containerF", 10, 560, 100, 100)
	containerG := NewContainerWidthXY("containerG", 10, 670, 100, 100)
	containerH := NewContainerWidthXY("containerG", 10, 780, 100, 100)
	containerI := NewContainerWidthXY("containerI", 10, 890, 100, 100)

	linkAToFather := NewLinkWithFather(
		"linkFToFather",
		father,
		containerF,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkAToB := NewLink(
		"linkAToB",
		containerB,
		containerA,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkBToC := NewLink(
		"linkBToC",
		containerC,
		containerB,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkCToD := NewLink(
		"linkCToD",
		containerD,
		containerC,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkBToE := NewLink(
		"linkBToE",
		containerE,
		containerB,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkGToF := NewLink(
		"linkGToF",
		containerF,
		containerG,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkGToH := NewLink(
		"linkGToH",
		containerH,
		containerG,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	linkFToI := NewLink(
		"linkFToI",
		containerI,
		containerF,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkBToC, linkCToD, linkBToE, linkGToF, linkGToH, linkFToI}

	filter := Filter{}
	ret := filter.LinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(father, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkFToFather
	// linkGToF
	// linkFToI
	// linkGToH
}

//  +-father------------------------------------------------------------------+
//  |                                                                         |
//  |     +-containerA----------------+     +-containerD----------------+     |
//  |     |                           |     |                           |     |
//  O<-+->O 1: a to father            O<-+  |                 5: d to f O->+  |
//  |  |  |                           |  |  |                           |  |  |
//  |  |  +---------------------------+  |  +---------------------------+  |  |
//  |  |                                 |                                 |  |
//  |  |  +-containerB----------------+  |  +-containerE----------------+  |  |
//  |  |  |                           |  |  |                           |  |  |
//  |  |  |                 4: b to a O->+  |                 6: e to f O->+  |
//  |  |  |                           |  |  |                           |  |  |
//  |  |  +---------------------------+  |  +---------------------------+  |  |
//  |  |                                 |                                 |  |
//  |  |  +-containerC----------------+  |  +-containerF----------------+  |  |
//  |  |  |                           |  |  |                           |  |  |
//  |  +<-O 2: c to a          c to a O->+<-O 3: a to f                 O<-+  |
//  |     |                           |     |                           |     |
//  |     +---------------------------+     +---------------------------+     |
//  |                                                                         |
//  +-------------------------------------------------------------------------+
func ExampleLinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_5() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 120, 10, 100, 100)
	containerE := NewContainerWidthXY("containerE", 230, 120, 100, 100)
	containerF := NewContainerWidthXY("containerF", 340, 230, 100, 100)

	// 1: a to father
	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	// 2: c to a
	linkCToA := NewLink(
		"linkCToA",
		containerA,
		containerC,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	// 3: a to f
	linkAToF := NewLink(
		"linkAToF",
		containerA,
		containerF,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnRight,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	// 4: b to a
	linkBToA := NewLink(
		"linkBToA",
		containerA,
		containerB,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	// 5: d to f
	linkDToF := NewLink(
		"linkDToF",
		containerF,
		containerD,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	// 6: e to f
	linkEToF := NewLink(
		"linkEToF",
		containerF,
		containerE,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkCToA, linkAToF, linkBToA, linkDToF, linkEToF}

	filter := Filter{}
	ret := filter.LinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(father, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkAToFather
	// linkCToA
	// linkAToF
	// linkBToA
	// linkDToF
	// linkEToF
}

//  +-father------------------------------------------------------------------+
//  |                                                                         |
//  |     +-containerA----------------+     +-containerD----------------+     |
//  |     |                           |     |                           |     |
//  O<-+->O 1: a to father            O<-+->O                           O<-+  |
//  |  |  |                           |  |  |                           |  |  |
//  |  |  +---------------------------+  |  +---------------------------+  |  |
//  |  |                                 |                                 |  |
//  |  |  +-containerB----------------+  |  +-containerE----------------+  |  |
//  |  |  |                           |  |  |                           |  |  |
//  |  |  |                 4: b to a O->+<-O 5: e to d  <---->  e to d O->+  |
//  |  |  |                           |  |  |                           |     |
//  |  |  +---------------------------+  |  +---------------------------+     |
//  |  |                                 |                                    |
//  |  |  +-containerC----------------+  |  +-containerF----------------+     |
//  |  |  |                           |  |  |                           |     |
//  |  +<-O 2: c to a  <---->  c to a O->+<-O 3: a to f                 O     |
//  |     |                           |     |                           |     |
//  |     +---------------------------+     +---------------------------+     |
//  |                                                                         |
//  +-------------------------------------------------------------------------+
func ExampleLinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_6() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 120, 10, 100, 100)
	containerE := NewContainerWidthXY("containerE", 230, 120, 100, 100)
	containerF := NewContainerWidthXY("containerF", 340, 230, 100, 100)

	// 1: a to father
	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	// 2: c to a
	linkCToA := NewLink(
		"linkCToA",
		containerA,
		containerC,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	// 3: a to f
	linkAToF := NewLink(
		"linkAToF",
		containerF,
		containerA,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnRight,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	// 4: b to a
	linkBToA := NewLink(
		"linkBToA",
		containerA,
		containerB,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	// 5: e to d
	linkEToD := NewLink(
		"linkEToD",
		containerD,
		containerE,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkCToA, linkAToF, linkBToA, linkEToD}

	filter := Filter{}
	ret := filter.LinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(father, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkAToFather
	// linkCToA
	// linkAToF
	// linkBToA
}

//  +-father------------O-----------------------------------------------------+
//  |                   | 1: a to d                                           |
//  |                   +---------------------------------+                   |
//  |                   |                                 |                   |
//  |     +-containerA--O-------------+     +-containerD--O-------------+     |
//  |     |      2: father to a       |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |         3: a to b         |     |                           |     |
//  |     +-------------O-------------+     +-------------O-------------+     |
//  |                   | 1: a to d                       |                   |
//  |                   +---------------------------------+                   |
//  |                   |                                                     |
//  |                   +---------------------------------+                   |
//  |                   | 4: b to e                       |                   |
//  |     +-containerB--O-------------+     +-containerE--O-------------+     |
//  |     |                           |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |         5: b to c         |     |         8: e to f         |     |
//  |     +-------------O-------------+     +-------------O-------------+     |
//  |                   |                                 |                   |
//  |                   |                                 |                   |
//  |                   |                                 |                   |
//  |     +-containerC--O-------------+     +-containerF--O-------------+     |
//  |     |                           |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |       6: c to father      |     |                           |     |
//  |     +-------------O-------------+     +-------------O-------------+     |
//  |                   | 7: c to f                       |                   |
//  |                   +---------------------------------+                   |
//  |                   |                                                     |
//  +-------------------O-----------------------------------------------------+
func ExampleLinkVerticalFilterContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainerDirectlyOrIndirectlyLinked_1() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 120, 10, 100, 100)
	containerE := NewContainerWidthXY("containerE", 230, 120, 100, 100)
	containerF := NewContainerWidthXY("containerF", 340, 230, 100, 100)

	linkAToD := NewLink(
		"1:linkAToD",
		containerA,
		containerD,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkFatherToA := NewLinkWithFather(
		"2:linkFatherToA",
		containerA,
		father,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkAToB := NewLink(
		"3:linkAToB",
		containerB,
		containerA,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkBToE := NewLink(
		"4:linkBToE",
		containerE,
		containerB,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkCToE := NewLink(
		"5:linkCToB",
		containerB,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkCToFather := NewLinkWithFather(
		"6:linkCToFather",
		containerC,
		father,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkFToC := NewLink(
		"7:linkFToC",
		containerC,
		containerF,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkFToE := NewLink(
		"8:linkFToE",
		containerE,
		containerF,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToD, linkFatherToA, linkAToB, linkBToE, linkCToE, linkCToFather, linkFToC, linkFToE}
	filter := Filter{}

	fatherFound := filter.FilterContainerFather(listLink)
	if fatherFound == nil && fatherFound != father {
		fmt.Println("error")
	}

	ret := filter.LinkVerticalFilterContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainerDirectlyOrIndirectlyLinked(fatherFound, listLink)
	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// 2:linkFatherToA
	// 6:linkCToFather
	// 1:linkAToD
	// 3:linkAToB
	// 5:linkCToB
	// 7:linkFToC
	// 4:linkBToE
	// 8:linkFToE
}

//  +-father------------O-----------------------------------------------------+
//  |                   | 1: a to d                                           |
//  |                   +---------------------------------+                   |
//  |                   |                                 |                   |
//  |     +-containerA--O-------------+     +-containerD--O-------------+     |
//  |     |      2: father to a       |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |         3: a to b         |     |                           |     |
//  |     +-------------O-------------+     +-------------O-------------+     |
//  |                   | 1: a to d                       |                   |
//  |                   +---------------------------------+                   |
//  |                                                                         |
//  |     +-containerB--X-------------+     +-containerE--X-------------+     |
//  |     |                           |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |         4: b to c         |     |         6: e to f         |     |
//  |     +-------------O-------------+     +-------------O-------------+     |
//  |                   |                                 |                   |
//  |                   |                                 |                   |
//  |                   |                                 |                   |
//  |     +-containerC--O-------------+     +-containerF--O-------------+     |
//  |     |                           |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |       5: c to father      |     |                           |     |
//  |     +-------------O-------------+     +-------------X-------------+     |
//  |                   |                                                     |
//  |                   |                                                     |
//  |                   |                                                     |
//  +-------------------O-----------------------------------------------------+
func ExampleLinkVerticalFilterContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainerDirectlyOrIndirectlyLinked_2() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 120, 10, 100, 100)
	containerE := NewContainerWidthXY("containerE", 230, 120, 100, 100)
	containerF := NewContainerWidthXY("containerF", 340, 230, 100, 100)

	linkAToD := NewLink(
		"1:linkAToD",
		containerA,
		containerD,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkFatherToA := NewLinkWithFather(
		"2:linkFatherToA",
		containerA,
		father,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkAToB := NewLink(
		"3:linkAToB",
		containerB,
		containerA,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkCToE := NewLink(
		"4:linkCToB",
		containerB,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkCToFather := NewLinkWithFather(
		"5:linkCToFather",
		containerC,
		father,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkFToE := NewLink(
		"6:linkFToE",
		containerE,
		containerF,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	listLink := []*Link{linkAToD, linkFatherToA, linkAToB, linkCToE, linkCToFather, linkFToE}
	filter := Filter{}

	fatherFound := filter.FilterContainerFather(listLink)
	if fatherFound == nil && fatherFound != father {
		fmt.Println("error")
	}

	ret := filter.LinkVerticalFilterContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainerDirectlyOrIndirectlyLinked(fatherFound, listLink)
	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// 2:linkFatherToA
	// 5:linkCToFather
	// 1:linkAToD
	// 3:linkAToB
	// 4:linkCToB
}

//  +-father------------O-----------------------------------------------------+
//  |                   | 1: a to d                                           |
//  |                   +---------------------------------+                   |
//  |                   |                                 |                   |
//  |     +-containerA--O-------------+     +-containerD--O-------------+     |
//  |     |      2: father to a       |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |         3: a to b         |     |                           |     |
//  |     +-------------O-------------+     +-------------O-------------+     |
//  |                   | 1: a to d                       |                   |
//  |                   +---------------------------------+                   |
//  |                                                                         |
//  |     +-containerB--X-------------+     +-containerE--X-------------+     |
//  |     |                           |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |         4: b to c         |     |         6: e to f         |     |
//  |     +-------------O-------------+     +-------------O-------------+     |
//  |                   |                                 |                   |
//  |                   |                                 |                   |
//  |                   |                                 |                   |
//  |     +-containerC--O-------------+     +-containerF--O-------------+     |
//  |     |                           |     |                           |     |
//  |     |                           |     |                           |     |
//  |     |       5: c to father      |     |                           |     |
//  |     +-------------O-------------+     +-------------X-------------+     |
//  |                   |                                                     |
//  |                   |                                                     |
//  |                   |                                                     |
//  +-------------------O-----------------------------------------------------+
func ExampleLinkVerticalFilterContainersWithErrorAndFindContainersWithIsNotAlignsFromTheTopOrTheBottomInRelationToFatherDirectlyOrIndirectlyLinked_1() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 120, 10, 100, 100)
	containerE := NewContainerWidthXY("containerE", 230, 120, 100, 100)
	containerF := NewContainerWidthXY("containerF", 340, 230, 100, 100)

	linkAToD := NewLink(
		"1:linkAToD",
		containerA,
		containerD,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkFatherToA := NewLinkWithFather(
		"2:linkFatherToA",
		containerA,
		father,
		KWallTopContainerBLinkOnTopToContainerALinkOnTop,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkAToB := NewLink(
		"3:linkAToB",
		containerB,
		containerA,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkCToE := NewLink(
		"4:linkCToB",
		containerB,
		containerC,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	linkCToFather := NewLinkWithFather(
		"5:linkCToFather",
		containerC,
		father,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkFToE := NewLink(
		"6:linkFToE",
		containerE,
		containerF,
		KWallTopContainerBLinkOnTopToContainerALinkOnBottom,
		KWallLeftNotSet,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	listOfContainers := []*Dimensions{father, containerA, containerB, containerC, containerD, containerE, containerF}
	listOfLinks := []*Link{linkAToD, linkFatherToA, linkAToB, linkCToE, linkCToFather, linkFToE}
	filter := Filter{}

	err, ret := filter.LinkVerticalFilterContainersWithErrorAndFindContainersWithIsNotAlignsFromTheTopOrTheBottomInRelationToFatherDirectlyOrIndirectlyLinked(listOfContainers, listOfLinks)
	if err != nil {
		fmt.Println("error")
	}
	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// containerE
	// containerF
}

//  +-father------------------------------------------------------------------+
//  |                                                                         |
//  |     +-containerA----------------+     +-containerD----------------+     |
//  |     |                           |     |                           |     |
//  O<-+->O 1: a to father            O<-+->O                           O<-+  |
//  |  |  |                           |  |  |                           |  |  |
//  |  |  +---------------------------+  |  +---------------------------+  |  |
//  |  |                                 |                                 |  |
//  |  |  +-containerB----------------+  |  +-containerE----------------+  |  |
//  |  |  |                           |  |  |                           |  |  |
//  |  |  |                 4: b to a O->+<-O 5: e to d  <---->  e to d O->+  |
//  |  |  |                           |  |  |                           |     |
//  |  |  +---------------------------+  |  +---------------------------+     |
//  |  |                                 |                                    |
//  |  |  +-containerC----------------+  |  +-containerF----------------+     |
//  |  |  |                           |  |  |                           |     |
//  |  +<-O 2: c to a  <---->  c to a O->+  X                           X     |
//  |     |                           |     |                           |     |
//  |     +---------------------------+     +---------------------------+     |
//  |                                                                         |
//  +-------------------------------------------------------------------------+
func ExampleLinkHorizontalFilterContainersWithErrorAndFindContainersWithIsNotAlignsFromTheTopOrTheBottomInRelationToFatherDirectlyOrIndirectlyLinked_1() {
	father := NewContainer("father", 300, 600)
	containerA := NewContainerWidthXY("containerA", 10, 10, 100, 100)
	containerB := NewContainerWidthXY("containerB", 10, 120, 100, 100)
	containerC := NewContainerWidthXY("containerC", 10, 230, 100, 100)
	containerD := NewContainerWidthXY("containerD", 120, 10, 100, 100)
	containerE := NewContainerWidthXY("containerE", 230, 120, 100, 100)
	containerF := NewContainerWidthXY("containerF", 340, 230, 100, 100)

	// 1: a to father
	linkAToFather := NewLinkWithFather(
		"linkAToFather",
		father,
		containerA,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightNotSet,
		KWallBottomNotSet,
	)

	// 2: c to a
	linkCToA := NewLink(
		"linkCToA",
		containerA,
		containerC,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	// 4: b to a
	linkBToA := NewLink(
		"linkBToA",
		containerA,
		containerB,
		KWallTopNotSet,
		KWallLeftNotSet,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	// 5: e to d
	linkEToD := NewLink(
		"linkEToD",
		containerD,
		containerE,
		KWallTopNotSet,
		KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KWallRightContainerBLinkOnRightToContainerALinkOnRight,
		KWallBottomNotSet,
	)

	listOfContainers := []*Dimensions{father, containerA, containerB, containerC, containerD, containerE, containerF}
	listOfLinks := []*Link{linkAToFather, linkCToA, linkBToA, linkEToD}

	filter := Filter{}
	err, ret := filter.LinkHorizontalFilterContainersWithErrorAndFindContainersWithIsNotAlignsFromTheTopOrTheBottomInRelationToFatherDirectlyOrIndirectlyLinked(listOfContainers, listOfLinks)
	if err != nil {
		fmt.Println("error")
	}
	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// containerD
	// containerE
	// containerF
}
