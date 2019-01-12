package request

import (
	AppCore "api_gateway/gateway/core"
	AppLogger "api_gateway/gateway/core/logger"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"net"
)

func createRequest(r *http.Request, forwardPath AppCore.TargetPath, originalReq string, msg string) (*http.Request, error) {
	logger := AppLogger.GetLogInstance()
	newPath := forwardPath.Path + originalReq
	if r.URL.RawQuery != "" {
		newPath += "?" + r.URL.RawQuery
	}

	req_content_type := r.Header.Get("Content-Type")
	req, err := http.NewRequest(r.Method, newPath, r.Body)
	if err != nil {
		logger.AddStep("createRequest", err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", req_content_type)
	req.Header.Set("Message", msg)
	// ToDO temp when finish integration
	req_token := r.Header.Get("Authorization")
	req.Header.Set("Authorization", req_token)

	logger.ForwardPath = newPath
	logger.AddStep("createRequest : Every Thing Is Good ", "")

	return req, nil
}

func sendRequest(w http.ResponseWriter, req *http.Request, router AppCore.Router) int {

	logger := AppLogger.GetLogInstance()

	client := createHttpClient(router)

	defer handelPanicRequest()
	resp, err := client.Do(req)
	if err != nil {
		logger.AddStep("HttpHandler", err.Error())
		AppLogger.DestroyLogInstance()
		fmt.Println("not req")
		AppCore.ShowError(w, err, http.StatusBadGateway)
		return 0
	}
	//
	defer resp.Body.Close()
	//
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.AddStep("HttpHandler", err.Error())
		AppLogger.DestroyLogInstance()
		AppCore.ShowError(w, err, http.StatusBadGateway)
		return 0
	}

	headerResp := strings.Join(resp.Header["Content-Type"], "")
	w.Header().Set("Content-Type", headerResp)
	logger.AddStep("HttpHandler : Request Send Successfully", "")
	logger.EndTime = time.Now()
	logger.Status = true

	AppLogger.DestroyLogInstance()
	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(body))
	return 0
}

func handelPanicRequest() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func createHttpClient(router AppCore.Router) *http.Client {
	timeOutValue := router.Settings.TimeOut
	timeout := time.Duration(timeOutValue * time.Second)
	keepAliveTimeout := 600 * time.Second
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAliveTimeout,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: netTransport,
		Timeout:   timeout}
	return client
}
