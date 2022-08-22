# syntax=docker/dockerfil1e:1
FROM golang:1.16-alpine
WORKDIR /smartcard_service
COPY go.mod ./
COPY go.sum ./
COPY cmd/ /smartcard_service/cmd/
COPY internal/ /smartcard_service/internal/
COPY pkg/ /smartcard_service/pkg/
RUN go mod download
RUN go run cmd/main.go
USER tatun2000