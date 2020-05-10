package resource

import (
	"mikelangelon/m/v2/api/model"
	"mikelangelon/m/v2/api/rest/operation"

	"github.com/go-openapi/runtime/middleware"
)

type Resource struct{}

func (r *Resource) Register(api *operation.CoolappAPI) {
	api.GetUsersHandler = operation.GetUsersHandlerFunc(r.Handler)
}

func (r *Resource) Handler(input operation.GetUsersParams) middleware.Responder {
	i := int64(1)
	s := "hola"
	return operation.NewGetUsersOK().WithPayload(
		model.Users{
			&model.User{
				ID:   &i,
				Name: &s,
			},
		},
	)
}
