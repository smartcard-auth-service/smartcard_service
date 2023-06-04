FROM golang:latest
WORKDIR /home/smartcard
COPY . .
RUN go build -o signal cmd/main.go

FROM scratch
WORKDIR /home/smartcard
COPY --from=builder /opt/signal /opt/signal
ENTRYPOINT [ "/opt/signal" ] 
