package factoryTween

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.platform/fps"
	"github.com/helmutkemper/util"
	"sync"
	"time"
)

func ExampleTween_NewEaseInOutQuadratic_1() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	fps.Set(3)

	start := time.Now()

	NewEaseInOutQuadratic(
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
	//value on interaction event: 0.2
	//value on interaction event: 0.7
	//value on interaction event: 1.5
	//value on interaction event: 2.3
	//value on interaction event: 2.8
	//value on interaction event: 3
	//value on done event: 3
	//total time: 2
}

func ExampleTween_NewEaseInOutQuadratic_2() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	fps.Set(3)

	start := time.Now()

	NewEaseInOutQuadratic(
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
	//value on interaction event: 2.8
	//value on interaction event: 2.3
	//value on interaction event: 1.5
	//value on interaction event: 0.7
	//value on interaction event: 0.2
	//value on interaction event: 0
	//value on done event: 0
	//total time: 2
}
