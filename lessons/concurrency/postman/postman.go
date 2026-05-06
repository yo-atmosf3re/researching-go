package postman

import (
	"context"
	"researching-go/pkg/logger"
	"sync"
	"time"
)

func postman(
	transferPoint chan<- string,
	n int, mail string,
	wg *sync.WaitGroup,
	ctx context.Context,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			logger.Ptc("postman number of ", n, "finished to deliver letter")
			return
		default:
			logger.Ptc("i'm postman number of :", n, "take letter")
			time.Sleep(1 * time.Second)
			logger.Ptc("i'm postman number of :", n, "bring letter to post, letter - ", mail)

			transferPoint <- mail

			logger.Ptc("i'm postman number of :", n, "delivered letter")

		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go postman(mailTransferPoint, i, postmanToMail(i), wg, ctx)
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint

}

func postmanToMail(postmanNumber int) string {
	letters := map[int]string{
		1: "Magazine",
		2: "List",
		3: "Book",
	}

	letter, ok := letters[postmanNumber]
	if !ok {
		return "Lottery"
	}
	return letter
}
