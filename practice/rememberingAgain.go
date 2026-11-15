package practice

import (
	"researching-go/pkg/logger"
	"sync"
	"time"
)

func worker(ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- "готово"
}

func startWorker() {
	ch := make(chan string)
	go worker(ch)

	res := <-ch
	logger.Ptc("result: ", res)
}

func simpleWgWait() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(2)
		go func(num int) {
			defer wg.Done()
			logger.Ptc("goroutine: ", num)
		}(i)
	}
	wg.Wait()
	logger.Ptc("finished")
}

func chTable() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)

	v1, ok1 := <-ch
	v2, ok2 := <-ch
	v3, ok3 := <-ch
	logger.Ptc(v1, ok1)
	logger.Ptc(v2, ok2)
	logger.Ptc(v3, ok3)
}

func RememberingAgain() {
	//startWorker()
	//simpleWgWait()
	chTable()
}
