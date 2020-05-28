package resource

import (
	"mikelangelon/m/v2/api/model"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/api/rest/operation/budget_requests"
	"mikelangelon/m/v2/internal/app/budgetrequest"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type budgetRequestResource struct {
	service BudgetRequestService
}

type BudgetRequestService interface {
	CreateBudgetRequest(budgetRequest *budgetrequest.BudgetRequest) error
	GetAllBudgetRequests() ([]*budgetrequest.BudgetRequest, error)
}

func (b *budgetRequestResource) Register(api *operation.CoolappAPI) {
	api.BudgetRequestsGetBudgetRequestsHandler = budget_requests.GetBudgetRequestsHandlerFunc(b.HandlerGetAll)
	api.BudgetRequestsPostBudgetRequestsHandler = budget_requests.PostBudgetRequestsHandlerFunc(b.HandlerPost)
}

func (b *budgetRequestResource) HandlerPost(input budget_requests.PostBudgetRequestsParams) middleware.Responder {
	bq := b.convertToBudgetRequest(input.BudgetRequest)
	err := b.service.CreateBudgetRequest(bq)
	if err != nil {
		return budget_requests.NewPostBudgetRequestsInternalServerError()
	}
	return budget_requests.NewPostBudgetRequestsOK().WithPayload(b.convertToModel(bq))
}

func (b *budgetRequestResource) HandlerGetAll(input budget_requests.GetBudgetRequestsParams) middleware.Responder {
	bqs, err := b.service.GetAllBudgetRequests()
	if err != nil {
		return budget_requests.NewGetBudgetRequestsInternalServerError()
	}
	var modelBQs []*model.BudgetRequest
	for _, v := range bqs {
		modelBQs = append(modelBQs, b.convertToModel(v))
	}
	return budget_requests.NewGetBudgetRequestsOK().WithPayload(modelBQs)
}

func (b *budgetRequestResource) convertToBudgetRequest(bud *model.BudgetRequest) *budgetrequest.BudgetRequest {
	return &budgetrequest.BudgetRequest{
		ID:                 bud.ID,
		Comments:           bud.Comments,
		CompanyID:          *bud.CompanyID,
		DeliveryConditions: bud.DeliveryConditions,
		Material:           bud.Material,
		Name:               *bud.Name,
		Deadline:           time.Time(bud.Deadline),
		OfferDeadline:      time.Time(bud.OfferDeadline),
		PaymentConditions:  bud.PaymentConditions,
		Public:             bud.Public,
		Quantity:           *bud.Quantity,
		QuantityDetail:     bud.QuantityDetails,
	}
}

func (b *budgetRequestResource) convertToModel(bud *budgetrequest.BudgetRequest) *model.BudgetRequest {
	return &model.BudgetRequest{
		Comments:           bud.Comments,
		CompanyID:          &bud.CompanyID,
		Deadline:           strfmt.Date(bud.Deadline),
		ID:                 bud.ID,
		Material:           bud.Material,
		Name:               &bud.Name,
		OfferDeadline:      strfmt.Date(bud.OfferDeadline),
		PaymentConditions:  bud.PaymentConditions,
		Public:             bud.Public,
		Quantity:           &bud.Quantity,
		QuantityDetails:    bud.QuantityDetail,
		DeliveryConditions: bud.DeliveryConditions,
	}
}
