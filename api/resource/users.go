package resource

import (
	"fmt"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/api/rest/operation/users"
	"mikelangelon/m/v2/internal/app/user"

	"github.com/go-openapi/runtime/middleware"
)

type userResource struct {
	UserService userService
}

type userService interface {
	Save(User user.User) error
	Login(User user.User) error
}

func (u *userResource) Register(api *operation.CoolappAPI) {
	api.UsersPostUsersHandler = users.PostUsersHandlerFunc(u.Handler)
}

func (u *userResource) Handler(input users.PostUsersParams) middleware.Responder {
	user := user.User{
		User:     *input.User.Email,
		Password: *input.User.Password,
	}
	err := u.UserService.Save(user)
	if err != nil {
		fmt.Print(err)
		return users.NewPostUsersInternalServerError()
	}
	return users.NewPostUsersOK()
}
