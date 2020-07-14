package budgetrequest

type service struct {
	repo *inMemoryRepo
}

func New() *service {
	return &service{
		repo: newRepo(),
	}
}

func (s *service) Create(budgetRequest *BudgetRequest) error {
	return s.repo.Create(budgetRequest)
}

func (s *service) GetAll() ([]*BudgetRequest, error) {
	return s.repo.GetAll()
}
func (s *service) GetLasts() ([]*BudgetRequest, error) {
	return s.repo.GetLasts()
}

func (s *service) Get(id string) (*BudgetRequest, error) {
	return s.repo.Get(id)
}
func (s *service) Search() ([]*BudgetRequest, error) {
	return s.repo.Search()
}
