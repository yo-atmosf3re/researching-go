package backend

import (
	"encoding/json"
	"net/http"
	"researching-go/pkg/logger"
	"strconv"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Time        time.Time
}

type HttpResponse struct {
	Money          int       `json:"money"`
	PaymentHistory []Payment `json:"paymentHistory"`
}

func (p Payment) Buyer() string {
	return p.FullName + ", " + p.Address
}

func getDescriptions(ph []Payment) string {
	var descriptions string
	for _, v := range ph {
		descriptions += v.Description + " "
	}
	return descriptions
}

var useMutex = sync.Mutex{}
var balance = 1000
var paymentHistory = make([]Payment, 0)

func simplePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil { // I need to use pointer for payment variable here, because this variable will be changed
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payment.Time = time.Now()

	useMutex.Lock()
	if balance-payment.USD < 0 {
		if err := ResponseStringWriter(w, []string{"insufficient funds. balance:", strconv.Itoa(balance)}); err != nil {
			logger.Ptc("fail to write HTTP response: ", err)
		}
		return
	}
	balance -= payment.USD
	paymentHistory = append(paymentHistory, payment)

	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}
	b, err := json.MarshalIndent(httpResponse, "", "	") // simple Marshal method returns nonformatted JSON, but MarshalIndent returns formatted JSON via indent (tab, space, etc.)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		logger.Ptc("fail to write HTTP response: ", err)
		return
	}
	defer useMutex.Unlock()

}

func SimplePaymentModuleSetup() {
	//http.HandleFunc("/pay", simplePaymentHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil { // environment err variable only inside "if"
		logger.Ptc("error starting server")
	}
}
