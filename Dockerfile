FROM golang:1.19-alpine3.16 AS build-stage
WORKDIR /artifacts
COPY go.mod go.sum ./
RUN GOOS=linux GOARCH=amd64 go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o smartcard-service cmd/main.go

FROM --platform=linux/amd64 alpine:3.16
WORKDIR /app
COPY --from=build-stage /artifacts/smartcard-service .
CMD ["/app/smartcard-service"]