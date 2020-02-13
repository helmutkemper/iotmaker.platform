package dimensions

import (
	"fmt"
	"reflect"
)

func ExampleNewLink() {

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

	list := []*Link{linkAToFather, linkAToB, linkAToC}

	filter := Filter{}
	ret := filter.LinkAssemblyHorizontalFilterContainerIsCentralizedInRelationToTheFather(list)

	if len(ret) == 1 && reflect.DeepEqual(ret[0], *linkAToFather) {
		fmt.Println("passou")
	}

	ret = filter.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftAndTheRightInRelationToAnotherContainer(containerA, list)
	if len(ret) == 3 {

		afPass := reflect.DeepEqual(ret[0], *linkAToFather)
		abPass := reflect.DeepEqual(ret[1], *linkAToB)
		acPass := reflect.DeepEqual(ret[2], *linkAToC)
		pass := afPass && abPass && acPass
		if pass == true {
			fmt.Println("passou")
		}
	}

	ret = filter.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftAndTheRightInRelationToAnotherContainer(containerB, list)
	if len(ret) == 1 {

		abPass := reflect.DeepEqual(ret[0], *linkAToB)
		if abPass == true {
			fmt.Println("passou")
		}
	}

	// output:
	// passou
	// passou
	// passou
}
