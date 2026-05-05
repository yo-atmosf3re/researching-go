package lessons

import (
	"context"
	"researching-go/pkg/logger"
	"time"
)

func boo(context context.Context) {
	for {
		select {
		case <-context.Done():
			logger.Ptc("boo is completed")
			return
		default:
			logger.Ptc("boo")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func doo(context context.Context) {
	for {
		select {
		case <-context.Done():
			logger.Ptc("doo is completed")
			return
		default:
			logger.Ptc("doo")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func contextUsage() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext) // child context nested parent context, when parent context will be canceled, child context too

	go boo(parentContext)
	go doo(childContext)

	time.Sleep(3 * time.Second)
	childCancel()
	time.Sleep(2 * time.Second)
	parentCancel()
}

func ContextExample() {
	contextUsage()
}
