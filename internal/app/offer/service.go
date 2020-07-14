package offer

type service struct {
	repo inMemoryRepo
}

func New() *service {
	return &service{
		repo: inMemoryRepo{
			repo: map[string]*Offer{},
		},
	}
}

func (s *service) Create(offer *Offer) error {
	return s.Create(offer)
}

func (s *service) GetAll() ([]*Offer, error) {
	return s.repo.GetAll()
}
