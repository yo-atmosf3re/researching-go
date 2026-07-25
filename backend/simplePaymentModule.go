package backend

import (
	"encoding/json"
	"net/http"
	"researching-go/pkg/logger"
	"strconv"
	"sync"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
}

func (p Payment) Buyer() string {
	return p.FullName + ", " + p.Address
}

var useMutex = sync.Mutex{}
var balance = 1000
var paymentHistory = make([]Payment, 0)

func simplePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		// I need to use pointer for payment variable here, because this variable will be changed
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	useMutex.Lock()
	if balance-payment.USD < 0 {
		if err := ResponseWriter(w, []string{"insufficient funds. balance:", strconv.Itoa(balance)}); err != nil {
			logger.Ptc("fail to write HTTP response: ", err)
		}
		return
	}
	balance -= payment.USD
	paymentHistory = append(paymentHistory, payment)
	var descriptions string
	for _, v := range paymentHistory {
		descriptions += v.Description + " "
	}
	if err := ResponseWriter(w, []string{"balance:", strconv.Itoa(balance), "history purchase:", descriptions}); err != nil {
		logger.Ptc("fail to write HTTP response: ", err)
	}
	defer useMutex.Unlock()

}

func SimplePaymentModuleSetup() {
	http.HandleFunc("/pay", simplePaymentHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil { // environment err variable only inside "if"
		logger.Ptc("error starting server")
	}
}
