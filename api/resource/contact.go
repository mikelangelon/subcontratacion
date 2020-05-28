package resource

import (
	"mikelangelon/m/v2/api/model"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/api/rest/operation/contacts"
	"mikelangelon/m/v2/internal/app/company"

	"github.com/go-openapi/runtime/middleware"
)

type contactResource struct {
	service ContactService
}

type ContactService interface {
	CreateContact(contact *company.Contact) error
	UpdateContact(contact *company.Contact) error
	GetContact(contactId string) (*company.Contact, error)
	GetAllContactsFor(companyId string) ([]*company.Contact, error)
}

func (c *contactResource) Register(api *operation.CoolappAPI) {
	api.ContactsPostContactsHandler = contacts.PostContactsHandlerFunc(c.HandlerPost)
	api.ContactsPutContactsContactIDHandler = contacts.PutContactsContactIDHandlerFunc(c.HandlerPut)
	api.ContactsGetContactsContactIDHandler = contacts.GetContactsContactIDHandlerFunc(c.HandlerGet)
}

func (c *contactResource) HandlerPost(input contacts.PostContactsParams) middleware.Responder {
	con := c.convertToContact(input.Contact)
	err := c.service.CreateContact(con)
	if err != nil {
		return contacts.NewPostContactsInternalServerError()
	}
	return contacts.NewPostContactsOK().WithPayload(c.convertToModel(con))
}

func (c *contactResource) HandlerPut(input contacts.PutContactsContactIDParams) middleware.Responder {
	con := c.convertToContact(input.Contact)
	err := c.service.UpdateContact(con)
	if err != nil {
		return contacts.NewPutContactsContactIDInternalServerError()
	}
	return contacts.NewPutContactsContactIDOK().WithPayload(c.convertToModel(con))
}

func (c *contactResource) HandlerGet(input contacts.GetContactsContactIDParams) middleware.Responder {
	con, err := c.service.GetContact(input.ContactID)
	if err != nil {
		return contacts.NewGetContactsContactIDInternalServerError()
	}
	return contacts.NewGetContactsContactIDOK().WithPayload(c.convertToModel(con))
}

func (c *contactResource) convertToContact(model *model.Contact) *company.Contact {
	return &company.Contact{
		ID:        model.ID,
		Name:      *model.Name,
		Surname:   *model.Surname,
		CompanyID: *model.CompanyID,
		Email:     *model.Email,
		Fax:       model.Fax,
		Phone:     *model.Phone,
	}
}

func (c *contactResource) convertToModel(con *company.Contact) *model.Contact {
	return &model.Contact{
		CompanyID: &con.CompanyID,
		Email:     &con.Email,
		Fax:       con.Fax,
		ID:        con.ID,
		Name:      &con.Name,
		Phone:     &con.Phone,
		Surname:   &con.Surname,
	}
}
