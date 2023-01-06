package carddatacontroller

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	card_data_control "smartcard/internal/app/mongocontrol/card_data_control"
	carddatacases "smartcard/internal/clearcore/use_cases/card_data_cases"
	"smartcard/internal/tools/conversion"

	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// for get Object by id
// curl -k --cert tls_client/client.crt --key tls_client/client.key -X GET https://localhost:1443/get/id=12345678
func (ctrl *CardDataController) GetHandler(w http.ResponseWriter, req *http.Request) {
	var idCard primitive.ObjectID
	var card *carddatacases.CardData
	var query url.Values
	var status card_data_control.StateOper
	var errText string

	status = card_data_control.SUCCESS
	parseUrl, err := url.Parse(req.URL.String())
	if err != nil {
		log.Logrus.Errorf("Error parse url: %v", err)
		errText = fmt.Sprintf("Error parse url :%v", err)
		status = card_data_control.FAILED
	}
	if !card_data_control.IsFailed(status) {
		query, err = url.ParseQuery(parseUrl.RawQuery)
		if err != nil {
			log.Logrus.Errorf("Error parse query %v Error: %v", parseUrl.RawQuery, err)
			errText = fmt.Sprintf("Error parse query %v Error: %v", parseUrl.RawQuery, err)
			status = card_data_control.FAILED
		}
	}
	if !card_data_control.IsFailed(status) {
		id := query.Get("id")
		idCard, err = conversion.GetIdCard(id)
		if err != nil {
			log.Logrus.Errorf("Unable get object with id = '%v' Error %v", id, err)
			errText = fmt.Sprintf("Unable get object with id = '%v' Error %v", id, err)
			status = card_data_control.FAILED
		}
	}
	if !card_data_control.IsFailed(status) {
		card, err = ctrl.cardDataInteractor.GetOneCardData(context.TODO(), idCard)
		if err != nil {
			log.Logrus.Errorf("Error get data : %v", err)
			errText = fmt.Sprintf("Error get card data : %v", err)
			status = card_data_control.FAILED
		}
	}

	if err == nil && errText == "" && status == card_data_control.SUCCESS {
		fmt.Fprintf(w, "status = %v\ndata:\n%v", status, card.String())
	} else {
		fmt.Fprintf(w, "error = %v, status = %v", errText, status)
	}
}
