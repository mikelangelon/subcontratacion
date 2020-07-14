package web

import (
	"html/template"
	"mikelangelon/m/v2/internal/app/budgetrequest"
	"mikelangelon/m/v2/internal/app/company"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type BudgetRequestResource struct {
	BudgetService budgetService
	OkCall        func(w http.ResponseWriter, r *http.Request)
}

type budgetService interface {
	Create(budgetRequest *budgetrequest.BudgetRequest) error
	GetLasts() ([]*budgetrequest.BudgetRequest, error)
	Search() ([]*budgetrequest.BudgetRequest, error)
	Get(id string) (*budgetrequest.BudgetRequest, error)
}

func (r BudgetRequestResource) Create(resp http.ResponseWriter, req *http.Request) {

	deadline, err := time.Parse("2006-01-02", req.FormValue("Deadline"))
	if err != nil {

	}
	offerDeadline, err := time.Parse("2006-01-02", req.FormValue("OfferDeadline"))
	if err != nil {

	}
	quantity, err := strconv.ParseInt(req.FormValue("Quantity"), 10, 64)
	if err != nil {

	}
	b := &budgetrequest.BudgetRequest{
		Name:               req.FormValue("Name"),
		Material:           req.FormValue("Material"),
		QuantityDetail:     req.FormValue("QuantityDetail"),
		Deadline:           deadline,
		OfferDeadline:      offerDeadline,
		Quantity:           quantity,
		DeliveryConditions: req.FormValue("DeliveryConditions"),
		PaymentConditions:  req.FormValue("PaymentConditions"),
	}
	r.BudgetService.Create(b)

	log.WithField("budgetRequest", b).Info("Storing budget")
	r.OkCall(resp, req)
}

type enrichedBudgetRequest struct {
	budgetrequest.BudgetRequest
	Company company.Company
}
type budgetSearch struct {
	Profile    *Profile
	AllBudgets []enrichedBudgetRequest
}

func (r BudgetRequestResource) Search(resp http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles(
		"web/template/layout.html",
		"web/template/topmenu.html",
		"web/template/leftmenu.html",
		"web/template/banner.html",
		"web/content/budgetrequest-search.html")

	result, err := r.BudgetService.Search()
	if err != nil {
		// TODO handle errors
	}

	var enriched []enrichedBudgetRequest
	for _, v := range result {
		enriched = append(enriched, enrichedBudgetRequest{
			BudgetRequest: *v,
			Company: company.Company{
				Name: "Blablabla SA",
			},
		})
	}

	prof := budgetSearch{
		AllBudgets: enriched,
	}
	err = t.Execute(resp, prof)
	if err != nil {
		log.Println(err)
	}
}

func (r BudgetRequestResource) View(resp http.ResponseWriter, req *http.Request) {
	params, ok := req.URL.Query()["id"]
	if !ok || len(params[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		return
	}
	id := params[0]
	log.Println(id)
}
