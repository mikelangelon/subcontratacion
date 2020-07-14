package company

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type inMemoryRepo struct {
	repo map[string]*Company
}

func (i *inMemoryRepo) Create(company *Company) error {
	logrus.Info("repo save company")
	company.ID = fmt.Sprintf("%d", len(i.repo)+1)
	i.repo[company.ID] = company
	return nil
}

func (i *inMemoryRepo) Update(company *Company) error {
	_, ok := i.repo[company.ID]
	if !ok {
		return errors.New("not found")
	}
	i.repo[company.ID] = company
	return nil
}

func (i *inMemoryRepo) Get(companyId string) (*Company, error) {
	val, ok := i.repo[companyId]
	if !ok {
		return nil, errors.New("not found")
	}
	return val, nil
}

func (i *inMemoryRepo) GetLastCompanies() ([]*Company, error) {
	logrus.Info("repo GetLastCompanies")
	var companies []*Company
	for _, v := range i.repo {
		companies = append(companies, v)
	}
	return companies, nil
}
