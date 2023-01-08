FROM golang:latest
WORKDIR /home/smartcard
COPY . .
RUN go build -o main cmd/main.go
CMD ./main