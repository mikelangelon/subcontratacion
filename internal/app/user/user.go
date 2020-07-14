package user

import (
	"errors"
)

var ErrUserNotValid = errors.New("user not valid")

type User struct {
	User     string
	Password string
}

type UserService struct {
	repo repo
}

type repo interface {
	Save(user User) error
	Exists(user User) (bool, error)
}

func New(repo repo) UserService {
	return UserService{
		repo: repo,
	}
}

func (u UserService) Register(user User) error {
	return u.repo.Save(user)
}

func (u UserService) Auth(username, password string) error {
	val, err := u.repo.Exists(User{
		User:     username,
		Password: password,
	})
	if err != nil {
		return err
	}
	if !val {
		return ErrUserNotValid
	}
	return nil
}
