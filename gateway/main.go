package main

import (
	"fmt"
	"net/http"
	initApp "api_gateway/gateway/core/init"
	requestHandler "api_gateway/gateway/core/request"
	monitors "api_gateway/gateway/core/monitor"
	"time"
)

func main() {
	fmt.Println("start App .. ")
	Router := initApp.ReadConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		enableCORS(&w)
		if (*r).Method == "OPTIONS" {
			return
		}
		monitors.PrintMonitor()
		requestHandler.HttpHandler(w, r, Router)
	})

	srv := &http.Server{
		Addr:         ":" + Router.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	srv.ListenAndServe()
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE , PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
