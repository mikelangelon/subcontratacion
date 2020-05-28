package company

import (
	"errors"
	"fmt"
)

type inMemoryRepoContacts struct {
	repo map[string]*Contact
}

func (i *inMemoryRepoContacts) Create(contact *Contact) error {
	contact.ID = fmt.Sprintf("%d", len(i.repo)+1)
	i.repo[contact.ID] = contact
	return nil
}

func (i *inMemoryRepoContacts) Update(contact *Contact) error {
	_, ok := i.repo[contact.ID]
	if !ok {
		return errors.New("not found")
	}
	i.repo[contact.ID] = contact
	return nil
}

func (i *inMemoryRepoContacts) Get(contactID string) (*Contact, error) {
	val, ok := i.repo[contactID]
	if !ok {
		return nil, errors.New("not found")
	}
	return val, nil
}

func (i *inMemoryRepoContacts) GetAllFor(companyID string) ([]*Contact, error) {
	var contacts []*Contact
	for _, v := range i.repo {
		if v.CompanyID == companyID {
			contacts = append(contacts, v)
		}
	}
	return contacts, nil
}
