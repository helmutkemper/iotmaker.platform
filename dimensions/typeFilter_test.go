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

func ExampleLinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_1() {
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
	fatherFromFilter := filter.FilterContainerFather(listLink)
	if fatherFromFilter == nil {
		fmt.Println("Error!")
	}
	ret := filter.LinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(fatherFromFilter, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkAToB
	// linkBToC
	// linkBToE
	// linkCToD
}

func ExampleLinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_2() {
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
	ret := filter.LinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(containerF, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkFToFather
}

func ExampleLinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_3() {
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
	ret := filter.LinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(father, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkFToFather
}

func ExampleLinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_4() {
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
		KCornerTopContainerBLinkOnTopToContainerALinkOnTop,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomContainerBLinkOnBottomToContainerALinkOnBottom,
	)

	linkAToB := NewLink(
		"linkAToB",
		containerB,
		containerA,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkBToC := NewLink(
		"linkBToC",
		containerC,
		containerB,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkCToD := NewLink(
		"linkCToD",
		containerD,
		containerC,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkBToE := NewLink(
		"linkBToE",
		containerE,
		containerB,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkGToF := NewLink(
		"linkGToF",
		containerF,
		containerG,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkGToH := NewLink(
		"linkGToH",
		containerH,
		containerG,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	linkFToI := NewLink(
		"linkFToI",
		containerI,
		containerF,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkAToB, linkBToC, linkCToD, linkBToE, linkGToF, linkGToH, linkFToI}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(father, listLink)

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
func ExampleLinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_5() {
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
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightNotSet,
		KCornerBottomNotSet,
	)

	// 2: c to a
	linkCToA := NewLink(
		"linkCToA",
		containerA,
		containerC,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightNotSet,
		KCornerBottomNotSet,
	)

	// 3: a to f
	linkAToF := NewLink(
		"linkAToF",
		containerA,
		containerF,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnRight,
		KCornerRightNotSet,
		KCornerBottomNotSet,
	)

	// 4: b to a
	linkBToA := NewLink(
		"linkBToA",
		containerA,
		containerB,
		KCornerTopNotSet,
		KCornerLeftNotSet,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	// 5: d to f
	linkDToF := NewLink(
		"linkDToF",
		containerF,
		containerD,
		KCornerTopNotSet,
		KCornerLeftNotSet,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	// 6: e to f
	linkEToF := NewLink(
		"linkEToF",
		containerF,
		containerE,
		KCornerTopNotSet,
		KCornerLeftNotSet,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkCToA, linkAToF, linkBToA, linkDToF, linkEToF}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(father, listLink)

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
func ExampleLinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked_6() {
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
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightNotSet,
		KCornerBottomNotSet,
	)

	// 2: c to a
	linkCToA := NewLink(
		"linkCToA",
		containerA,
		containerC,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	// 3: a to f
	linkAToF := NewLink(
		"linkAToF",
		containerF,
		containerA,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnRight,
		KCornerRightNotSet,
		KCornerBottomNotSet,
	)

	// 4: b to a
	linkBToA := NewLink(
		"linkBToA",
		containerA,
		containerB,
		KCornerTopNotSet,
		KCornerLeftNotSet,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	// 5: e to d
	linkEToD := NewLink(
		"linkEToD",
		containerD,
		containerE,
		KCornerTopNotSet,
		KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft,
		KCornerRightContainerBLinkOnRightToContainerALinkOnRight,
		KCornerBottomNotSet,
	)

	listLink := []*Link{linkAToFather, linkCToA, linkAToF, linkBToA, linkEToD}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(father, listLink)

	for k := range ret {
		fmt.Printf("%v\n", ret[k].ToDebug)
	}

	// Output:
	// linkAToFather
	// linkCToA
	// linkAToF
	// linkBToA
}
