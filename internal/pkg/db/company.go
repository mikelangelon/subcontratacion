package db

import (
	"context"
	"mikelangelon/m/v2/internal/app/company"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CompanysRepo struct {
	collection *mongo.Collection
}

func NewCompanyRepo(db *mongo.Database) CompanysRepo {
	return CompanysRepo{
		collection: db.Collection("company"),
	}
}

type companyDB struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	CIF           string             `bson:"cif"`
	Address       AddressDB          `bson:"address"`
	Description   string             `bson:"description"`
	Employees     string             `bson:"employees"`
	FundationYear int64              `bson:"fundationYear"`
	Web           string             `bson:"web"`
	Email         string             `bson:"email"`
	CreationAt    time.Time          `bson:"creationAt"`
	Logo          []byte             `bson:"logo"`
}

type AddressDB struct {
	City   string `bson:"city"`
	CP     string `bson:"cp"`
	Street string `bson:"street"`
}

func (c CompanysRepo) Save(company company.Company) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	comp := toCompanyDB(company)
	comp.CreationAt = time.Now()
	comp.ID = primitive.NewObjectID()
	_, err := c.collection.InsertOne(ctx, comp)
	return err
}

func (c CompanysRepo) Update(company company.Company) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := c.collection.FindOneAndUpdate(ctx, bson.M{"_id": company.ID}, bson.M{"$set": company})
	if s.Err() != nil {
		return s.Err()
	}
	return nil
}

func (c CompanysRepo) Get(id string) (*company.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := c.collection.FindOne(ctx, bson.M{"_id": id})
	if s.Err() != nil {
		return nil, s.Err()
	}

	var company *company.Company
	err := s.Decode(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c CompanysRepo) Latests(limit int64) ([]*company.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"creationAt", -1}})
	findOptions.SetLimit(limit)

	cursor, err := c.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "problem finding latest companies")
	}
	defer cursor.Close(ctx)

	var companies []*company.Company
	err = cursor.All(ctx, &companies)
	if err != nil {
		return nil, errors.Wrap(err, "problem getting all companies")
	}
	return companies, nil
}

func toCompanyDB(company company.Company) companyDB {
	hex, _ := primitive.ObjectIDFromHex(company.ID)
	return companyDB{
		ID:   hex,
		Name: company.Name,
		CIF:  company.CIF,
		Address: AddressDB{
			City:   company.City,
			CP:     company.CP,
			Street: company.Street,
		},
		Description:   company.Description,
		Employees:     company.Employees,
		FundationYear: company.FundationYear,
		Web:           company.Web,
		Email:         company.Email,
		Logo:          company.Logo,
	}
}
