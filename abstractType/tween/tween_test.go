package tween

import (
	"fmt"
	"sync"
	"time"
)

func ExampleTween_Start() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	SetFPS(10)

	start := time.Now()

	t := Tween{
		Duration:   time.Second * 2,
		StartValue: 0,
		EndValue:   3,
		Func:       Linear,
		Interaction: func(value float64) {
			fmt.Printf("value on interaction event: %v\n", value)
		},
		Done: func(value float64) {
			wg.Done()
			fmt.Printf("value on done event: %v\n", value)
		},
	}
	t.Start()

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("total time: %v", elapsed)

	// output:
	//
}
