package listen

import (
	"context"
	"smartcard/internal/app/mongocontrol/service"
	grpc "smartcard/pkg/grpc/server"
	log "smartcard/pkg/logging"
	transfer "smartcard/pkg/tls/tls_server/transfer"
)

func ListenBatch() {
	for {
		for batch := range transfer.BatchCh {
			cardData, err := grpc.GetDecodingCardData(batch)
			if err != nil {
				log.Logrus.Errorf("Error Unmarshal = %v", err)
				continue
			}
			id, err := service.AddOne(context.TODO(), cardData)
			if err != nil {
				log.Logrus.Errorf("Error inserting data : %v", err)
				continue
			}
			log.Logrus.Infof("Inserted one document with id: %v", id)
		}
	}
}
