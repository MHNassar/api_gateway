package main

import (
	"fmt"
	"net/http"
	initApp "api_gateway/gateway/core/init"
	requestHandler "api_gateway/gateway/core/request"
	"time"
)

func main() {
	fmt.Println("start App .. ")
	Router := initApp.ReadConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/healthy-check" {
			return
		}

		enableCORS(&w)
		if (*r).Method == "OPTIONS" {
			return
		}

		requestHandler.HttpHandler(w, r, Router)
	})

	srv := &http.Server{
		Addr:         ":" + Router.Port,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	fmt.Println("ListenAndServe HTTP 80")

	srv.ListenAndServe()

}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE , PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
