package dimensions

type WallTop int

const (
	// en: Wall top not set
	//  +-father------------X-------------------+
	//  |                                       |
	//  |     +-containerB--X-------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +---------------------------+     |
	//  |                                       |
	//  +---------------------------------------+
	//
	//  +-father--------------------------------+
	//  |                                       |
	//  |     +-containerB--X-------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +---------------------------+     |
	//  |                                       |
	//  +-------------------X-------------------+
	//
	//  +-father--------------------------------+
	//  |                                       |
	//  |     +-containerA----------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +-------------X-------------+     |
	//  |                                       |
	//  |     +-containerB--X-------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +---------------------------+     |
	//  |                                       |
	//  +---------------------------------------+
	KWallTopNotSet WallTop = iota

	// en: Wall top container B link on top to container A link on top
	//  +-father------------O-------------------+
	//  |                   |                   |
	//  |     +-containerB--O-------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +---------------------------+     |
	//  |                                       |
	//  +---------------------------------------+
	//
	//  +-father--------------------------------+
	//  |                                       |
	//  |     +-containerA--O-------------+     |
	//  |     |             |             |     |
	//  |     |             |             |     |
	//  |     |             |             |     |
	//  |     +-------------+-------------+     |
	//  |                   |                   |
	//  |     +-containerB--O-------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +---------------------------+     |
	//  |                                       |
	//  +---------------------------------------+
	KWallTopContainerBLinkOnTopToContainerALinkOnTop

	// en: Wall top from container B link on top to container A link on bottom
	//  +-father--------------------------------+
	//  |                                       |
	//  |     +-containerB--O-------------+     |
	//  |     |             |             |     |
	//  |     |             |             |     |
	//  |     |             |             |     |
	//  |     +-------------+-------------+     |
	//  |                   |                   |
	//  +-------------------O-------------------+
	//
	//  +-father--------------------------------+
	//  |                                       |
	//  |     +-containerA----------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +-------------O-------------+     |
	//  |                   |                   |
	//  |     +-containerB--O-------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +---------------------------+     |
	//  |                                       |
	//  +---------------------------------------+
	KWallTopContainerBLinkOnTopToContainerALinkOnBottom
)

type WallBottom int

const (
	// en: Wall bottom not set
	//
	//   +-father--------------------------------+
	//   |                                       |
	//   |     +-containerB----------------+     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     +-------------X-------------+     |
	//   |                                       |
	//   +-------------------X-------------------+
	//
	//   +-father--------------------------------+
	//   |                                       |
	//   |     +-containerA----------------+     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     +-------------X-------------+     |
	//   |                                       |
	//   |     +-containerB----------------+     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     +-------------X-------------+     |
	//   |                                       |
	//   +---------------------------------------+
	KWallBottomNotSet WallBottom = iota

	// en: Wall bottom from container B link on top to container A link on top
	//
	//   +-father------------O-------------------+
	//   |                   |                   |
	//   |     +-containerB--+-------------+     |
	//   |     |             |             |     |
	//   |     |             |             |     |
	//   |     |             |             |     |
	//   |     +-------------O-------------+     |
	//   |                                       |
	//   +---------------------------------------+
	//
	//   +-father--------------------------------+
	//   |                                       |
	//   |     +-containerA--O-------------+     |
	//   |     |             |             |     |
	//   |     |             |             |     |
	//   |     |             |             |     |
	//   |     +-------------+-------------+     |
	//   |                   |                   |
	//   |     +-containerB--+-------------+     |
	//   |     |             |             |     |
	//   |     |             |             |     |
	//   |     |             |             |     |
	//   |     +-------------O-------------+     |
	//   |                                       |
	//   +---------------------------------------+
	KWallBottomContainerBLinkOnTopToContainerALinkOnTop

	// en: Wall bottom from container B link on top to container A link on bottom
	//
	//   +-father--------------------------------+
	//   |                                       |
	//   |     +-containerB----------------+     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     +-------------O-------------+     |
	//   |                   |                   |
	//   +-------------------O-------------------+
	//
	//   +-father--------------------------------+
	//   |                                       |
	//   |     +-containerA----------------+     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     +---------------------------+     |
	//   |                                       |
	//   |     +-containerB----------------+     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     |                           |     |
	//   |     +---------------------------+     |
	//   |                                       |
	//   +---------------------------------------+
	KWallBottomContainerBLinkOnTopToContainerALinkOnBottom

	// en: Wall bottom from container B link on bottom to container A link on bottom
	//  +-father--------------------------------+
	//  |                                       |
	//  |     +-containerB----------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +-------------O-------------+     |
	//  |                   |                   |
	//  +-------------------O-------------------+
	//
	//  +-father--------------------------------+
	//  |                                       |
	//  |     +-containerA----------------+     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     |                           |     |
	//  |     +-------------O-------------+     |
	//  |                   |                   |
	//  |     +-containerB--+-------------+     |
	//  |     |             |             |     |
	//  |     |             |             |     |
	//  |     |             |             |     |
	//  |     +-------------O-------------+     |
	//  |                                       |
	//  +---------------------------------------+
	KWallBottomContainerBLinkOnBottomToContainerALinkOnBottom
)

type WallLeft int

const (
	// en: Wall left not set
	//  +-father----------------------------+
	//  |                                   |
	//  |   +-containerB----------------+   |
	//  |   |                           |   |
	//  X   X                           |   |
	//  |   |                           |   |
	//  |   +---------------------------+   |
	//  |                                   |
	//  +-----------------------------------+
	//
	//  +-father------------------------------------------------------------+
	//  |                                                                   |
	//  |   +-containerA----------------+   +-containerB----------------+   |
	//  |   |                           |   |                           |   |
	//  |   |                           X   X                           |   |
	//  |   |                           |   |                           |   |
	//  |   +---------------------------+   +---------------------------+   |
	//  |                                                                   |
	//  +-------------------------------------------------------------------+
	KWallLeftNotSet WallLeft = iota

	// en: Wall left from container B link on left to container A link on right
	//  +-father------------------------------------------------------------+
	//  |                                                                   |
	//  |   +-containerA----------------+   +-containerB----------------+   |
	//  |   |                           |   |                           |   |
	//  |   |                           O---O                           |   |
	//  |   |                           |   |                           |   |
	//  |   +---------------------------+   +---------------------------+   |
	//  |                                                                   |
	//  +-------------------------------------------------------------------+
	KWallLeftContainerBLinkOnLeftToContainerALinkOnRight

	// en: Wall left from container B link on left to container A link on left
	//  +-father----------------------------+
	//  |                                   |
	//  |   +-containerB----------------+   |
	//  |   |                           |   |
	//  O---O                           |   |
	//  |   |                           |   |
	//  |   +---------------------------+   |
	//  |                                   |
	//  +-----------------------------------+
	//
	//  +-father------------------------------------------------------------+
	//  |                                                                   |
	//  |   +-containerA----------------+   +-containerB----------------+   |
	//  |   |                           |   |                           |   |
	//  |   O---------------------------+---O                           |   |
	//  |   |                           |   |                           |   |
	//  |   +---------------------------+   +---------------------------+   |
	//  |                                                                   |
	//  +-------------------------------------------------------------------+
	KWallLeftContainerBLinkOnLeftToContainerALinkOnLeft
)

type WallRight int

const (
	// en: Wall right not set
	//  +-father------------------------------------------------------------+
	//  |                                                                   |
	//  |   +-containerA----------------+   +-containerB----------------+   |
	//  |   |                           |   |                           |   |
	//  |   |                           X   |                           X   |
	//  |   |                           |   |                           |   |
	//  |   +---------------------------+   +---------------------------+   |
	//  |                                                                   |
	//  +-------------------------------------------------------------------+
	//
	//  +-father----------------------------+
	//  |                                   |
	//  |   +-containerB----------------+   |
	//  |   |                           |   |
	//  |   |                           X   X
	//  |   |                           |   |
	//  |   +---------------------------+   |
	//  |                                   |
	//  +-----------------------------------+
	KWallRightNotSet WallRight = iota

	// en: Wall right from container B lin on right to container A link on right
	//  +-father------------------------------------------------------------+
	//  |                                                                   |
	//  |   +-containerA----------------+   +-containerB----------------+   |
	//  |   |                           |   |                           |   |
	//  |   |                           O---+---------------------------O   |
	//  |   |                           |   |                           |   |
	//  |   +---------------------------+   +---------------------------+   |
	//  |                                                                   |
	//  +-------------------------------------------------------------------+
	//
	//  +-father----------------------------+
	//  |                                   |
	//  |   +-containerB----------------+   |
	//  |   |                           |   |
	//  |   |                           O---O
	//  |   |                           |   |
	//  |   +---------------------------+   |
	//  |                                   |
	//  +-----------------------------------+
	KWallRightContainerBLinkOnRightToContainerALinkOnRight
)

type Link struct {
	ToDebug          string
	ContainerA       *Dimensions
	ContainerAFather bool
	ContainerB       *Dimensions
	Top              WallTop
	Left             WallLeft
	Right            WallRight
	Bottom           WallBottom
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
func (el *Dimensions) LinkIsPresentAtTheBottom(containerA, containerB *Dimensions) {
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
func (el *Dimensions) LinkIsPresentAtTheTop(containerA, containerB *Dimensions) {
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
func (el *Dimensions) LinkIsPresentAtTheTopAndBottom(containerA, containerB *Dimensions) {
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
func (el *Dimensions) LinkIsPresentAtTheRight(containerA, containerB *Dimensions) {
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
func (el *Dimensions) LinkIsPresentAtTheLeft(containerA, containerB *Dimensions) {
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
func (el *Dimensions) LinkIsPresentAtTheRightAndLeft(containerA, containerB *Dimensions) {
	AX := containerA.X
	AWidth := containerA.Width
	CWidth := containerB.Width
	containerB.X = AX + AWidth/2 - CWidth/2
}

func (el *Link) LinkAssembly(father Link, list []Link) error {

	return nil
}

/*func (el *Link) SimpleLinkAssembly(father *Dimensions, container LinkCollection) error {

	//
	// +-------------O------------+
	// |             |topIsSet    |
	// | +-----------O----------+ |
	// O-O leftIsSet            O-O rightIsSet
	// | +-----------O----------+ |
	// |             |bottomIsSet |
	// +-------------O------------+
	topIsSet := !reflect.DeepEqual(*container.Top, Link{})
	leftIsSet := !reflect.DeepEqual(*container.Left, Link{})
	rightIsSet := !reflect.DeepEqual(*container.Right, Link{})
	bottomIsSet := !reflect.DeepEqual(*container.Bottom, Link{})

	if topIsSet == false && bottomIsSet == false {
		return errors.New("this container can't be assembled. links a and c are not connected")
	} else if topIsSet == false && bottomIsSet == true {
		father.LinkIsPresentAtTheBottom(father, container.Bottom.Container)
	} else if topIsSet == true && bottomIsSet == false {
		father.LinkIsPresentAtTheTop(father, container.Top.Container)
	} else if topIsSet == true && bottomIsSet == true {
		father.LinkIsPresentAtTheTopAndBottom(father, container.Bottom.Container)
	}

	if leftIsSet == false && rightIsSet == false {
		return errors.New("this container can't be assembled. links b and d are not connected")
	} else if leftIsSet == false && rightIsSet == true {
		father.LinkIsPresentAtTheRight(father, container.Right.Container)
	} else if leftIsSet == true && rightIsSet == false {
		father.LinkIsPresentAtTheLeft(father, container.Left.Container)
	} else if leftIsSet == true && rightIsSet == true {
		father.LinkIsPresentAtTheRightAndLeft(father, container.Right.Container)
	}

	return nil
}*/

//                Top
//      +----------O---------+
//      |                    |
//      |   +------O-----+   |
// Left O   O            O   O Right
//      |   +------O-----+   |
//      |                    |
//      +----------O---------+
//               Bottom
//
//                      A
//             +--------O-------+
//             |                |
//           D O                O B
//             |                |
//             +--------O-------+
//                      C
//
//  Molde:
//  +-father------------O-------------------+
//  |                   |                   |
//  |     +-containerA--O-------------+     |
//  |     |                           |     |
//  O-----O                           O-----O
//  |     |                           |     |
//  |     +-------------O-------------+     |
//  |                   |                   |
//  |     +-containerB--O-------------+     |
//  |     |                           |     |
//  O-----O                           O-----O
//  |     |                           |     |
//  |     +-------------O-------------+     |
//  |                   |                   |
//  +-------------------O-------------------+
//
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
//
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
//
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
//
//  Ponto de vista horizontal
//  Option: D - os containers a e b têm o seu tamanho ajustados pelo maior
//              o container a é centralizado pelo pai
//              o container b y é ajustado por a y
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O--+--O a link to father          O--+--O
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  |  |                                 |  |
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O b link to a               O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: E - o container a é ajustado a direita pelo pai
//              o container b y+w é ajustado por a y+w
//              w não é alterado nem em a nem em b
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |     | a link to father          O--+--O
//  |     |                           |  |  |
//  |     +---------------------------+  |  |
//  |                                    |  |
//  |     +-containerB----------------+  |  |
//  |     |                           |  |  |
//  |     | b link to a               O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: F - o container a é ajustado a direita pelo pai
//              o container b y+w é ajustado por a y+w
//              w é alterado pelo maior
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O a link to father          O--+--O
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  |  |                                 |  |
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O b link to a               O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: G
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O--+--O a link to father          |     |
//  |  |  |                           |     |
//  |  |  +---------------------------+     |
//  |  |                                    |
//  |  |  +-containerB----------------+     |
//  |  |  |                           |     |
//  |  +--O b link to a               |     |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: H
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O--+--O a link to father          O--+  |
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  |  |                                 |  |
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O b link to a               O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: I
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O                           O-----O
//  |  |  |                           |     |
//  |  |  +---------------------------+     |
//  |  |                                    |
//  |  |  +-containerB----------------+     |
//  |  |  |                           |     |
//  |  +--O                           O-----O
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Option: J
//  Correto:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  O-----O                           O--+  |
//  |     |                           |  |  |
//  |     +---------------------------+  |  |
//  |                                    |  |
//  |     +-containerB----------------+  |  |
//  |     |                           |  |  |
//  O-----O                           O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Errado:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O                           O--+  |
//  |  |  |                           |  |  |
//  |  |  +---------------------------+  |  |
//  X  |                                 |  X
//  |  |  +-containerB----------------+  |  |
//  |  |  |                           |  |  |
//  |  +--O                           O--+  |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  Ponto de vista horizontal
//  Errado:
//  +-father--------------------------------+
//  |                                       |
//  |     +-containerA----------------+     |
//  |     |                           |     |
//  |  +--O                           O-----O
//  |  |  |                           |     |
//  |  |  +---------------------------+     |
//  X  |                                    |
//  |  |  +-containerB----------------+     |
//  |  |  |                           |     |
//  |  +--O                           X     |
//  |     |                           |     |
//  |     +---------------------------+     |
//  |                                       |
//  +---------------------------------------+
//
//  +-father--------------------------------O---------------------------------------+
//  |                                       |                                       |
//  |                   +-------------------+-------------------+                   |
//  |                   |                                       |                   |
//  |     +-containerA--O-------------+           +-containerB--O-------------+     |
//  |     |                           |           |                           |     |
//  |  +--O                           O-----------O                           O--+  |
//  |  |  |                           |           |                           |  |  |
//  |  |  +-------------O-------------+           +-------------O-------------+  |  |
//  O--+                |                                       |                +--O
//  |  |  +-containerC--O-------------+           +-containerD--O-------------+  |  |
//  |  |  |                           |           |                           |  |  |
//  |  +--O                           O-----------O                           O--+  |
//  |     |                           |           |                           |     |
//  |     +-------------O-------------+           +-------------O-------------+     |
//  |                   |                                       |                   |
//  |                   +-------------------+-------------------+                   |
//  |                                       |                                       |
//  +---------------------------------------O---------------------------------------+
//
type Dimensions struct {
	ToDebug string
	Level   int

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

func NewContainer(debug string, width, height int) *Dimensions {
	return &Dimensions{
		ToDebug: debug,

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

func NewContainerWidthXY(debug string, x, y, width, height int) *Dimensions {
	return &Dimensions{
		ToDebug: debug,

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

func NewLink(debug string, containerA, containerB *Dimensions, top WallTop, left WallLeft, right WallRight, bottom WallBottom) *Link {
	return &Link{
		ToDebug:          debug,
		ContainerA:       containerA,
		ContainerAFather: false,
		ContainerB:       containerB,
		Top:              top,
		Left:             left,
		Right:            right,
		Bottom:           bottom,
	}
}

func NewLinkWithFather(debug string, father, containerB *Dimensions, top WallTop, left WallLeft, right WallRight, bottom WallBottom) *Link {
	return &Link{
		ToDebug:          debug,
		ContainerA:       father,
		ContainerAFather: true,
		ContainerB:       containerB,
		Top:              top,
		Left:             left,
		Right:            right,
		Bottom:           bottom,
	}
}
