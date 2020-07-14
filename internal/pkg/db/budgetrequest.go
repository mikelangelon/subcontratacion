package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BudgetRequestRepo struct {
	collection *mongo.Collection
}

func NewBudgetRequestRepo(db *mongo.Database) BudgetRequestRepo {
	return BudgetRequestRepo{
		collection: db.Collection("budgetrequest"),
	}
}

type BudgetRequestDB struct {
	ID                 primitive.ObjectID
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
