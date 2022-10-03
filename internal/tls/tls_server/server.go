package tls_server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"smartcard/config"
	log "smartcard/pkg/logging"
	"smartcard/pkg/tls/tls_server/handlers"
	"smartcard/pkg/tls/tls_server/listen"
	"smartcard/pkg/tls/tls_server/transfer"
	"sync"

	"github.com/gorilla/mux"
)

func Run(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
	defer transfer.CloseTransferChan()
	transfer.InitTransferChan()

	caCert, err := ioutil.ReadFile(config.Cfg.TLS_CLIENT_CRT)
	if err != nil {
		log.Logrus.Fatalf("Error read client.crt file :%v", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	r := mux.NewRouter()
	r.HandleFunc("/generate", handlers.GenerateHandler)
	r.HandleFunc("/get", handlers.GetHandler)

	srv := &http.Server{
		Addr:      config.Cfg.TLS_SERVER_LISTEN_PORT,
		Handler:   r,
		TLSConfig: cfg,
	}
	go listen.ListenBatch()

	log.Logrus.Debugf("Listening TLS server to port number %v", config.Cfg.TLS_SERVER_LISTEN_PORT)
	log.Logrus.Info(srv.ListenAndServeTLS(config.Cfg.TLS_SERVER_CRT, config.Cfg.TLS_SERVER_KEY))
}
