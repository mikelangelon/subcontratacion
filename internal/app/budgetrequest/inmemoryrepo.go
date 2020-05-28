package budgetrequest

import "fmt"

type inMemoryRepo struct {
	repo map[string]*BudgetRequest
}

func (i *inMemoryRepo) Create(budgetRequest *BudgetRequest) error {
	budgetRequest.ID = fmt.Sprintf("%d", len(i.repo)+1)
	i.repo[budgetRequest.ID] = budgetRequest
	return nil
}

func (i *inMemoryRepo) GetAll() ([]*BudgetRequest, error) {
	var budgetRequests []*BudgetRequest
	for _, v := range i.repo {
		budgetRequests = append(budgetRequests, v)
	}
	return budgetRequests, nil
}
