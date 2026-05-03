package lessons

import (
	"researching-go/pkg/logger"
	"strconv"
	"time"
)

func selectExample1() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		i := 1
		for {
			strCh <- "str" + strconv.Itoa(i)
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case number := <-intCh:
			logger.Ptc("intCh: ", number)
		case str := <-strCh:
			logger.Ptc("strCh: ", str)
		}
	}

}

func SelectExample() {
	selectExample1()
}
