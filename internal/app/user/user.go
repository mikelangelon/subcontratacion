package user

import "errors"

var errUserNotValid = errors.New("user not valid")

type User struct {
	User     string
	Password string
}

type userService struct {
	redis Redis
}

func New(redis Redis) userService {
	return userService{
		redis: redis,
	}
}

func (u userService) Save(user User) error {
	return u.redis.SetPair(user.User, user.Password)
}

func (u userService) Login(user User) error {
	val, err := u.redis.GetPair(user.User)
	if err != nil {
		return errUserNotValid
	}
	if val != user.Password {
		return errUserNotValid
	}
	return nil
}

type Redis interface {
	SetPair(key, value string) error
	GetPair(key string) (string, error)
}
