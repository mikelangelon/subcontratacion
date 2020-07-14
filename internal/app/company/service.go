package company

type CompanyService struct {
	repo         Repo
	repoContacts RepoContacts
}

func New(repo Repo) CompanyService {
	return CompanyService{
		repo: repo,
		repoContacts: &inMemoryRepoContacts{
			map[string]*Contact{},
		},
	}
}

type Repo interface {
	Save(company Company) error
	Update(company Company) error
	Get(id string) (*Company, error)
	Latests(limit int64) ([]*Company, error)
}

type RepoContacts interface {
	Create(contact *Contact) error
	Update(contact *Contact) error
	Get(contactID string) (*Contact, error)
	GetAllFor(companyID string) ([]*Contact, error)
}

func (s CompanyService) CreateCompany(company Company) error {
	return s.repo.Save(company)
}

func (s CompanyService) UpdateCompany(company Company) error {
	return s.repo.Update(company)
}
func (s CompanyService) GetCompany(companyId string) (*Company, error) {
	return s.repo.Get(companyId)
}

func (s CompanyService) GetLastCompanies() ([]*Company, error) {
	return s.repo.Latests(4)
}

func (s CompanyService) CreateContact(contact *Contact) error {
	return s.repoContacts.Create(contact)
}

func (s CompanyService) UpdateContact(contact *Contact) error {
	return s.repoContacts.Create(contact)
}

func (s CompanyService) GetContact(contactId string) (*Contact, error) {
	return s.repoContacts.Get(contactId)
}

func (s CompanyService) GetAllContactsFor(companyId string) ([]*Contact, error) {
	return s.repoContacts.GetAllFor(companyId)
}
