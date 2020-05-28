package budgetrequest

type service struct {
	repo inMemoryRepo
}

func (s *service) Create(budgetRequest *BudgetRequest) error {
	return s.Create(budgetRequest)
}

func (s *service) GetAll() ([]*BudgetRequest, error) {
	return s.repo.GetAll()
}
