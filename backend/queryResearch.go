package backend

import (
	"net/http"
	"researching-go/pkg/logger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	logger.Ptc("foo: ", fooParam, "boo: ", booParam)
}

func QueryResearch() {
	//http.HandleFunc("/query", handler)
	if err := http.ListenAndServe(":9091", http.HandlerFunc(handler)); err != nil {
		logger.Ptc("error starting server")
		return
	}
}
