package carddatacontroller

import (
	"context"
	"fmt"
	"net/http"
	card_data_control "smartcard/internal/app/mongocontrol/card_data_control"
	transfer "smartcard/internal/tls/tls_server/transfer"
	"smartcard/internal/tools/conversion"
	log "smartcard/pkg/logging"
)

func errorResponse(w http.ResponseWriter, errText string) {
	fmt.Fprintf(w, "error = %v\nstatus:%v\n", errText, card_data_control.FAILED)
	return
}

func (ctrl *CardDataController) ListenBatch() {
	for {
		for batch := range transfer.BatchCh {
			cardData, err := conversion.GetDecodingCardData(batch)
			if err != nil {
				log.Logrus.Errorf("Error Unmarshal = %v", err)
				continue
			}
			id, err := ctrl.cardDataInteractor.AddOneCardData(context.TODO(), cardData)
			if err != nil {
				log.Logrus.Errorf("Error inserting data : %v", err)
				continue
			}
			log.Logrus.Infof("Inserted one document with id: %v", id.Hex())
		}
	}
}
