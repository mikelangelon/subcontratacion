package budgetrequest

import (
	"time"
)

type BudgetRequest struct {
	ID                 string
	Comments           string
	CompanyID          string
	DeliveryConditions string
	Material           string
	Name               string
	Deadline           time.Time
	OfferDeadline      time.Time
	PaymentConditions  string
	Public             bool
	Quantity           int64
	QuantityDetail     string
}
