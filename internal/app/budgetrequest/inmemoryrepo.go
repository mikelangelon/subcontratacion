package budgetrequest

import (
	"errors"
	"fmt"
	"time"
)

type inMemoryRepo struct {
	repo map[string]*BudgetRequest
}

func newRepo() *inMemoryRepo {
	m := feedMap()
	return &inMemoryRepo{m}
}

func (i *inMemoryRepo) Create(budgetRequest *BudgetRequest) error {
	budgetRequest.ID = fmt.Sprintf("%d", len(i.repo)+1)
	i.repo[budgetRequest.ID] = budgetRequest
	return nil
}

func (i *inMemoryRepo) Search() ([]*BudgetRequest, error) {
	return i.GetAll()
}

func (i *inMemoryRepo) GetAll() ([]*BudgetRequest, error) {
	var budgetRequests []*BudgetRequest
	for _, v := range i.repo {
		budgetRequests = append(budgetRequests, v)
	}
	return budgetRequests, nil
}

func (i *inMemoryRepo) GetLasts() ([]*BudgetRequest, error) {
	var budgetRequests []*BudgetRequest
	for _, v := range i.repo {
		budgetRequests = append(budgetRequests, v)
	}
	len := len(budgetRequests)
	if len < 4 {
		return budgetRequests[0:len], nil
	}

	return budgetRequests[len-4 : len], nil
}

func (i *inMemoryRepo) Get(id string) (*BudgetRequest, error) {
	val, ok := i.repo[id]
	if ok {
		return val, nil
	}
	return nil, errors.New("not found")
}

func feedMap() map[string]*BudgetRequest {
	return map[string]*BudgetRequest{
		"1": {
			ID:                 "1",
			Name:               "Cool stuff",
			Material:           "Acer0",
			Treatments:         []string{"Cosas", "blablabla"},
			Quantity:           3,
			QuantityDetail:     "Algo",
			Deadline:           time.Date(2020, 1, 1, 1, 1, 2, 1, time.UTC),
			OfferDeadline:      time.Date(2020, 1, 1, 1, 1, 2, 1, time.UTC),
			DeliveryConditions: "Mas cosas",
			PaymentConditions:  "Quizas nunca",
			Comments:           "Lalalala",
			CompanyID:          "",
			Public:             false,
		},
		"2": {
			ID:                 "2",
			Name:               "Cool stuff",
			Material:           "Metaquilato",
			Treatments:         []string{"Cosas", "blablabla"},
			Quantity:           100,
			QuantityDetail:     "Algo",
			Deadline:           time.Date(2020, 1, 1, 1, 1, 2, 1, time.UTC),
			OfferDeadline:      time.Date(2020, 1, 1, 1, 1, 2, 1, time.UTC),
			DeliveryConditions: "Mas cosas",
			PaymentConditions:  "Quizas nunca",
			Comments:           "Lalalala",
			CompanyID:          "",
			Public:             false,
		},
		"3": {
			ID:                 "3",
			Name:               "Cool stuff",
			Material:           "Madera",
			Treatments:         []string{"Cosas", "blablabla"},
			Quantity:           25,
			QuantityDetail:     "Algo",
			Deadline:           time.Date(2020, 1, 1, 1, 1, 2, 1, time.UTC),
			OfferDeadline:      time.Date(2020, 1, 1, 1, 1, 2, 1, time.UTC),
			DeliveryConditions: "Mas cosas",
			PaymentConditions:  "Quizas nunca",
			Comments:           "Lalalala",
			CompanyID:          "",
			Public:             false,
		},
		"4": {
			ID:                 "4",
			Name:               "Cool stuff",
			Material:           "Acero",
			Quantity:           300,
			QuantityDetail:     "Algo",
			Treatments:         []string{"Cosas", "blablabla"},
			Deadline:           time.Date(2020, 1, 1, 1, 1, 2, 1, time.UTC),
			OfferDeadline:      time.Date(2020, 1, 1, 1, 1, 2, 1, time.UTC),
			DeliveryConditions: "Mas cosas",
			PaymentConditions:  "Quizas nunca",
			Comments:           "Lalalala",
			CompanyID:          "",
			Public:             false,
		},
	}
}
