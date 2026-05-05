package lessons

import (
	"researching-go/pkg/logger"
	"sync"
	"sync/atomic"
)

// var number int = 0
var number atomic.Int64 // atomic - blocks race condition, because it's make operations this variable simple. if we will not to use atomic, that we got various values while same code executing

func increase(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		number.Add(1) // struct Int64 has methods, e.g. - add value
	}
}

func raceConditionConception1() {
	wg := &sync.WaitGroup{}

	wg.Add(10)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	logger.Ptc("total value: ", number.Load()) // read value atomic-type variable
}

var slice []int
var mtx sync.Mutex // mutex blocks completing codes after Lock() and before Unlock(), only one goroutine can to work this code, others is waiting

func addSliceValue(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		mtx.Lock()
		slice = append(
			slice,
			i,
		)
		mtx.Unlock()
	}

}

func raceConditionConception2() {
	wg := &sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		go addSliceValue(wg)
	}

	wg.Wait()

	logger.Ptc("slice: ", len(slice))
}

func RaceConditionAMExample() {
	//raceConditionConception1()
	raceConditionConception2()
}
