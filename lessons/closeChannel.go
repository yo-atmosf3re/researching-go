package lessons

import (
	"math/rand"
	"researching-go/pkg/logger"
	"time"
)

func axioms() {
	//__________________state: nil channel | state: closed channel|
	// operation:                          |                      |
	// close(channel)           panic      |       panic          |
	// read val: <-c            block      |    default value     |
	// write c <- val           block      |       panic          |
	// ____________________________________|______________________|

	//closeChannelAgain := func() {
	//	channel := make(chan int)
	//	close(channel)
	//	close(channel)
	//}
	//closeChannelAgain() // panic - close of closed channel

	//closeNilChannelAgainAndReadValue := func() {
	//	var channel chan int // nil value
	//	val := <-channel
	//	logger.Ptc(val, ": nil channel")
	//	close(channel)
	//	close(channel)
	//}
	//closeNilChannelAgain() // happened panic + block, because we tried read value from nil-channel

	//readValueFromDefaultValueChannel := func() {
	//	channel := make(chan int) // create new channel, this channel is open, value of channel - 0
	//
	//	go func() {
	//		channel <- 5   // now value of channel - 5
	//		close(channel) // now channel is closed
	//	}()
	//
	//	v := <-channel
	//	logger.Ptc("value of closed channel: ", v) // 5 value
	//}
	//readValueFromDefaultValueChannel()

	//writeInClosedChannel := func() {
	//	var channel chan int
	//	channel <- 5
	//}
	//writeInClosedChannel() // block because forbidden write value in nil-channel
}

func mineWork() {
	var transferPoint = make(chan int)

	go func() {
		iterations := 3 + rand.Intn(4)
		logger.Ptc("iterations: ", iterations)
		for range iterations {
			transferPoint <- 10
			time.Sleep(time.Millisecond * 200)
		}
		close(transferPoint) // 1. if we don't close this channel, that program is deadlock
	}()

	coal := 0
	for { // first variant getting value from channel
		value, ok := <-transferPoint // 2. because we will always get value from channel, "ok" never will not false. when channel will not have value we will try to get any value from this channel, that we got deadlock error
		if !ok {
			break
		}
		coal += value
		logger.Ptc("coal mined: ", coal)
	}

	for value := range transferPoint { // second variant, cycle will be completed, when all value from this channel will get
		coal += value
		logger.Ptc("coal mined: ", coal)
	}

	// CHANNEL, WHEN IT'S DON'T NEED NO MORE, SHOULD BE CLOSE ALWAYS. GETTING VALUE FROM CHANNEL = TRANSFER VALUE IN CHANNEL, ELSE WE WILL TO GET DEFAULT VALUE FROM CHANNEL.

}

func CloseChannelExample() {
	axioms()
	mineWork()
}
