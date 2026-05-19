package backend

import (
	"io"
	"net/http"
	"researching-go/pkg/logger"
	"strconv"
	"sync/atomic"
)

var money = atomic.Int64{} // handler it is function inside goroutine, so atomic should be used
var bank = atomic.Int64{}

func handlePayment(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	logger.Ptc("body read", string(rBody))
	if err != nil {
		logger.Ptc("during reading body error", err.Error())
		return
	}

	rBodyString := string(rBody)
	paymentAmount, err := strconv.Atoi(rBodyString)
	if err != nil {
		logger.Ptc("during parsing payment amount occurred error, maybe payment amount is not integer", err.Error())
		return
	}

	if money.Load()-int64(paymentAmount) >= 0 {
		money.Add(int64(-paymentAmount))
	} else {
		logger.Ptc("is not enough money")
	}
	logger.Ptc("money", money.Load())

}

func handleSaveBalance(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Ptc("during reading body error", err.Error())
		return
	}
	httpRequestBodyString := string(httpRequestBody)       // convert byte[] to string
	saveAmount, err := strconv.Atoi(httpRequestBodyString) // than convert string to int
	if err != nil {
		logger.Ptc("during parsing save amount occurred error", err.Error()) // if saveAmount is not int, that we got error
		return
	}

	if money.Load()-int64(saveAmount) >= 0 {
		money.Add(int64(-saveAmount)) // convert int to int64 for atomic
		bank.Add(int64(saveAmount))
		logger.Ptc("money :", money.Load(), "bank :", bank.Load())
	} else {
		logger.Ptc("insufficient funds")
	}
}

func run() {
	money.Add(1000)
	http.HandleFunc("/pay", handlePayment)
	http.HandleFunc("/save", handleSaveBalance)
	logger.Ptc("starting server")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		logger.Ptc("during start server error is occurred", err.Error())
	}
}

func BodyHttpRequestExample() {
	run()
}
