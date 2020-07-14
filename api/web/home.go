package web

import (
	"fmt"
	"html/template"
	"mikelangelon/m/v2/internal/app/budgetrequest"
	"mikelangelon/m/v2/internal/app/company"
	"net/http"

	"github.com/sirupsen/logrus"
)

type HomeResource struct {
	BudgetService  budgetService
	CompanyService CompanyService
}

type home struct {
	Profile       *Profile
	Offers        []offer
	LastProviders []provider
}

type offer struct {
	Name         string
	Quantity     int64
	Description  string
	AmountOffers int64
}

type provider struct {
	Name      string
	City      string
	Web       string
	Employees string
}

func (h HomeResource) Home(resp http.ResponseWriter, req *http.Request) {
	profile, _ := LoginResource{}.Auth(resp, req)
	t, err := template.ParseFiles(
		"web/template/layout.html",
		"web/template/topmenu.html",
		"web/template/leftmenu.html",
		"web/template/budgetrequests.html",
		"web/template/banner.html",
		"web/template/providers.html",
		"web/content/home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	budgetRequests, err := h.BudgetService.GetLasts()
	if err != nil {
		// TODO
	}

	companies, err := h.CompanyService.GetLastCompanies()
	if err != nil {
		logrus.WithError(err).Info("eh")
	}

	home := &home{
		Profile:       profile,
		Offers:        prepareBudgetRequests(budgetRequests),
		LastProviders: prepareCompanies(companies)}
	t.Execute(resp, home)
}

func prepareBudgetRequests(requests []*budgetrequest.BudgetRequest) []offer {
	var offers []offer
	for _, v := range requests {
		offers = append(offers, offer{
			Name:         v.Name,
			Quantity:     v.Quantity,
			Description:  v.Treatments[0],
			AmountOffers: 0,
		})
	}
	return offers
}

func prepareCompanies(companies []*company.Company) []provider {
	var providers []provider
	for _, v := range companies {
		providers = append(providers, provider{
			Name:      v.Name,
			City:      v.City,
			Web:       v.Web,
			Employees: v.Employees,
		})
	}
	return providers
}
