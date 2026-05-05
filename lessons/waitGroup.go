package lessons

import (
	"researching-go/pkg/logger"
	"sync"
)

func postman(text string, wg *sync.WaitGroup) {
	defer wg.Done() // this function adding e.g. in the end goroutine, it's decrement counter waitGroup by 1.

	for i := 1; i <= 3; i++ {
		logger.Ptc("postman bring the newspaper", text, " in ", i, " times")
		//time.Sleep(200 * time.Millisecond)
	}
}

func postmanExample() {
	wg := &sync.WaitGroup{} // implement waitGroup should be by pointer

	wg.Add(1) // we add in waitGroup before call each goroutine number of 1 (e.g.). counter of calling goroutine updated. e.g., we have 3 goroutines, that means final counter waitGroup - 3.
	go postman("News", wg)

	wg.Add(1)
	go postman("Letter", wg)

	wg.Add(1)
	go postman("Cars magazine", wg)

	wg.Wait() // Wait function is waiting while every goroutine in waitGroup will be done. when counter will be 0 - program is completed.
}

func WaitGroupExample() {
	postmanExample()
}
