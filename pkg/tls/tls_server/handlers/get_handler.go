package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"smartcard/internal/app/mongocontrol/model"
	"smartcard/internal/app/mongocontrol/service"
	grpc "smartcard/pkg/grpc/server"
	log "smartcard/pkg/logging"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetHandler(w http.ResponseWriter, req *http.Request) {
	var idCard *primitive.ObjectID
	var card *model.CardData
	var query url.Values
	var status service.StateOper
	var errText string

	status = service.SUCCESS
	parseUrl, err := url.Parse(req.URL.String())
	if err != nil {
		log.Logrus.Errorf("Error parse url: %v", err)
		errText = fmt.Sprintf("Error parse url :%v", err)
		status = service.FAILED
	}
	if !service.IsFailed(status) {
		query, err = url.ParseQuery(parseUrl.RawQuery)
		if err != nil {
			log.Logrus.Errorf("Error parse query %v Error: %v", parseUrl.RawQuery, err)
			errText = fmt.Sprintf("Error parse query %v Error: %v", parseUrl.RawQuery, err)
			status = service.FAILED
		}
	}
	if !service.IsFailed(status) {
		for key, value := range query {
			if key != "id" {
				log.Logrus.Errorf("Key must be equal 'id', key = '%v' Error: %v", key, err)
				errText = fmt.Sprintf("Key must be equal 'id', key = '%v' Error: %v", key, err)
				status = service.FAILED
				break
			}
			idCard, err = grpc.GetIdCard(value[0])
			if err != nil {
				log.Logrus.Errorf("Unable get object with id = '%v' Error %v", value[0], err)
				errText = fmt.Sprintf("Unable get object with id = '%v' Error %v", value[0], err)
				status = service.FAILED
			}
			break
		}
	}
	if !service.IsFailed(status) {
		card, err = service.GetOne(context.TODO(), idCard)
		if err != nil {
			log.Logrus.Errorf("Error get data : %v", err)
			errText = fmt.Sprintf("Error get card data : %v", err)
			status = service.FAILED
		}
	}

	if err == nil && errText == "" && status == service.SUCCESS {
		fmt.Fprintf(w, "status = %v\ndata:\n%v", status, card.String())
	} else {
		fmt.Fprintf(w, "error = %v, status = %v", errText, status)
	}
}
