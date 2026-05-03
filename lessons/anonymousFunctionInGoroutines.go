package lessons

import (
	"researching-go/pkg/logger"
	"time"
)

func foo() {
	for {
		logger.Ptc("foo")
		time.Sleep(100 * time.Millisecond)
	}
}

func AnonFuncInGoroutines() {
	go foo()

	go func() {
		logger.Ptc("anon")
		time.Sleep(100 * time.Millisecond)
	}()

	time.Sleep(1 * time.Second)
}
