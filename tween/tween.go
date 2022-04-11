package tween

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"time"
)

type Tween struct {
	engine             engine.IEngine
	StartValue         float64
	endValue           float64
	arguments          []interface{}
	startTime          time.Time
	duration           time.Duration
	tweenFunc          func(currentTime, duration, startValue, endValue, changeInValue float64) float64
	interaction        func(value, percentToComplete float64, arguments ...interface{})
	onCycleStart       func(value float64, arguments ...interface{})
	onCycleEnd         func(value float64, arguments ...interface{})
	onStart            func(value float64, arguments ...interface{})
	onEnd              func(value float64, arguments ...interface{})
	onInvert           func(value float64, arguments ...interface{})
	doNotReverseMotion bool
	invert             bool
	repeat             int
	fpsUId             string
	loopStartValue     float64
	loopEndValue       float64
}

// SetEngine
//
// English:
//
//  Defines a new engine for time control.
//
//   Input:
//     value: object compatible with the engine.IEEngine interface
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Define uma nova engine para controle de tempo.
//
//   Entrada:
//     value: objeto compatível com ã interface engine.IEngine
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetEngine(value engine.IEngine) (object *Tween) {
	el.engine = value
	return el
}

// SetTweenFunc
//
// English:
//
//  Defines the tween math function to control the loop of interactions
//
//   Input:
//     value: tween math function.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Define a função matemática tween para controle do ciclo de interações
//
//   Entrada:
//     value: função matemática tween.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetTweenFunc(value func(currentTime, duration, startValue, endValue, changeInValue float64) float64) (object *Tween) {
	el.tweenFunc = value
	return el
}

// SetValues
//
// English:
//
//  Defines the initial and final values of the interactions cycle.
//
//   Input:
//     start: initial value for the beginning of the cycle of interactions;
//     end:   final value for the end of the iteration cycle.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Defines os valores inicial e final do ciclo de interações.
//
//   Entrada:
//     start: valor inicial para o início do ciclo de interações;
//     end:   valor final para o fim do ciclo de interações.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetValues(start, end float64) (object *Tween) {
	el.StartValue = start
	el.endValue = end
	return el
}

// SetDuration
//
// English:
//
//  Defines the total cycle time of interactions.
//
//   Input:
//     value: time.Duration contendo o tempo do ciclo de interações.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Define o tempo total do ciclo de interações.
//
//   Entrada:
//     value: time.Duration contendo o tempo do ciclo de interações.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetDuration(value time.Duration) (object *Tween) {
	el.duration = value
	return el
}

// SetDoNotReverseMotion
//
// English:
//
// Português:
//
//  Define a opção de reversão do movimento.
//
//   Entrada:
//     value: true para não reverter o movimento.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetDoNotReverseMotion(value bool) (object *Tween) {
	el.doNotReverseMotion = value
	return el
}

// SetLoops
//
// Português:
//
//  Define a quantidade de laços antes do fim da função.
//
//   Notas:
//     * A cada nova interação do laço ocorrerá uma inversão de movimento, a não ser que seja usada a
//       função SetDoNotReverseMotion(true);
//     * Para laços infinitos, defina o valor como sendo -1;
//     * Em caso de laço, a ordem das funções de eventos são: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnEndFunc() e SetOnInvertFunc()
//
//
//2022/04/10 22:38:39 start function
//wasm_exec.js:51 2022/04/10 22:38:39 start ciclo
//wasm_exec.js:51 2022/04/10 22:38:42 end ciclo
//wasm_exec.js:51 2022/04/10 22:38:42 end function
//wasm_exec.js:51 2022/04/10 22:38:42 inverte
//wasm_exec.js:51 2022/04/10 22:38:42 start ciclo
//wasm_exec.js:51 2022/04/10 22:38:44 end ciclo
//wasm_exec.js:51 2022/04/10 22:38:44 end function
//wasm_exec.js:51 2022/04/10 22:38:44 inverte
//wasm_exec.js:51 2022/04/10 22:38:44 start ciclo
//wasm_exec.js:51 2022/04/10 22:38:47 end ciclo
//wasm_exec.js:51 2022/04/10 22:38:47 end function
//wasm_exec.js:51 2022/04/10 22:38:47 inverte
//wasm_exec.js:51 2022/04/10 22:38:47 start ciclo
//wasm_exec.js:51 2022/04/10 22:38:49 end ciclo
//wasm_exec.js:51 2022/04/10 22:38:49 end function
//wasm_exec.js:51 2022/04/10 22:38:49 inverte
//wasm_exec.js:51 2022/04/10 22:38:49 start ciclo
//wasm_exec.js:51 2022/04/10 22:38:51 end ciclo
//wasm_exec.js:51 2022/04/10 22:38:51 end function
//wasm_exec.js:51 2022/04/10 22:38:51 inverte
//wasm_exec.js:51 2022/04/10 22:38:51 start ciclo
//wasm_exec.js:51 2022/04/10 22:38:54 end ciclo
//wasm_exec.js:51 2022/04/10 22:38:54 end function
//wasm_exec.js:51 2022/04/10 22:38:54 inverte
//wasm_exec.js:51 2022/04/10 22:38:54 start ciclo
//wasm_exec.js:51 2022/04/10 22:38:56 end ciclo
//wasm_exec.js:51 2022/04/10 22:38:56 end function
//wasm_exec.js:51 2022/04/10 22:38:56 inverte
//
//
func (el *Tween) SetLoops(value int) (object *Tween) {
	el.repeat = value
	return el
}

// SetOnStartFunc
//
// English:
//
//  Add the function to be called when the animation starts.
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: initial value defined in startValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada quando a animação inicia.
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor inicial definido em startValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnStartFunc(function func(value float64, arguments ...interface{})) (object *Tween) {
	el.onStart = function
	return el
}

// SetOnEndFunc
//
// English:
//
//  Add the function to be called when the animation ends.
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: final value defined in endValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada quando a animação inicia.
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor final definido em endValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnEndFunc(function func(value float64, arguments ...interface{})) (object *Tween) {
	el.onEnd = function
	return el
}

// SetOnCycleStartFunc
//
// English:
//
//  Adds the function to be called at the beginning of the interpolation cycle
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: initial value defined in startValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada no início do ciclo de interpolação
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor inicial definido em startValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnCycleStartFunc(function func(value float64, arguments ...interface{})) (object *Tween) {
	el.onCycleStart = function
	return el
}

// SetOnCycleEndFunc
//
// English:
//
//  Adds the function to be called at the ending of the interpolation cycle
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: final value defined in endValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada no fim do ciclo de interpolação
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor final definido em endValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnCycleEndFunc(function func(value float64, arguments ...interface{})) (object *Tween) {
	el.onCycleEnd = function
	return el
}

// SetOnStepFunc
//
// English:
//
//  Adds the function to be called for each iteration.
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: current value
//       percentToComplete: value between 0.0 and 1.0 indicating the percentage of the process
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada a cada interação
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor corrente
//       percentToComplete: valor entre 0.0 e 1.0 indicando o percentual do processo
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnStepFunc(function func(value, percentToComplete float64, arguments ...interface{})) (object *Tween) {
	el.interaction = function
	return el
}

// SetOnInvertFunc
//
// English:
//
//  Adds the function to be called on inversion of the interpolation cycle
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: current value
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada a cada interação
//
//   Entrada:
//     function: func(value, percentToComplete float64, arguments ...interface{})
//       value: valor corrente
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnInvertFunc(function func(value float64, arguments ...interface{})) (object *Tween) {
	el.onInvert = function
	return el
}

// Start
//
// English:
//
//  Starts the interaction according to the chosen tween function.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Inicia a interação conforme a função tween escolhida.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) Start() (object *Tween) {

	if el.engine == nil {
		el.engine = global.Global.Engine
	}

	if el.tweenFunc == nil {
		//fixme:descomentar
		//el.tweenFunc = KLinear
	}

	el.startTime = time.Now()
	el.invert = true

	if el.tweenFunc == nil {
		return
	}

	if el.onStart != nil {
		el.onStart(el.StartValue, el.arguments)
	}

	el.tickerRunnerPrepare(el.StartValue, el.endValue)

	return el
}

func (el *Tween) tickerRunnerPrepare(startValue, endValue float64) {
	if el.onCycleStart != nil {
		el.onCycleStart(el.StartValue, el.arguments)
	}

	el.loopStartValue = startValue
	el.loopEndValue = endValue

	// fixme: _ é um index do array? tem uso aqui?
	el.fpsUId, _ = el.engine.MathAddToFunctions(el.tickerRunnerRun)
}

func (el *Tween) tickerRunnerRun() {
	elapsed := time.Since(el.startTime)
	value := el.tweenFunc(elapsed.Seconds(), el.duration.Seconds(), el.loopStartValue, el.loopEndValue, el.loopEndValue-el.loopStartValue)
	percent := elapsed.Seconds() / el.duration.Seconds()

	if el.interaction != nil {
		el.interaction(value, percent, el.arguments)
	}

	if elapsed >= el.duration {

		el.Stop()

		if el.repeat == 0 {
			if el.onEnd != nil {
				el.onEnd(value)
			}
		}

		if el.repeat != 0 {
			el.startTime = time.Now()

			if el.onInvert != nil {
				el.onInvert(value)
			}

			if el.doNotReverseMotion == true {
				el.tickerRunnerPrepare(el.StartValue, el.endValue)
			} else {
				if el.invert == true {
					el.tickerRunnerPrepare(el.endValue, el.StartValue)
				} else {
					el.tickerRunnerPrepare(el.StartValue, el.endValue)
				}
				el.invert = !el.invert
			}

			el.repeat -= 1
		}
	}
}

// End
//
// English:
//
//  Terminates all interactions of the chosen Tween function, without invoking the onCycleEnd and
//  onEnd functions.
//
//   Saída:
//     object: reference to the current Tween object.
//
// Português:
//
// Termina todas as interações da função Tween escolhida, sem invocar as funções onCycleEnd e onEnd.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) End() (object *Tween) {
	el.engine.MathDeleteFromFunctions(el.fpsUId)

	return el
}

// Stop
//
// English:
//
//  Ends all interactions of the chosen Tween function, interacting with the onCycleEnd and onEnd
//  functions, respectively, in that order, if they have been defined.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Termina todas as interações da função Tween escolhida, interagindo com as funções onCycleEnd e
//  onEnd, respectivamente nessa ordem, se elas tiverem sido definidas.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) Stop() (object *Tween) {
	el.engine.MathDeleteFromFunctions(el.fpsUId)
	if el.onCycleEnd != nil {
		el.onCycleEnd(el.endValue, el.arguments)
	}

	if el.onEnd != nil {
		el.onEnd(el.endValue, el.arguments)
	}

	return el
}
