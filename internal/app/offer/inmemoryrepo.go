package offer

import (
	"fmt"
)

type inMemoryRepo struct {
	repo map[string]*Offer
}

func (i *inMemoryRepo) Create(offer *Offer) error {
	offer.ID = fmt.Sprintf("%d", len(i.repo)+1)
	i.repo[offer.ID] = offer
	return nil
}

func (i *inMemoryRepo) GetAll() ([]*Offer, error) {
	var budgetRequests []*Offer
	for _, v := range i.repo {
		budgetRequests = append(budgetRequests, v)
	}
	return budgetRequests, nil
}
