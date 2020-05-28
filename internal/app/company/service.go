package company

type service struct {
	repo         Repo
	repoContacts RepoContacts
}

func New() *service {
	return &service{
		repo: &inMemoryRepo{
			repo: map[string]*Company{},
		},
		repoContacts: &inMemoryRepoContacts{
			map[string]*Contact{},
		},
	}
}

type Repo interface {
	Create(company *Company) error
	Update(company *Company) error
	Get(companyId string) (*Company, error)
}

type RepoContacts interface {
	Create(contact *Contact) error
	Update(contact *Contact) error
	Get(contactID string) (*Contact, error)
	GetAllFor(companyID string) ([]*Contact, error)
}

func (s *service) CreateCompany(company *Company) error {
	return s.repo.Create(company)
}

func (s *service) UpdateCompany(company *Company) error {
	return s.repo.Create(company)
}
func (s *service) GetCompany(companyId string) (*Company, error) {
	return s.repo.Get(companyId)
}

func (s *service) CreateContact(contact *Contact) error {
	return s.repoContacts.Create(contact)
}

func (s *service) UpdateContact(contact *Contact) error {
	return s.repoContacts.Create(contact)
}

func (s *service) GetContact(contactId string) (*Contact, error) {
	return s.repoContacts.Get(contactId)
}

func (s *service) GetAllContactsFor(companyId string) ([]*Contact, error) {
	return s.repoContacts.GetAllFor(companyId)
}
