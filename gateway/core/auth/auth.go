package auth

import (
	"net/http"
	"fmt"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"os"
	AppLogger "api_gateway/gateway/core/logger"
	"encoding/base64"
)

func CheckAuth(r *http.Request, auth bool) (string, error) {
	logger := AppLogger.GetLogInstance()
	if auth {
		token := getToken(r)
		if token == "" {

			err := fmt.Errorf("Token not found")
			return "", err
		}
		tokenData, err := parseToken(token)
		if err != nil {
			return "", err
		}

		base64encoded := base64.StdEncoding.EncodeToString([]byte(tokenData))
		encText, err := encryption(base64encoded)
		if err != nil {
			return "", err
		}
		fmt.Println(encText)
		logger.AddStep("CheckAuth : Every Thing Is Good ", "")
		return encText, nil
	}
	logger.AddStep("CheckAuth : Request Not Needs Auth", " ")
	return "", nil
}

func getToken(r *http.Request) string {
	logger := AppLogger.GetLogInstance()
	var token string
	token = r.Header.Get("Authorization")
	if token == "" {
		logger.AddStep("getToken", "Token Not Found")
	}
	logger.AddStep("getToken", "")
	return token
}

func parseToken(tokenString string) (string, error) {
	logger := AppLogger.GetLogInstance()
	tokenList := strings.Split(tokenString, "Bearer")
	if len(tokenList) != 2 {
		err := fmt.Errorf("Token Not Found")
		logger.AddStep("parseToken", "Token Not Found")
		return "", err
	}
	tokenString = strings.TrimSpace(tokenList[1])

	signingKey := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		logger.AddStep("parseToken", err.Error())
		return "", err
	}
	var tokenData string
	for key, val := range claims {
		if key == "sub" {
			data, _ := json.Marshal(val)
			tokenData = string(data)
		}
	}
	logger.AddStep("parseToken", "")
	return tokenData, nil
}
