
FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates
FROM scratch
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/



FROM golang:latest

RUN go get github.com/dgrijalva/jwt-go github.com/sadlil/gologger golang.org/x/crypto/acme/autocert

COPY . /go/src/api_gateway/gateway

WORKDIR /go/src/api_gateway/gateway

RUN go install api_gateway/gateway


ENTRYPOINT /go/bin/gateway
