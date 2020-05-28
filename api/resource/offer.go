package resource

import (
	"mikelangelon/m/v2/api/model"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/api/rest/operation/offers"
	"mikelangelon/m/v2/internal/app/offer"

	"github.com/go-openapi/runtime/middleware"
)

type offerResource struct {
	service OfferService
}

type OfferService interface {
	CreateOffer(offer *offer.Offer) error
	GetAllOffers() ([]*offer.Offer, error)
}

func (o *offerResource) Register(api *operation.CoolappAPI) {
	api.OffersGetOffersHandler = offers.GetOffersHandlerFunc(o.HandlerGetAll)
	api.OffersPostOffersHandler = offers.PostOffersHandlerFunc(o.HandlerPost)
}

func (o *offerResource) HandlerPost(input offers.PostOffersParams) middleware.Responder {
	of := o.convertToOffer(input.Offer)
	err := o.service.CreateOffer(of)
	if err != nil {
		return offers.NewPostOffersInternalServerError()
	}
	return offers.NewPostOffersOK().WithPayload(o.convertToModel(of))
}

func (o *offerResource) HandlerGetAll(input offers.GetOffersParams) middleware.Responder {
	ofs, err := o.service.GetAllOffers()
	if err != nil {
		return offers.NewGetOffersInternalServerError()
	}
	var modelOfs []*model.Offer
	for _, v := range ofs {
		modelOfs = append(modelOfs, o.convertToModel(v))
	}
	return offers.NewGetOffersOK().WithPayload(modelOfs)
}

func (o *offerResource) convertToOffer(of *model.Offer) *offer.Offer {
	return &offer.Offer{
		ID:              of.ID,
		BudgetRequestID: *of.BudgetRequestID,
		CompanyID:       *of.CompanyID,
		Description:     of.Description,
	}
}

func (o *offerResource) convertToModel(of *offer.Offer) *model.Offer {
	return &model.Offer{
		ID:              of.ID,
		BudgetRequestID: &of.BudgetRequestID,
		CompanyID:       &of.CompanyID,
		Description:     of.Description,
	}
}
