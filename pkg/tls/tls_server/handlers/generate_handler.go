package handlers

import (
	"encoding/json"
	"net/http"
	mgo "smartcard/internal/app/model/mongocontrol"
	log "smartcard/pkg/logging"
	transfer "smartcard/pkg/tls/tls_server/transfer"
	"time"
)

func GenerateOneObject(w http.ResponseWriter, req *http.Request) {
	var batch []byte
	var err error
	object := mgo.CardData{
		Owner:      "Dima",
		TypeOfCard: "credit",
		CVC:        "453",
		Number:     "1234432112344321",
		ExpireDate: time.Now(),
	}
	batch, err = json.Marshal(object)
	if err != nil {
		log.Logrus.Fatalf("Error marshal object: %v", err)
	}
	log.Logrus.Info("BatchInfo is ready")
	transfer.BatchCh <- batch
	log.Logrus.Info("BatchInfo is successful go out")
}
