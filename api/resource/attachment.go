package resource

import (
	"mikelangelon/m/v2/api/model"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/api/rest/operation/attachments"

	"github.com/go-openapi/runtime/middleware"
)

type ResourceA struct {
	Input map[string]string
}

func (r *ResourceA) Register(api *operation.CoolappAPI) {
	api.AttachmentsGetAttachmentsHandler = attachments.GetAttachmentsHandlerFunc(r.Handler)
	api.AttachmentsPostAttachmentsHandler = attachments.PostAttachmentsHandlerFunc(r.HandlerPost)

}

func (r *ResourceA) Handler(input attachments.GetAttachmentsParams) middleware.Responder {

	return attachments.NewGetAttachmentsOK().WithPayload(model.AttachmentMetaList{
		{
			Description: "test",
			Name:        r.Input[*input.OfferID],
			OfferID:     *input.OfferID,
		},
	})
}

func (r *ResourceA) HandlerPost(input attachments.PostAttachmentsParams) middleware.Responder {
	r.Input[input.OfferID] = input.Name
	return attachments.NewPostAttachmentsOK()
}
