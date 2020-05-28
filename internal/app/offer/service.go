package offer

type service struct {
	repo inMemoryRepo
}

func (s *service) Create(budgetRequest *Offer) error {
	return s.Create(budgetRequest)
}

func (s *service) GetAll() ([]*Offer, error) {
	return s.repo.GetAll()
}
