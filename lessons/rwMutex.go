package lessons

import (
	"fmt"
	"sync"
	"time"
)

var likes = 0
var rwMtx = sync.RWMutex{}

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for range 100_000 {
		rwMtx.Lock() // usually, when values set somewhere goroutines, might happen nonfull writing these values, these values might be incorrect. when we will be reading these values, that we will be to get incorrect values, unpredictable values. blocking guaranteed gives correct values.
		likes++
		rwMtx.Unlock()
	}
}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for range 100_000 {
		rwMtx.RLock() // if setting value is unlock, that parallel reading values is available. goroutines reads values parallel. program works is quicker.
		_ = likes
		rwMtx.RUnlock()
	}
}

func exampleWithLikes() {
	wg := &sync.WaitGroup{}
	initTime := time.Now()

	for range 10 {
		wg.Add(1)
		go setLike(wg)
	}

	for range 10 {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait()
	fmt.Println("time is over: ", time.Since(initTime))
}

func RwMutexExample() {
	exampleWithLikes()
}
