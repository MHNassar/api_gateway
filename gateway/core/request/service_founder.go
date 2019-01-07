package request

import (
	"strings"
	AppCore "api_gateway/gateway/core"
	AppLogger "api_gateway/gateway/core/logger"
)

func checkServiceExist(router AppCore.Router, originalPath string) (AppCore.Services, error) {
	service_name_rray := strings.Split(originalPath, "/")
	service_prefix := service_name_rray[1]

	service := getService(router, service_prefix)

	if service.ServicePrefix == "" {
		service = getService(router, "default")
	}

	logger := AppLogger.GetLogInstance()
	logger.AddStep("checkServiceExist : Every Thing Is Good", "")
	return service, nil

}

func getService(router AppCore.Router, service_prefix string) AppCore.Services {
	var service AppCore.Services

	for _, v := range router.Services {
		if v.ServicePrefix == service_prefix {
			service = v
		}
	}
	return service
}
