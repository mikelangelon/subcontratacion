package resource

import (
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/internal/app/attachmentrepo"
	"mikelangelon/m/v2/internal/app/company"
	"mikelangelon/m/v2/internal/app/user"
)

type Dependencies struct {
	Store attachmentrepo.Store
	Redis user.Redis
}

func Register(api *operation.CoolappAPI, dependencies Dependencies) {
	attachment := &AttachmentResource{
		Store: dependencies.Store,
		Input: map[string]string{},
	}
	attachment.Register(api)

	userService := user.New(dependencies.Redis)
	user := &userResource{
		UserService: userService,
	}
	user.Register(api)

	session := &sessionResource{
		UserService: userService,
	}
	session.Register(api)

	comResource := &companyResource{service: company.New()}
	comResource.Register(api)

	conResource := &contactResource{service: company.New()}
	conResource.Register(api)
}
