package db

import (
	"context"
	"mikelangelon/m/v2/internal/app/user"
	"mikelangelon/m/v2/internal/pkg/encrypt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	collection *mongo.Collection
	encryptor  encrypt.Encryptor
}

func NewUserRepo(db *mongo.Database) UserRepo {
	return UserRepo{
		collection: db.Collection("users"),
		encryptor:  encrypt.New("test"),
	}
}

func (u UserRepo) Save(user user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.Password = u.encryptor.Encrypt(user.Password)
	_, err := u.collection.InsertOne(ctx, toUserDB(user))
	return err
}

func (u UserRepo) Exists(user user.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := u.collection.FindOne(ctx, bson.M{"_id": user.User})
	if s.Err() != nil {
		switch s.Err() {
		case mongo.ErrNoDocuments:
			return false, nil
		default:
			return false, s.Err()
		}
	}
	var db userDB
	err := s.Decode(&db)
	if err != nil {
		logrus.WithError(err).Error("error decoding user")
	}

	if u.encryptor.Decrypt(db.Password) != user.Password {
		return false, nil
	}
	return true, nil

}

type userDB struct {
	User     string `bson:"_id"`
	Password string `bson:"password"`
}

func toUserDB(user user.User) userDB {
	return userDB{
		User:     user.User,
		Password: user.Password,
	}
}
