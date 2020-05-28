package resource

import (
	"mikelangelon/m/v2/api/model"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/api/rest/operation/companies"
	"mikelangelon/m/v2/internal/app/company"

	"github.com/go-openapi/runtime/middleware"
)

type companyResource struct {
	service CompanyService
}

type CompanyService interface {
	CreateCompany(company *company.Company) error
	UpdateCompany(company *company.Company) error
	GetCompany(companyId string) (*company.Company, error)
}

func (c *companyResource) Register(api *operation.CoolappAPI) {
	api.CompaniesPostCompaniesHandler = companies.PostCompaniesHandlerFunc(c.HandlerPost)
	api.CompaniesPutCompaniesCompanyIDHandler = companies.PutCompaniesCompanyIDHandlerFunc(c.HandlerPut)
	api.CompaniesGetCompaniesCompanyIDHandler = companies.GetCompaniesCompanyIDHandlerFunc(c.HandlerGet)
}

func (c *companyResource) HandlerPost(input companies.PostCompaniesParams) middleware.Responder {
	com := convertToCompany(input.Company)
	err := c.service.CreateCompany(com)
	if err != nil {
		return companies.NewPostCompaniesInternalServerError()
	}
	return companies.NewPostCompaniesOK().WithPayload(convertToModel(com))
}

func (c *companyResource) HandlerPut(input companies.PutCompaniesCompanyIDParams) middleware.Responder {
	com := convertToCompany(input.Company)
	err := c.service.UpdateCompany(com)
	if err != nil {
		return companies.NewPutCompaniesCompanyIDInternalServerError()
	}
	return companies.NewPostCompaniesOK()
}

func (c *companyResource) HandlerGet(input companies.GetCompaniesCompanyIDParams) middleware.Responder {
	com, err := c.service.GetCompany(input.CompanyID)
	if err != nil {
		return companies.NewGetCompaniesCompanyIDInternalServerError()
	}
	return companies.NewGetCompaniesCompanyIDOK().WithPayload(convertToModel(com))
}

func convertToCompany(model *model.Company) *company.Company {
	return &company.Company{
		Name:          *model.Name,
		CIF:           model.Cif,
		City:          model.City,
		CP:            model.Cp,
		Description:   model.Description,
		Employees:     model.Employees,
		FundationYear: model.FundationYear,
		Quantity:      model.Quantity,
		Street:        model.Street,
		Web:           model.Web,
	}
}

func convertToModel(company *company.Company) *model.Company {
	return &model.Company{
		Cif:           company.CIF,
		City:          company.City,
		Cp:            company.CP,
		Description:   company.Description,
		Employees:     company.Employees,
		FundationYear: company.FundationYear,
		ID:            company.ID,
		Name:          &company.Name,
		Quantity:      company.Quantity,
		Street:        company.Street,
		Web:           company.Web,
	}
}
