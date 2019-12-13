package factoryFont

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/font"
	"image/color"
)

// en: Create a new font struct
//     size: size of the font
//     sizeUnit: size unit of the font
//     fontColor: a color.RGBA struct used with text
//     family: font name. Must be a string or a constant fontFamily.FontFamily
//             Example: 'Verdana' or fontFamily.KHelvetica
//     density: Please, see a density methods
//     iDensity: Please, see a density methods
//
// pt_br: Cria um novo struct de fonte
//     size: tamanho da fonte
//     sizeUnit: unidade do tamanho da fonte, exemplo: 'px'
//     fontColor: um struct color.RGBA usado com texto
//     family: nome da fonte escolhida. Pode ser string ou uma constante
//     fontFamily.FontFamily.
//             Exemplo: 'Verdana' ou fontFamily.KHelvetica
//     density: Por favor, veja o elemento densidade
//     iDensity: Por favor, veja o elemento densidade
func NewFont(size int, sizeUnit string, fontColor color.RGBA, family interface{}, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) font.Font {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(size)
	size = densityCalc.Int()

	f := font.Font{
		Size:     size,
		SizeUnit: sizeUnit,
		Family:   family,
		Color:    fontColor,
	}
	return f
}
