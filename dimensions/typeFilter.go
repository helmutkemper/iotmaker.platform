package dimensions

import (
	"errors"
)

type Filter struct{}

func (el Filter) findDimensionsInList(valueToFind *Dimensions, list []*Dimensions) bool {
	for _, valueInList := range list {
		if valueToFind == valueInList {
			return true
		}
	}

	return false
}

func (el Filter) findLinkInList(link *Link, list []*Link) bool {
	for k := range list {
		if list[k].ContainerA == link.ContainerA && list[k].ContainerB == link.ContainerB {
			return true
		}
	}

	return false
}

func (el Filter) SearchContainerFather(list []*Link) *Dimensions {
	for _, link := range list {
		if link.ContainerAFather == true {
			return link.ContainerA
		}
	}

	return nil
}

func (el Filter) CheckIfFatherExistsAndHasOnlyOne(list []*Link) bool {
	counter := 0
	for _, link := range list {
		if link.ContainerAFather == true {
			counter += 1
		}
	}

	return counter == 1
}

//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA--X-------------+     |
//  |     |                           |     |
//  |     X                           X     |
//  |     |                           |     |
//  |     +-------------X-------------+     |
//  |                                       |
//  +---------------------------------------+
func (el Filter) SearchForFloatingContainers(listLink []*Link, listDimensions []*Dimensions) []*Dimensions {
	notFoundList := make([]*Dimensions, 0)
	foundList := make([]*Dimensions, 0)

	for _, dimension := range listDimensions {
		for _, link := range listLink {
			if link.ContainerA == dimension && !el.findDimensionsInList(link.ContainerA, foundList) {
				foundList = append(foundList, link.ContainerA)
			}

			if link.ContainerB == dimension && !el.findDimensionsInList(link.ContainerB, foundList) {
				foundList = append(foundList, link.ContainerB)
			}
		}
	}

	for _, dimension := range listDimensions {
		if !el.findDimensionsInList(dimension, foundList) {

			if !el.findDimensionsInList(dimension, notFoundList) {
				notFoundList = append(notFoundList, dimension)
			}

		}
	}

	return notFoundList
}

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
func (el Filter) LinkHorizontalFilterContainerIsCentralizedInRelationToTheFather(list []*Link) []*Link {

	ret := make([]*Link, 0)
	for _, link := range list {
		fatherPass := link.ContainerAFather == true
		bRightARightPass := link.Right == KWallRightContainerBLinkOnRightToContainerALinkOnRight
		bLeftALeftPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft
		//bLeftARightPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnRight

		pass := bRightARightPass && bLeftALeftPass && fatherPass
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

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
func (el Filter) LinkHorizontalFilterEachContainerIsAlignsFromTheRightInRelationToTheFather(list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		fatherPass := link.ContainerAFather == true
		bRightARightPass := link.Right == KWallRightContainerBLinkOnRightToContainerALinkOnRight
		leftNotSetPass := link.Left == KWallLeftNotSet

		pass := bRightARightPass && leftNotSetPass && fatherPass
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

//  Ponto de vista horizontal
//  Option: cada container é alinhado a direita em relação ao container A
//  Correto:
//  +-father----------------------------------+
//  |                                         |
//  |     +-containerA----------------+       |
//  |     |                           |       |
//  |     |                           O--+-+  |
//  |     |                           |  | |  |
//  |     +---------------------------+  | |  |
//  |                                    | |  |
//  |     +-containerB----------------+  | |  |
//  |     |                           |  | |  |
//  |     |                           O--+ |  |
//  |     |                           |    |  |
//  |     +---------------------------+    |  |
//  |                                      |  |
//  |     +-containerC----------------+    |  |
//  |     |                           |    |  |
//  |     |                           O----+  |
//  |     |                           |       |
//  |     +---------------------------+       |
//  |                                         |
//  +-----------------------------------------+
func (el Filter) LinkHorizontalFilterEachContainerIsAlignsFromTheRightInRelationToAnother(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		bRightARightPass := link.Right == KWallRightContainerBLinkOnRightToContainerALinkOnRight
		leftNotSetPass := link.Left == KWallLeftNotSet

		pass := bRightARightPass && leftNotSetPass && (containerAPass || containerBPass)
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

//  Ponto de vista horizontal
//  Option: cada container é alinhado a direita em relação ao container A
//  Correto:
//  +-father------------------------------------+
//  |                                           |
//  |       +-containerA----------------+       |
//  |       |                           |       |
//  |  +-+--O                           O--+-+  |
//  |  | |  |                           |  | |  |
//  |  | |  +---------------------------+  | |  |
//  |  | |                                 | |  |
//  |  | |  +-containerB----------------+  | |  |
//  |  | |  |                           |  | |  |
//  |  | +--O                           O--+ |  |
//  |  |    |                           |    |  |
//  |  |    +---------------------------+    |  |
//  |  |                                     |  |
//  |  |    +-containerC----------------+    |  |
//  |  |    |                           |    |  |
//  |  +----O                           O----+  |
//  |       |                           |       |
//  |       +---------------------------+       |
//  |                                           |
//  +-------------------------------------------+
func (el Filter) LinkHorizontalFilterEachContainerIsAlignsFromTheRightAndTheLeftInRelationToAnother(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		bRightARightPass := link.Right == KWallRightContainerBLinkOnRightToContainerALinkOnRight
		bLeftALeftPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft
		//leftPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnRight

		pass := bRightARightPass && bLeftALeftPass && (containerAPass || containerBPass)
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

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
func (el Filter) LinkHorizontalFilterEachContainerIsAlignsFromTheLeftInRelationToTheFather(list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		fatherPass := link.ContainerAFather == true
		rightPass := link.Right == KWallRightNotSet
		bLeftALeftPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft
		//leftPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnRight

		pass := rightPass && bLeftALeftPass && fatherPass
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

//  Ponto de vista horizontal
//  Option: cada container é alinhado a direita em relação ao container A
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |       +-containerA----------------+   |
//  |       |                           |   |
//  |  +-+--O                           |   |
//  |  | |  |                           |   |
//  |  | |  +---------------------------+   |
//  |  | |                                  |
//  |  | |  +-containerB----------------+   |
//  |  | |  |                           |   |
//  |  | +--O                           |   |
//  |  |    |                           |   |
//  |  |    +---------------------------+   |
//  |  |                                    |
//  |  |    +-containerC----------------+   |
//  |  |    |                           |   |
//  |  +----O                           |   |
//  |       |                           |   |
//  |       +---------------------------+   |
//  |                                       |
//  +---------------------------------------+
func (el Filter) LinkHorizontalFilterEachContainerIsAlignsFromTheLeftInRelationToAnother(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		rightNotSetPass := link.Right == KWallRightNotSet
		bLeftALeftPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft

		pass := rightNotSetPass && bLeftALeftPass && (containerAPass || containerBPass)
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |  +<-O                           O->+  |
//  |  |  |                           |  |  |
//  |  +->O                           O<-+  |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
func (el Filter) LinkErrorFilterContainerIsLinkedInRelationToItSelf(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container

		pass := containerAPass && containerBPass
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

//  Ponto de vista horizontal
//  Option: D - os containers a e b têm o seu tamanho ajustados pelo maior
//              o container a é centralizado pelo pai
//              o container b y é ajustado por a y
//  Correto:
//  +-father--------------------------------+  +-father--------------------------------+
//  |                                       |  |                                       |
//  |     +-containerA----------------+     |  |     +-containerA----------------+     |
//  |     |                           |     |  |     |                           |     |
//  |  +--O a link to b               O--+  |  |  +>>O                           O<<+  |
//  |  |  |                           |  |  |  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |  |  |  +---------------------------+  |  |
//  |  |                                 |  |  |  |                                 |  |
//  |  |  +-containerB----------------+  |  |  |  |  +-containerB----------------+  |  |
//  |  +--O                           O--+  |  |  |  |                           |  |  |
//  |     | a link to b               |     |  |  |<<O b link to a               O>>|  |
//  |     | b link to c               |     |  |  |  |                           |  |  |
//  |  +--O                           O--+  |  |  |  +---------------------------+  |  |
//  |  |  +---------------------------+  |  |  |  |                                 |  |
//  |  |                                 |  |  |  |  +-containerC----------------+  |  |
//  |  |  +-containerC----------------+  |  |  |  |  |                           |  |  |
//  |  |  |                           |  |  |  |  +<<O c link to a               O>>+  |
//  |  +--O b link to c               O--+  |  |     |                           |     |
//  |     |                           |     |  |     +---------------------------+     |
//  |     +---------------------------+     |  |                                       |
//  |                                       |  |                                       |
//  +---------------------------------------+  +---------------------------------------+
func (el Filter) LinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(container *Dimensions, listLinks []*Link) []*Link {

	listOfLinksToVerify := el.LinkHorizontalFilterEachContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainer(container, listLinks)
	containerTested := []*Dimensions{container}
	containersToVerify := make([]*Dimensions, 0)

	for {
		for k := range listOfLinksToVerify {
			if el.findDimensionsInList(listOfLinksToVerify[k].ContainerA, containerTested) == false {
				containersToVerify = append(containersToVerify, listOfLinksToVerify[k].ContainerA)
			}

			if el.findDimensionsInList(listOfLinksToVerify[k].ContainerB, containerTested) == false {
				containersToVerify = append(containersToVerify, listOfLinksToVerify[k].ContainerB)
			}
		}

		if len(containersToVerify) == 0 {
			break
		}

		pass := false
		for k := range containersToVerify {
			container = containersToVerify[k]
			if el.findDimensionsInList(container, containerTested) == false {
				pass = true
				break
			}
		}

		if pass == false {
			break
		}

		containerTested = append(containerTested, container)
		newLinkList := el.LinkHorizontalFilterEachContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainer(container, listLinks)

		for k := range newLinkList {
			if el.findLinkInList(newLinkList[k], listOfLinksToVerify) == false {
				listOfLinksToVerify = append(listOfLinksToVerify, newLinkList[k])
			}
		}
	}

	return listOfLinksToVerify
}

//  Ponto de vista horizontal
//  Option: cada container é alinhado a direita em relação ao container A
//  Correto:
//  +-father---------------------------------------------------------------+
//  |                                                                      |
//  |     +-containerA----------------+   +-containerD----------------+    |
//  |     |                           |   |                           |    |
//  |  +--O                           O-+ |                           O-+  |
//  |  |  |                           | | |                           | |  |
//  |  |  +---------------------------+ | +---------------------------+ |  |
//  |  |                                |                               |  |
//  |  |  +-containerB----------------+ | +-containerE----------------+ |  |
//  |  |  |                           | | |                           | |  |
//  |  |  |                           O-+ |                           O-+  |
//  |  |  |                           | | |                           | |  |
//  |  |  +---------------------------+ | +---------------------------+ |  |
//  |  |                                |                               |  |
//  |  |  +-containerC----------------+ | +-containerF----------------+ |  |
//  |  |  |                           | | |                           | |  |
//  |  +--O                           O-+-O                           O-+  |
//  |     |                           |   |                           |    |
//  |     +---------------------------+   +---------------------------+    |
//  |                                                                      |
//  +----------------------------------------------------------------------+
func (el Filter) LinkHorizontalFilterEachContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainer(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		bRightARightPass := link.Right == KWallRightContainerBLinkOnRightToContainerALinkOnRight
		bLeftALeftPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft
		bLeftBRightPass := link.Left == KWallLeftContainerBLinkOnLeftToContainerALinkOnRight

		pass := (bRightARightPass || bLeftALeftPass || bLeftBRightPass) && (containerAPass || containerBPass)
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

func (el Filter) LinkVerticalFilterEachContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainer(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		topBTopATopPass := link.Top == KWallTopContainerBLinkOnTopToContainerALinkOnTop
		topBTopABottomPass := link.Top == KWallTopContainerBLinkOnTopToContainerALinkOnBottom
		bottomBTopABottomPass := link.Bottom == KWallBottomContainerBLinkOnTopToContainerALinkOnBottom
		bottomBBottomABottomPass := link.Bottom == KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom

		pass := (topBTopATopPass || topBTopABottomPass || bottomBTopABottomPass || bottomBBottomABottomPass) && (containerAPass || containerBPass)
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

func (el Filter) LinkHorizontalFilterContainersWithErrorAndFindContainersWithIsNotAlignsFromTheTopOrTheBottomInRelationToFatherDirectlyOrIndirectlyLinked(listOfContainers []*Dimensions, listOfLinks []*Link) (error, []*Dimensions) {

	fatherContainer := el.SearchContainerFather(listOfLinks)
	if fatherContainer == nil {
		return errors.New("father containers not found in containers list"), []*Dimensions{}
	}
	listOfLinksToVerify := el.LinkHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(fatherContainer, listOfLinks)
	return nil, el.SearchForFloatingContainers(listOfLinksToVerify, listOfContainers)
}

func (el Filter) LinkVerticalFilterContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainerDirectlyOrIndirectlyLinked(container *Dimensions, listLinks []*Link) []*Link {

	listOfLinksToVerify := el.LinkVerticalFilterEachContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainer(container, listLinks)
	containerTested := []*Dimensions{container}
	containersToVerify := make([]*Dimensions, 0)

	for {
		for k := range listOfLinksToVerify {
			if el.findDimensionsInList(listOfLinksToVerify[k].ContainerA, containerTested) == false {
				containersToVerify = append(containersToVerify, listOfLinksToVerify[k].ContainerA)
			}

			if el.findDimensionsInList(listOfLinksToVerify[k].ContainerB, containerTested) == false {
				containersToVerify = append(containersToVerify, listOfLinksToVerify[k].ContainerB)
			}
		}

		if len(containersToVerify) == 0 {
			break
		}

		pass := false
		for k := range containersToVerify {
			container = containersToVerify[k]
			if el.findDimensionsInList(container, containerTested) == false {
				pass = true
				break
			}
		}

		if pass == false {
			break
		}

		containerTested = append(containerTested, container)
		newLinkList := el.LinkVerticalFilterEachContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainer(container, listLinks)

		for k := range newLinkList {
			if el.findLinkInList(newLinkList[k], listOfLinksToVerify) == false {
				listOfLinksToVerify = append(listOfLinksToVerify, newLinkList[k])
			}
		}
	}

	return listOfLinksToVerify
}

func (el Filter) LinkVerticalFilterContainersWithErrorAndFindContainersWithIsNotAlignsFromTheTopOrTheBottomInRelationToFatherDirectlyOrIndirectlyLinked(listOfContainers []*Dimensions, listOfLinks []*Link) (error, []*Dimensions) {

	fatherContainer := el.SearchContainerFather(listOfLinks)
	if fatherContainer == nil {
		return errors.New("father containers not found in containers list"), []*Dimensions{}
	}
	listOfLinksToVerify := el.LinkVerticalFilterContainerIsAlignsFromTheTopOrTheBottomInRelationToAnotherContainerDirectlyOrIndirectlyLinked(fatherContainer, listOfLinks)
	return nil, el.SearchForFloatingContainers(listOfLinksToVerify, listOfContainers)
}

//todo: filters
//  +-father------------------------------------------------------------+
//  |                                                                   |
//  |  +-containerA--X-------------+     +-containerB--X-------------+  |
//  |  |                           |     |                           |  |
//  |  |                           |     |                           |  |
//  |  |                           |     |                           |  |
//  |  +-------------O-------------+     +-------------O-------------+  |
//  |                |                                 |                |
//  |                +---------------------------------+                |
//  |                                                                   |
//  +-------------------------------------------------------------------+
//
//  +-father------------------------------------------------------------+
//  |                                                                   |
//  |                +---------------------------------+                |
//  |                |                                 |                |
//  |  +-containerA--O-------------+     +-containerB--O-------------+  |
//  |  |                           |     |                           |  |
//  |  |                           |     |                           |  |
//  |  |                           |     |                           |  |
//  |  +-------------X-------------+     +-------------X-------------+  |
//  |                                                                   |
//  +-------------------------------------------------------------------+
//
//  +-father------------------------------------------------------------+
//  |                                                                   |
//  |                +---------------------------------+                |
//  |                |                                 |                |
//  |  +-containerA--O-------------+     +-containerB--O-------------+  |
//  |  |                           |     |                           |  |
//  |  |                           |     |                           |  |
//  |  |                           |     |                           |  |
//  |  +-------------O-------------+     +-------------O-------------+  |
//  |                |                                 |                |
//  |                +---------------------------------+                |
//  |                                                                   |
//  +-------------------------------------------------------------------+
//
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O                           O--+  |
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  |  |                                 |  |
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O                           O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |     X                           O--+  |
//  |     |                           |  |  |
//  |     +---------------------------+  |  |
//  |                                    |  |
//  |     +-containerB----------------+  |  |
//  |     |                           |  |  |
//  |     X                           O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O                           X     |
//  |  |  |                           |     |
//  |  |  +---------------------------+     |
//  |  |                                    |
//  |  |  +-containerB----------------+     |
//  |  |  |                           |     |
//  |  +--O                           X     |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
