package tls_client

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

func GetPatch(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
	URL := config.Cfg.TLS_URL
	caCert, err := ioutil.ReadFile(config.Cfg.TLS_SERVER_CRT)
	if err != nil {
		log.Logrus.Fatalf("Error read server.crt file :%v", err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cert, err := tls.LoadX509KeyPair(config.Cfg.TLS_CLIENT_CRT, config.Cfg.TLS_CLIENT_KEY)
	if err != nil {
		log.Logrus.Fatalf("Error load tls x509 :%v", err)
		return
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{cert},
			},
		},
	}
	resp, err := client.Get(URL)
	if err != nil {
		log.Logrus.Fatalf("Error performing GET Method :%v", err)
		return
	}
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logrus.Fatalf("Error read response body :%v", err)
		return
	}
	defer resp.Body.Close()
	log.Logrus.Infof("%v\n", resp.Status)
	log.Logrus.Info(string(htmlData))
}
