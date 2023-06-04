FROM golang:1.20.4-alpine3.17 AS builder
WORKDIR /opt/akd
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o signal cmd/main.go

FROM scratch
WORKDIR /opt/akd
COPY . .
COPY --from=builder /opt/akd/signal /opt/akd/signal
ENTRYPOINT [ "/opt/akd/signal" ] 
