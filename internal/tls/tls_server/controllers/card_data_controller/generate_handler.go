package carddatacontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	card_data_control "smartcard/internal/app/mongocontrol/card_data_control"
	carddatacases "smartcard/internal/clearcore/use_cases/card_data_cases"
	transfer "smartcard/internal/tls/tls_server/transfer"
	log "smartcard/pkg/logging"
	"time"

	"github.com/Pallinder/go-randomdata"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// for generate Object use
// curl -k --cert tls_client/client.crt --key tls_client/client.key -X GET https://localhost:1443/generate
func (ctrl *CardDataController) GenerateHandler(w http.ResponseWriter, req *http.Request) {
	var batch []byte
	var status card_data_control.StateOper
	var errText string

	status = card_data_control.SUCCESS
	object := carddatacases.CardData{
		ID:         primitive.NewObjectID(),
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
		status = card_data_control.FAILED
	}
	log.Logrus.Info("BatchInfo is ready")
	transfer.BatchCh <- batch
	log.Logrus.Info("BatchInfo is successful go out")
	if err == nil && errText == "" && status == card_data_control.SUCCESS {
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
