package miner

import (
	"context"
	"math/rand"
	"researching-go/pkg/logger"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n, power int,
) { // send-only channel type in function, read is forbidden
	defer wg.Done()
	for {
		logger.Ptc("i'm miner number of ", n, "is started to mine coal")

		select {
		case <-ctx.Done():
			logger.Ptc("i'm miner number of", n, "my work day is finished")
			return
		case <-time.After(1 * time.Second):
			logger.Ptc("i'm miner number of ", n, "is mined coal: ", power)
		}

		// in postman.go there is first of type using select{}, when it's using goroutines finishing is fully, even if goroutine is middle through case of select.
		// in miner.go goroutines is finished immediately, when context of goroutine is closed in parent goroutine, because before executing goroutine it's waits any time. if during waiting of time, context of goroutine is closed, that case when time.After is not completed, case context.Done() will be completed.
		// setting value in channel looks like above case select, if context is not closed - setting value will be completed, else goroutine will be finished

		select {
		case <-ctx.Done():
			logger.Ptc("i'm miner number of", n, "my work day is finished")
			return
		case transferPoint <- power:
			logger.Ptc("i'm miner number of ", n, "is finished to mine coal: ", power)
		}
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, 1+i*rand.Intn(10))
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
