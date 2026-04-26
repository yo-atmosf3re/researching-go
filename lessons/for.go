package lessons

import (
	"math/rand"
	"researching-go/pkg/logger"
	"time"
)

func ForExample() {
	index := 0
	for ; index <= 5; index++ {
		logger.Pt("first", index)
	}
	logger.Pt(index)
}

func WithRand() {
	randomValue := rand.Intn(10)

	for i := 0; i <= randomValue; i++ {
		logger.Pt("second", i)
	}
}

func SleepExample() {
	for i := 0; i <= 5; i++ {
		logger.Pt("third", i)
		time.Sleep(1 * time.Minute)
	}
}

func Infinity() {
	score := 0
	randValue := rand.Intn(20)
	for {
		logger.Pt("before: ", score)
		logger.Pt("rand value: ", randValue)
		if randValue == score {
			logger.Pt("end")
			break
		}
		score++
		logger.Pt("after: ", score)
		time.Sleep(500 * time.Millisecond)
	}
	logger.Pt("total: ", score)
}
