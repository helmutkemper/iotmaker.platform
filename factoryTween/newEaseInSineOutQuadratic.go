package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStartFunc: on start function
//     onEndFunc: on end function
//     onInvertFunc: on invert function. Only for loop != 0
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//                Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: x; args[1]: y}
//
// pt_br:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStartFunc: função do evento início
//     onEndFunc: função do evento fim
//     onInvertFunc: função para o evento inverção. Apenas se loop != 0
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//                Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: x; args[1]: y}
func NewEaseInSineOutQuadratic(duration time.Duration, startValue, endValue float64, onStartFunc, onEndFunc, onInvertFunc func(value float64, arguments ...interface{}), onStepFunc func(value, percentToComplete float64, arguments ...interface{}), loop int, arguments ...interface{}) *tween.Tween {
	t := &tween.Tween{
		OnStart:     onStartFunc,
		OnEnd:       onEndFunc,
		Arguments:   arguments,
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInSineOutQuadratic,
		Interaction: onStepFunc,
		OnInvert:    onInvertFunc,
		Repeat:      loop,
	}
	t.Start()

	return t
}
