package factoryTween

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/tween"
	"time"
)

// NewEaseInQuadratic
//
// English: Ease tween in quadratic
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStartFunc: on start function
//     onEndFunc: on end function
//     onCycleStartFunc: on tween cycle start function
//     onCycleEndFunc: on tween cycle end function
//     onInvertFunc: on invert function. Only for loop != 0
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//                Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: x; args[1]: y}
//
// Português: Facilitador de interpolação in quadratic
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStartFunc: função do evento início
//     onEndFunc: função do evento fim
//     onCycleStartFunc: função para o início do ciclo de interpolação
//     onCycleEndFunc: função para o fim do ciclo de interpolação
//     onInvertFunc: função para o evento inversão. Apenas se loop != 0
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//                Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: x; args[1]: y}
func NewEaseInQuadratic(
	duration time.Duration,
	startValue,
	endValue float64,
	onStartFunc,
	onEndFunc,
	onCycleStartFunc,
	onCycleEndFunc,
	onInvertFunc func(value float64, arguments ...interface{}),
	onStepFunc func(value, percentToComplete float64, arguments ...interface{}),
	loop int,
	arguments ...interface{},
) *tween.Tween {

	t := &tween.Tween{
		Engine:       global.Global.Engine,
		OnStart:      onStartFunc,
		OnEnd:        onEndFunc,
		OnCycleStart: onCycleStartFunc,
		OnCycleEnd:   onCycleEndFunc,
		Arguments:    arguments,
		Duration:     duration,
		StartValue:   startValue,
		EndValue:     endValue,
		Interaction:  onStepFunc,
		OnInvert:     onInvertFunc,
		Repeat:       loop,
		Func:         tween.KEaseInQuadratic,
	}
	t.Start()

	return t
}
