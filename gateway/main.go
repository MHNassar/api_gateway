package main

import (
	"fmt"
	"net/http"
	initApp "api_gateway/gateway/core/init"
	requestHandler "api_gateway/gateway/core/request"
	monitors "api_gateway/gateway/core/monitor"
	//"golang.org/x/crypto/acme/autocert"
	"time"
	//"context"
	//"crypto/tls"
)

func main() {
	fmt.Println("start App .. ")
	Router := initApp.ReadConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" {
			return
		}
		enableCORS(&w)
		if (*r).Method == "OPTIONS" {
			return
		}
		monitors.PrintMonitor()
		requestHandler.HttpHandler(w, r, Router)
	})

	//go func() {
	//	var m *autocert.Manager
	//	hostPolicy := func(ctx context.Context, host string) error {
	//		// Note: change to your real host
	//		allowedHost := "sandbox.bolt-play.com"
	//		if host == allowedHost {
	//			return nil
	//		}
	//		return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
	//	}
	//	dataDir := "."
	//	m = &autocert.Manager{
	//		Prompt:     autocert.AcceptTOS,
	//		HostPolicy: hostPolicy,
	//		Cache:      autocert.DirCache(dataDir),
	//	}
	//
	//	srv := &http.Server{
	//		Addr:         ":443",
	//		ReadTimeout:  5 * time.Second,
	//		WriteTimeout: 10 * time.Second,
	//		TLSConfig:    &tls.Config{GetCertificate: m.GetCertificate},
	//	}
	//	fmt.Println("ListenAndServe HTTPS 443")
	//	srv.ListenAndServe()
	//
	//}()

	srv := &http.Server{
		Addr:         ":" + Router.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("ListenAndServe HTTP 80")

	srv.ListenAndServe()

}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE , PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
