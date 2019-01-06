package request

import (
	AppCore "api_gateway/gateway/core"
	AppAuth "api_gateway/gateway/core/auth"
	AppLogger "api_gateway/gateway/core/logger"
	"net/http"
	"fmt"
)

func HttpHandler(w http.ResponseWriter, r *http.Request, router AppCore.Router) int {

	originalPath := r.URL.Path

	// init logger
	logger := AppLogger.GetLogInstance()
	logger.InitLog(originalPath)

	service, _ := checkServiceExist(router, originalPath)

	msg, err := AppAuth.CheckAuth(r, service.TargetPath.Auth)
	if err != nil {
		AppLogger.DestroyLogInstance()
		fmt.Println("not message")
		AppCore.ShowError(w, err, http.StatusUnauthorized)
		return 0
	}
	var req *http.Request

	defaultForwardPath := service.TargetPath

	req, err = createRequest(r, defaultForwardPath, originalPath, msg)
	if err != nil {
		logger.AddStep("HttpHandler", err.Error())
		AppLogger.DestroyLogInstance()
		fmt.Println("not createRequest")
		AppCore.ShowError(w, err, http.StatusBadGateway)
		return 0
	}

	fmt.Printf("forwarded to default :%v\n", req.URL)
	res := sendRequest(w, req, router)
	return res
}
