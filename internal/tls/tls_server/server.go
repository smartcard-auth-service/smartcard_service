package tls_server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	mongo_client "smartcard/internal/app/mongo_client/client"
	carddatacases "smartcard/internal/clearcore/use_cases/card_data_cases"
	card_data_controller "smartcard/internal/tls/tls_server/controllers/card_data_controller"
	"smartcard/internal/tls/tls_server/transfer"

	"smartcard/pkg/config"
	log "smartcard/pkg/logging"
	"sync"

	"github.com/gorilla/mux"
)

func Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer transfer.CloseTransferChan()
	transfer.InitTransferChan()

	caCert, err := ioutil.ReadFile(config.Cfg.TLS_CLIENT_CRT)
	if err != nil {
		log.Logrus.Info("Shutdown tls server")
		log.Logrus.Fatalf("Error read client.crt file :%v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	cardDataController := card_data_controller.NewCardDataController(
		carddatacases.NewCardDataInteractor(mongo_client.Mgo))
	r := mux.NewRouter()
	r.HandleFunc("/generate", cardDataController.GenerateHandler)
	r.HandleFunc("/get", cardDataController.GetHandler)

	srv := &http.Server{
		Addr:      config.Cfg.TLS_LISTEN_HOST,
		Handler:   r,
		TLSConfig: cfg,
	}
	go cardDataController.ListenBatch()

	log.Logrus.Tracef("Listening TLS server to host %v", config.Cfg.TLS_LISTEN_HOST)
	log.Logrus.Info(srv.ListenAndServeTLS(config.Cfg.TLS_SERVER_CRT, config.Cfg.TLS_SERVER_KEY))
}
