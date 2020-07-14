package budgetrequest

import "time"

type BudgetRequest struct {
	ID                 string
	Name               string
	Material           string
	Treatments         []string
	Quantity           int64
	QuantityDetail     string
	Deadline           time.Time
	OfferDeadline      time.Time
	DeliveryConditions string
	PaymentConditions  string
	Comments           string
	CompanyID          string
	Public             bool
}
