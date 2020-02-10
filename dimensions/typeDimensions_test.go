package dimensions

import "fmt"

// en: Explication:
//
// pt_br: Explicação
//
//  +-father----------O-----------------+
//  |                 |                 |
//  |   +-containerA--O-------------+   |
//  |   |                           |   |
//  O---O                           O---O
//  |   |                           |   |
//  |   +-------------O-------------+   |
//  |                 |                 |
//  +-----------------O-----------------+
func ExampleLink_SimpleLinkAssembly_1() {
	var err error

	var linkFatherTopToContainerATop = &Link{}
	var linkFatherLeftToContainerALeft = &Link{}
	var linkFatherRightToContainerARight = &Link{}
	var linkFatherBottomToContainerABottom = &Link{}

	father := NewContainer(600, 300)
	containerA := NewContainerWidthXY(10, 10, 300, 100)

	linkFatherTopToContainerATop = NewLink(containerA, KContainerBLinkOnTopToContainerALinkOnTop)
	linkFatherLeftToContainerALeft = NewLink(containerA, KContainerBLinkOnLeftToContainerALinkOnLeft)
	linkFatherRightToContainerARight = NewLink(containerA, KContainerBLinkOnRightToContainerALinkOnRight)
	linkFatherBottomToContainerABottom = NewLink(containerA, KContainerBLinkOnBottomToContainerALinkOnBottom)

	collection := LinkCollection{
		Top:    linkFatherTopToContainerATop,
		Left:   linkFatherLeftToContainerALeft,
		Right:  linkFatherRightToContainerARight,
		Bottom: linkFatherBottomToContainerABottom,
	}

	err = linkFatherTopToContainerATop.SimpleLinkAssembly(
		father,
		collection,
	)
	if err != nil {
		fmt.Printf("SimpleLinkAssembly().error: %v\n", err.Error())
	}

	fmt.Printf("father.X: %v\n", father.X)
	fmt.Printf("father.Y: %v\n", father.Y)
	fmt.Printf("father.W: %v\n", father.Width)
	fmt.Printf("father.H: %v\n", father.Height)

	fmt.Printf("containerA.X: %v\n", containerA.X)
	fmt.Printf("containerA.Y: %v\n", containerA.Y)
	fmt.Printf("containerA.W: %v\n", containerA.Width)
	fmt.Printf("containerA.H: %v\n", containerA.Height)

	// Output:
	// father.X: 0
	// father.Y: 0
	// father.W: 600
	// father.H: 300
	// containerA.X: 150
	// containerA.Y: 100
	// containerA.W: 300
	// containerA.H: 100
}

func ExampleLink_SimpleLinkAssembly_2() {
	var err error

	var linkFatherTopToContainerATop = &Link{}
	var linkFatherLeftToContainerALeft = &Link{}
	var linkFatherRightToContainerARight = &Link{}
	var linkFatherBottomToContainerABottom = &Link{}

	containerA := NewContainer(300, 150)
	containerB := NewContainerWidthXY(10, 10, 600, 300)

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

	fmt.Printf("containerA.X: %v\n", containerA.X)
	fmt.Printf("containerA.Y: %v\n", containerA.Y)
	fmt.Printf("containerA.W: %v\n", containerA.Width)
	fmt.Printf("containerA.H: %v\n", containerA.Height)

	fmt.Printf("containerB.X: %v\n", containerB.X)
	fmt.Printf("containerB.Y: %v\n", containerB.Y)
	fmt.Printf("containerB.W: %v\n", containerB.Width)
	fmt.Printf("containerB.H: %v\n", containerB.Height)

	// Output:
	//
}
