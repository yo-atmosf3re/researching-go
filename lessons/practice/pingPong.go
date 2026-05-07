package practice

import (
	"fmt"
	"researching-go/pkg/logger"
	"time"
)

func player(name string, table chan int) {
	for {
		ball, ok := <-table
		if !ok {
			return
		}

		fmt.Printf("player %s hit ball! hit №%d\n", name, ball)
		time.Sleep(500 * time.Millisecond)

		ball++
		table <- ball
	}
}

func PingPongExample() {
	table := make(chan int)

	go player("ping", table)
	go player("pong", table)

	table <- 1

	time.Sleep(5 * time.Second)

	close(table)

	logger.Ptc("game is over")
}
