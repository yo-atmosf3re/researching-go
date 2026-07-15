package practice

import (
	"researching-go/pkg/logger"
	"sync"
	"time"
)

func fetchStatus(serviceName string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	ch <- serviceName + " started is successful!"
}

func RememberingMedium() {
	wg := &sync.WaitGroup{}
	var serviceChannel = make(chan string)
	wg.Add(3)
	go fetchStatus("Auth", serviceChannel, wg)
	go fetchStatus("Database", serviceChannel, wg)
	go fetchStatus("Cache", serviceChannel, wg)

	go func() {
		wg.Wait()
		close(serviceChannel)
	}()

	for message := range serviceChannel {
		logger.Ptc(message)
	}
}
