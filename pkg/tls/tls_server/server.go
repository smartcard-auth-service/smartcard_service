package tls_server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"smartcard/config"
	log "smartcard/pkg/logging"
	"sync"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world!\n"))
}

func Run(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
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
	srv := &http.Server{
		Addr:      config.Cfg.TLS_SERVER_LISTEN_PORT,
		Handler:   &handler{},
		TLSConfig: cfg,
	}
	log.Logrus.Debugf("Listening TLS server to port number :%v", config.Cfg.TLS_SERVER_LISTEN_PORT)
	log.Logrus.Info(srv.ListenAndServeTLS(config.Cfg.TLS_SERVER_CRT, config.Cfg.TLS_SERVER_KEY))
}
