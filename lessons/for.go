package lessons

import (
	"fmt"
	"math/rand"
	"time"
)

func pt(value ...any) (n int, err error) {
	return fmt.Println(value...)
}

func ForExample() {
	index := 0
	for ; index <= 5; index++ {
		pt("first", index)
	}
	pt(index)
}

func WithRand() {
	randomValue := rand.Intn(10)

	for i := 0; i <= randomValue; i++ {
		pt("second", i)
	}
}

func SleepExample() {
	for i := 0; i <= 5; i++ {
		pt("third", i)
		time.Sleep(1 * time.Minute)
	}
}

func Infinity() {
	score := 0
	randValue := rand.Intn(20)
	for {
		pt("before: ", score)
		pt("rand value: ", randValue)
		if randValue == score {
			pt("end")
			break
		}
		score++
		pt("after: ", score)
		time.Sleep(500 * time.Millisecond)
	}
	pt("total: ", score)
}
