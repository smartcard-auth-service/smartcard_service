package carddatacases

import (
	"context"
	"fmt"
	cli "smartcard/internal/app/mongo_client/client"
	carddatacontrol "smartcard/internal/app/mongocontrol/card_data_control"
	domain "smartcard/internal/clearcore/domain/card_data_domain"
	log "smartcard/pkg/logging"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardData struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Owner      string             `bson:"owner" json:"owner"`
	TypeOfCard string             `bson:"type_of_card" json:"type_of_card"`
	CVC        string             `bson:"cvc" json:"cvc"`
	Number     string             `bson:"number" json:"number"`
	ExpireDate time.Time          `bson:"date" json:"date"`
}

func (c *CardData) String() string {
	return fmt.Sprintf("Owner: %v\nTypeOfCard: %v\nCVC: %v\nNumber: %v\nExpireDate: %v\n",
		c.Owner, c.TypeOfCard, c.CVC, c.Number, c.ExpireDate)
}

type CardDataInteractor struct {
	cardDataRepository domain.CardDataRepositorier
}

func NewCardDataInteractor(mgoDriver *cli.MgoDriver) *CardDataInteractor {
	return &CardDataInteractor{
		cardDataRepository: carddatacontrol.NewCardDataRepository(mgoDriver),
	}
}

func (interactor *CardDataInteractor) GetOneCardData(ctx context.Context, id primitive.ObjectID) (cardData *CardData, err error) {
	rawCardData, err := interactor.cardDataRepository.GetOne(ctx, id)
	if err != nil {
		log.Logrus.Errorf("Error getting card data: %v", err)
	}
	return cardDataOutDto(rawCardData), nil
}

func (interactor *CardDataInteractor) AddOneCardData(ctx context.Context, cardData *CardData) (id primitive.ObjectID, err error) {
	id, err = interactor.cardDataRepository.AddOne(ctx, cardDataInDto(cardData))
	if err != nil {
		log.Logrus.Errorf("Error adding card data: %v", err)
		return
	}
	return
}

func cardDataInDto(src *CardData) (dst *domain.CardData) {
	dst = &domain.CardData{
		ID:         src.ID,
		Owner:      src.Owner,
		TypeOfCard: src.TypeOfCard,
		CVC:        src.CVC,
		Number:     src.Number,
		ExpireDate: src.ExpireDate,
	}
	return dst
}

func cardDataOutDto(src *domain.CardData) (dst *CardData) {
	dst = &CardData{
		ID:         src.ID,
		Owner:      src.Owner,
		TypeOfCard: src.TypeOfCard,
		CVC:        src.CVC,
		Number:     src.Number,
		ExpireDate: src.ExpireDate,
	}
	return dst
}
