// Package dimensions
//
// pt_br:
// A base do sistema gráfico diz que todo container fica contido dentro de um
// container pai. Um container pai pode ser o container principal da aplicação ou un
// container qualquer que recebe outro container, por exemplo, um botão
//
//
//  +-main-application-----O----------------------+
//  |                      |                      |
//  |     +-containerA-----O-------children-+     |
//  |     |                                 |     |
//  O-----O                                 O-----O
//  |     |                                 |     |
//  |     +----------------O----------------+     |
//  |                      |                      |
//  |     +-buttonA--------O-------children-+     |
//  |     |                                 |     |
//  O-----O                                 O-----O
//  |     |                                 |     |
//  |     +----------------O----------------+     |
//  |                      |                      |
//  +----------------------O-----father-container-+
//
//  +-buttonA--O---------------------O------------+
//  |          |                     |            |
//  |   +-icon-O------+     +-text---O--------+   |
//  |   |             |     |                 |   |
//  O---O  children   O-----O     children    O---O
//  |   |             |     |                 |   |
//  |   +------O------+     +--------O--------+   |
//  |          |                     |            |
//  +----------O-----------------father-container-+
//
package dimensions
