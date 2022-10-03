package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smartcard/internal/app/mongocontrol/model"
	"smartcard/internal/app/mongocontrol/service"
	log "smartcard/pkg/logging"
	transfer "smartcard/pkg/tls/tls_server/transfer"
	"time"

	"github.com/Pallinder/go-randomdata"
)

// for generate Object use
// curl -k --cert tls_client/client.crt --key tls_client/client.key -X GET https://localhost:1443/generate
func GenerateHandler(w http.ResponseWriter, req *http.Request) {
	var batch []byte
	var status service.StateOper
	var errText string

	status = service.SUCCESS
	object := model.CardData{
		Owner:      getOwner(),
		TypeOfCard: "credit",
		CVC:        getCVC(),
		Number:     getNumberOfCard(),
		ExpireDate: getExpireDate(),
	}
	batch, err := json.Marshal(object)
	if err != nil {
		log.Logrus.Errorf("Error marshal object: %v", err)
		errText = fmt.Sprintf("Error marshal object: %v", err)
		status = service.FAILED
	}
	log.Logrus.Info("BatchInfo is ready")
	transfer.BatchCh <- batch
	log.Logrus.Info("BatchInfo is successful go out")
	if err == nil && errText == "" && status == service.SUCCESS {
		fmt.Fprintf(w, "status = %v\n", status)
	} else {
		fmt.Fprintf(w, "error = %v, status = %v", errText, status)
	}
}

func getOwner() string {
	return fmt.Sprintf("%v %v", randomdata.FirstName(randomdata.Male), randomdata.LastName())
}

func getCVC() string {
	return fmt.Sprint(randomdata.Number(100, 999))
}

func getNumberOfCard() string {
	return fmt.Sprint(randomdata.StringNumberExt(4, "-", 4))
}

func getExpireDate() time.Time {
	expTime, err := time.Parse("1970-01-01", randomdata.FullDateInRange("2023-01-01", "2029-12-31"))
	if err != nil {
		log.Logrus.Error(err)
		return time.Now()
	}
	return expTime
}
