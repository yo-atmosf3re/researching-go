package backend

import (
	"io"
	"net/http"
	"researching-go/pkg/logger"
	"strconv"
	"sync"
	"time"
)

// var money = atomic.Int64{} // handler it is function inside goroutine, so atomic should be used + money it is variable which using in different http handle function, mtx is expected
// var bank = atomic.Int64{}
var money = 1000 // second example without atomic, because we used mutex for work with variables
var bank = 0
var mtx sync.Mutex // declare mutex can outside functions globally, e.g. near other variables, because we can't receive mtx as argument in http handle function

func handlePayment(_ http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	logger.Ptc("body read", string(rBody))
	if err != nil {
		msg := "during reading body error" + err.Error()
		logger.Ptc(msg)
		return
	}

	rBodyString := string(rBody)
	paymentAmount, err := strconv.Atoi(rBodyString)
	if err != nil {
		logger.Ptc("during parsing payment amount occurred error, maybe payment amount is not integer", err.Error())
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {
		time.Sleep(3 * time.Second)
		money = -paymentAmount
	} else {
		logger.Ptc("payment: is not enough money")
	}
	logger.Ptc("money", money)
	mtx.Unlock()

}

func handleSaveBalance(_ http.ResponseWriter, r *http.Request) {
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

	mtx.Lock()
	if money-saveAmount >= 0 {
		time.Sleep(3 * time.Second)
		money = -saveAmount // convert int to int64 for atomic
		bank = saveAmount
		logger.Ptc("money :", money, "bank :", bank)
	} else {
		logger.Ptc("save balance: insufficient funds")
	}
	mtx.Unlock()
}

func run() {
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
