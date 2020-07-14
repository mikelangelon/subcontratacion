package web

import (
	o "mikelangelon/m/v2/internal/app/offer"
	"net/http"
)

type OfferResource struct {
	OfferService offerService
	OkCall       func(w http.ResponseWriter, r *http.Request)
}

type offerService interface {
	Create(budgetRequest *o.Offer) error
}

func (r OfferResource) Create(resp http.ResponseWriter, req *http.Request) {
	r.OfferService.Create(
		&o.Offer{
			Description: "",
		},
	)

	r.OkCall(resp, req)
}
