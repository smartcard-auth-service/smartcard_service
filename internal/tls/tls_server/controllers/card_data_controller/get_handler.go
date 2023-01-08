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
	var errText string

	parseUrl, err := url.Parse(req.URL.String())
	if err != nil {
		errText = fmt.Sprintf("Error parse url :%v", err)
		log.Logrus.Errorf(errText)
		errorResponse(w, errText)
		return
	}
	query, err = url.ParseQuery(parseUrl.RawQuery)
	if err != nil {
		errText = fmt.Sprintf("Error parse query %v: %v", parseUrl.RawQuery, err)
		log.Logrus.Errorf(errText)
		errorResponse(w, errText)
		return
	}
	id := query.Get("id")
	idCard, err = conversion.GetIdCard(id)
	if err != nil {
		errText = fmt.Sprintf("Unable get object with id = '%v' Error %v", id, err)
		log.Logrus.Errorf(errText)
		errorResponse(w, errText)
		return
	}
	card, err = ctrl.cardDataInteractor.GetOneCardData(context.TODO(), idCard)
	if err != nil {
		errText = fmt.Sprintf("Error get card data : %v", err)
		log.Logrus.Errorf(errText)
		errorResponse(w, errText)
		return
	}
	fmt.Fprintf(w, "status = %v\ndata:\n%v", card_data_control.SUCCESS, card.String())
	return
}
