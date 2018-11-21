// Copyright 2016 kadiray karakaya. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	initApp "api_gateway/gateway/core/init"
	requestHandler "api_gateway/gateway/core/request"
)

func main() {
	Router := initApp.ReadConfig()
	fmt.Printf("listening on port: %v\n", Router.Port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// To Enable CORS
		enableCORS(&w)
		if (*r).Method == "OPTIONS" {
			return
		}
		requestHandler.HttpHandler(w, r, Router)
	})

	http.ListenAndServe(":"+Router.Port, nil)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
