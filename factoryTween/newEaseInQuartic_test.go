package factoryTween

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.platform/fps"
	"github.com/helmutkemper/util"
	"sync"
	"time"
)

func ExampleTween_NewEaseInQuartic_1() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	fps.Set(3)

	start := time.Now()

	NewEaseInQuartic(
		time.Second*2,
		0,
		3,
		func(value float64) {
			fmt.Printf("value on interaction event: %v\n", util.Round(value, 0.5, 1.0))
		},
		func(value float64) {
			wg.Done()
			fmt.Printf("value on done event: %v\n", util.Round(value, 0.5, 1.0))
		},
	)

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("total time: %v", util.Round(elapsed.Seconds(), 0.5, 1.0))

	//output:
	//value on interaction event: 0
	//value on interaction event: 0
	//value on interaction event: 0.2
	//value on interaction event: 0.6
	//value on interaction event: 1.4
	//value on interaction event: 3
	//value on done event: 3
	//total time: 2
}

func ExampleTween_NewEaseInQuartic_2() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	fps.Set(3)

	start := time.Now()

	NewEaseInQuartic(
		time.Second*2,
		3,
		0,
		func(value float64) {
			fmt.Printf("value on interaction event: %v\n", util.Round(value, 0.5, 1.0))
		},
		func(value float64) {
			wg.Done()
			fmt.Printf("value on done event: %v\n", util.Round(value, 0.5, 1.0))
		},
	)

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("total time: %v", util.Round(elapsed.Seconds(), 0.5, 1.0))

	//output:
	//value on interaction event: 3
	//value on interaction event: 3
	//value on interaction event: 2.8
	//value on interaction event: 2.4
	//value on interaction event: 1.6
	//value on interaction event: -0.1
	//value on done event: -0.1
	//total time: 2
}
