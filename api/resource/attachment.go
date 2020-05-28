package resource

import (
	"mikelangelon/m/v2/api/model"
	"mikelangelon/m/v2/api/rest/operation"
	"mikelangelon/m/v2/api/rest/operation/attachments"
	"mikelangelon/m/v2/internal/app/attachmentrepo"

	"github.com/go-openapi/runtime/middleware"
)

type AttachmentResource struct {
	Input map[string]string
	Store attachmentrepo.Store
}

func (r *AttachmentResource) Register(api *operation.CoolappAPI) {
	api.AttachmentsGetAttachmentsHandler = attachments.GetAttachmentsHandlerFunc(r.Handler)
	api.AttachmentsPostAttachmentsHandler = attachments.PostAttachmentsHandlerFunc(r.HandlerPost)
}

func (r *AttachmentResource) Handler(input attachments.GetAttachmentsParams) middleware.Responder {
	list, err := r.Store.ListAttachments(*input.ID)
	if err != nil {
		return attachments.NewGetAttachmentsInternalServerError()
	}
	size := len(list)
	payload := make(model.AttachmentMetaList, size, size)
	for i, m := range list {
		payload[i] = &model.AttachmentMeta{
			Description: m.Description,
			Name:        m.Name,
			OfferID:     m.OfferID,
		}
	}

	return attachments.NewGetAttachmentsOK().WithPayload(payload)
}

func (r *AttachmentResource) HandlerPost(input attachments.PostAttachmentsParams) middleware.Responder {
	r.Input[input.ID] = input.Name

	var desc string
	if input.Description != nil {
		desc = *input.Description
	}
	meta := attachmentrepo.UploadMeta{
		Name:        input.Name,
		Description: desc,
		OfferID:     input.ID,
	}
	_, err := r.Store.Save(input.ID, meta, input.File)
	if err != nil {
		return attachments.NewPostAttachmentsOK()
	}
	return attachments.NewPostAttachmentsOK()
}
