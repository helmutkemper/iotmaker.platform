package dimensions

type Filter struct{}

func (el Filter) FilterContainerFather(list []*Link) []*Dimensions {
	ret := make([]*Dimensions, 0)
	for _, link := range list {
		if link.ContainerAFather == true {
			ret = append(ret, link.ContainerA)
		}
	}

	return ret
}

func (el Filter) LinkAssemblyCheckIfFatherExistsAndHasOnlyOne(list []*Link) bool {
	counter := 0
	for _, link := range list {
		if link.ContainerAFather == true {
			counter += 1
		}
	}

	return counter == 1
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
func (el Filter) LinkAssemblyHorizontalFilterContainerIsCentralizedInRelationToTheFather(list []*Link) []*Link {

	ret := make([]*Link, 0)
	for _, link := range list {
		fatherPass := link.ContainerAFather == true
		rightPass := link.Right == KCornerRightContainerBLinkOnRightToContainerALinkOnRight
		leftPass := link.Left == KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft

		pass := rightPass && leftPass && fatherPass
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
func (el Filter) LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheRightInRelationToTheFather(list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		fatherPass := link.ContainerAFather == true
		rightPass := link.Right == KCornerRightContainerBLinkOnRightToContainerALinkOnRight
		leftPass := link.Left == KCornerLeftNotSet

		pass := rightPass && leftPass && fatherPass
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
func (el Filter) LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheRightInRelationToAnother(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		rightPass := link.Right == KCornerRightContainerBLinkOnRightToContainerALinkOnRight
		leftPass := link.Left == KCornerLeftNotSet

		pass := rightPass && leftPass && (containerAPass || containerBPass)
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
func (el Filter) LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheRightAndTheLeftInRelationToAnother(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		rightPass := link.Right == KCornerRightContainerBLinkOnRightToContainerALinkOnRight
		leftPass := link.Left == KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft

		pass := rightPass && leftPass && (containerAPass || containerBPass)
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
func (el Filter) LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftInRelationToTheFather(list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		fatherPass := link.ContainerAFather == true
		rightPass := link.Right == KCornerRightNotSet
		leftPass := link.Left == KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft

		pass := rightPass && leftPass && fatherPass
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
func (el Filter) LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftInRelationToAnother(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		rightPass := link.Right == KCornerRightNotSet
		leftPass := link.Left == KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft

		pass := rightPass && leftPass && (containerAPass || containerBPass)
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
//  |  |  +-containerA----------------+  |  |  |  |  +-containerA----------------+  |  |
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

func (el Filter) LinkAssemblyErrorFilterContainerIsLinkedInRelationToItSelf(container *Dimensions, list []*Link) []*Link {
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

func (el Filter) LinkAssemblyHorizontalFilterContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainerDirectlyOrIndirectlyLinked(container *Dimensions, listLinks []*Link) []*Link {

	listToVerify := el.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainer(container, listLinks)
	containerTested := []*Dimensions{container}

	for {

		pass := false
		for _, link := range listToVerify {

			if el.findDimensionsInList(link.ContainerA, containerTested) == false {
				pass = true

				containerTested = append(containerTested, link.ContainerA)

				newList := el.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainer(link.ContainerA, listToVerify)

				for k := range newList {
					if el.findLinkInList(newList[k], listToVerify) == false {
						listToVerify = append(listToVerify, newList[k])
					}
				}
			}

			if el.findDimensionsInList(link.ContainerB, containerTested) == false {
				pass = true
				containerTested = append(containerTested, link.ContainerB)

				newList := el.LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainer(link.ContainerB, listToVerify)
				for k := range newList {
					if el.findLinkInList(newList[k], listToVerify) == false {
						listToVerify = append(listToVerify, newList[k])
					}
				}
			}
		}

		if pass == false {
			break
		}
	}

	return listToVerify
}

func (el Filter) LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftOrTheRightInRelationToAnotherContainer(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		rightPass := link.Right == KCornerRightContainerBLinkOnRightToContainerALinkOnRight
		leftPass := link.Left == KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft

		pass := (rightPass || leftPass) && (containerAPass || containerBPass)
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

func (el Filter) LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheLeftInRelationToAnotherContainer(container *Dimensions, list []*Link) []*Link {
	ret := make([]*Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		rightPass := link.Right == KCornerRightNotSet
		leftPass := link.Left == KCornerLeftContainerBLinkOnLeftToContainerALinkOnLeft

		pass := rightPass && leftPass && (containerAPass || containerBPass)
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}

func (el Filter) LinkAssemblyHorizontalFilterEachContainerIsAlignsFromTheRightInRelationToAnotherContainer(container *Dimensions, list []Link) []Link {
	ret := make([]Link, 0)

	for _, link := range list {
		containerAPass := link.ContainerA == container
		containerBPass := link.ContainerB == container
		rightPass := link.Right == KCornerRightContainerBLinkOnRightToContainerALinkOnRight
		leftPass := link.Left == KCornerLeftNotSet

		pass := rightPass && leftPass && (containerAPass || containerBPass)
		if pass == true {
			ret = append(ret, link)
		}
	}

	return ret
}
