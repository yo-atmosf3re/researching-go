package backend

import (
	"io"
	"net/http"
	"researching-go/pkg/logger"
	"strconv"
	"sync"
)

func responseWriter(w http.ResponseWriter, messages []string) error {
	var msg string
	for i, message := range messages {
		msg += message
		if i != len(messages)-1 {
			msg += " "
		}
	}
	_, err := w.Write([]byte(msg))
	return err
}

// var money = atomic.Int64{} // handler it is function inside goroutine, so atomic should be used + money it is variable which using in different http handle function, mtx is expected
// var bank = atomic.Int64{}
var money = 1000 // second example without atomic, because we used mutex for work with variables
var bank = 0
var mtx sync.Mutex // declare mutex can outside functions globally, e.g. near other variables, because we can't receive mtx as argument in http handle function

func handlePayment(w http.ResponseWriter, r *http.Request) {
	//r.Header - contains headers of request/response, client or server
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rBody, err := io.ReadAll(r.Body)
	logger.Ptc("body read", string(rBody))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // easy way type status code response with help WriteHeader method
		// WriteHeader should to call before Write method, else status code response will be 200
		err = responseWriter(w, []string{"during reading body error", err.Error()})
		if err != nil {
			logger.Ptc("fail to write HTTP response: ", err)
		}
		return
	}

	rBodyString := string(rBody)
	paymentAmount, err := strconv.Atoi(rBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = responseWriter(w, []string{"during parsing payment amount occurred error, maybe payment amount is not integer.", err.Error()})
		if err != nil {
			logger.Ptc("fail to write HTTP response: ", err)
		}
		return
	}

	mtx.Lock()
	defer mtx.Unlock()
	if money-paymentAmount < 0 {
		err := responseWriter(w, []string{"insufficient funds. balance:", strconv.Itoa(money)})
		if err != nil {
			logger.Ptc("fail to write HTTP response: ", err)
		}
		return
	}
	money -= paymentAmount
	err = responseWriter(w, []string{"payment is success. balance:", strconv.Itoa(money)})
	if err != nil {
		logger.Ptc("fail to write HTTP response: ", err)
	}
}

func handleSaveBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "during reading body error" + err.Error()
		logger.Ptc(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			logger.Ptc("fail to write HTTP response: ", err)
		}
		return
	}
	httpRequestBodyString := string(httpRequestBody)       // convert byte[] to string
	saveAmount, err := strconv.Atoi(httpRequestBodyString) // than convert string to int
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "during parsing save amount occurred error" + err.Error() // if saveAmount is not int, that we got error
		logger.Ptc(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			logger.Ptc("fail to write HTTP response: ", err)
		}
		return
	}

	mtx.Lock()
	defer mtx.Unlock()
	if money-saveAmount < 0 {
		err := responseWriter(w, []string{"insufficient funds for save balance. balance:", strconv.Itoa(money), "bank:", strconv.Itoa(bank)})
		if err != nil {
			logger.Ptc("fail to write HTTP response: ", err)
		}
		return
	}
	money -= saveAmount
	bank += saveAmount
	err = responseWriter(w, []string{"money:", strconv.Itoa(money), "bank:", strconv.Itoa(bank)})
	if err != nil {
		logger.Ptc("fail to write HTTP response: ", err)
	}
	return
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
