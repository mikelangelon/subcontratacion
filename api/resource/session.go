package resource

import (
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/api/rest/operation/sessions"
	"mikelangelon/m/v2/internal/app/user"

	"github.com/go-openapi/runtime/middleware"
)

type sessionResource struct {
	UserService userService
}

func (s *sessionResource) Register(api *operation.CoolappAPI) {
	api.SessionsPostSessionsHandler = sessions.PostSessionsHandlerFunc(s.Handler)
}

func (s *sessionResource) Handler(input sessions.PostSessionsParams) middleware.Responder {
	err := s.UserService.Login(user.User{
		User:     *input.User.Email,
		Password: *input.User.Password,
	})
	if err != nil {
		return sessions.NewPostSessionsInternalServerError()
	}
	return sessions.NewPostSessionsOK()
}
