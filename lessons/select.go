package lessons

import (
	"researching-go/pkg/logger"
	"time"
)

func selectExample1() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		time.Sleep(40 * time.Millisecond)
		intCh <- 1

	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		strCh <- "str"

	}()

	time.Sleep(50 * time.Millisecond)

	select {
	case number := <-intCh:
		logger.Ptc("intCh: ", number)
	case str := <-strCh:
		logger.Ptc("strCh: ", str)
	default: // this code works if case-channels is not ready. if case-channel condition is ready, that default don't work
		logger.Ptc("channel is not ready")
	}

}

type message struct {
	author string
	text   string
}

func selectExample2() {
	messageChannel1 := make(chan message)
	messageChannel2 := make(chan message)

	go func() {
		for {
			messageChannel1 <- message{
				author: "Alex",
				text:   "hello",
			}
			time.Sleep(10 * time.Second)
		}

	}()
	go func() {
		for {
			messageChannel2 <- message{
				author: "Dave",
				text:   "hello",
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case msg1 := <-messageChannel1:
			logger.Ptc("i got message from: ", msg1.author, "text of message", msg1.text)
		case msg2 := <-messageChannel2:
			logger.Ptc("i got message from: ", msg2.author, "text of message", msg2.text)
		}
	}
}

func SelectExample() {
	selectExample2()
}
