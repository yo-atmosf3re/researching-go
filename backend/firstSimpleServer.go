package backend

import (
	"net/http"
	"researching-go/pkg/logger"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	str := "hello world"
	b := []byte(str)

	_, err := w.Write(b) // first argument - quantity writing bytes, second argument - error
	if err != nil {
		logger.Ptc("error is occurred", err.Error())
	} else {
		logger.Ptc("http response handle is correctly")
	}

}

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "pay is finished"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		logger.Ptc("error is occurred", err.Error())
	} else {
		logger.Ptc("http response handle is correctly")
	}

}

func handlerWithEmptyPattern(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world, empty pattern - root pattern"))
	if err != nil {
		logger.Ptc("error is occurred", err.Error())
	}
	logger.Ptc("http response handle is correctly")
}

func FirstServerExample() {
	port := ":9091"
	http.HandleFunc("/default", defaultHandler) // second arg - handler, some function, which allows response writer and request reader, more details later, first arg - it is endpoint, e.g. "/default", "/pay" etc.
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/", handlerWithEmptyPattern) // "/" - root endpoint
	logger.Ptc("starting server on port ", port)
	err := http.ListenAndServe(port, nil) // start server on specified port in first arg, second arg - handler, e.g. function defaultHandler above
	if err != nil {
		logger.Ptc("during starting server error is occurred", err.Error())
	}

	// every handler is it function inside goroutine, handlers works concurrency

}
