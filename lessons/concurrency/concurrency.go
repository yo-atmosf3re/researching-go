package concurrency

import (
	"context"
	"researching-go/lessons/concurrency/miner"
	"researching-go/lessons/concurrency/postman"
	"researching-go/pkg/logger"
	"sync"
	"sync/atomic"
	"time"
)

func minerAndPostmanExample() {
	var coal atomic.Int64
	var letters []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	mtx := sync.Mutex{}

	initTime := time.Now()

	go func() {
		time.Sleep(3 * time.Second)
		logger.Ptc("MINERS IS OVER")
		minerCancel()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		logger.Ptc("POSTMAN'S IS OVER")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 3)
	lettersTransferPoint := postman.PostmanPool(postmanContext, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := range coalTransferPoint {
			coal.Add(int64(c))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := range lettersTransferPoint {
			mtx.Lock()
			letters = append(letters, c)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	logger.Ptc("total coal :", coal.Load())
	mtx.Lock()
	logger.Ptc("total letters :", len(letters))
	mtx.Unlock()
	logger.Ptc("time is over: ", time.Since(initTime).Seconds())
}

func ConcurrencyExample() {
	minerAndPostmanExample()
}
