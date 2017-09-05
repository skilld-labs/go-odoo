package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type EmailTemplatePreviewService struct {
	client *Client
}

func NewEmailTemplatePreviewService(c *Client) *EmailTemplatePreviewService {
	return &EmailTemplatePreviewService{client: c}
}

func (svc *EmailTemplatePreviewService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.EmailTemplatePreviewModel, name)
}

func (svc *EmailTemplatePreviewService) GetByIds(ids []int) (*types.EmailTemplatePreviews, error) {
	e := &types.EmailTemplatePreviews{}
	return e, svc.client.getByIds(types.EmailTemplatePreviewModel, ids, e)
}

func (svc *EmailTemplatePreviewService) GetByName(name string) (*types.EmailTemplatePreviews, error) {
	e := &types.EmailTemplatePreviews{}
	return e, svc.client.getByName(types.EmailTemplatePreviewModel, name, e)
}

func (svc *EmailTemplatePreviewService) GetByField(field string, value string) (*types.EmailTemplatePreviews, error) {
	e := &types.EmailTemplatePreviews{}
	return e, svc.client.getByName(types.EmailTemplatePreviewModel, field, value, e)
}

func (svc *EmailTemplatePreviewService) GetAll() (*types.EmailTemplatePreviews, error) {
	e := &types.EmailTemplatePreviews{}
	return e, svc.client.getAll(types.EmailTemplatePreviewModel, e)
}

func (svc *EmailTemplatePreviewService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.EmailTemplatePreviewModel, fields, relations)
}

func (svc *EmailTemplatePreviewService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.EmailTemplatePreviewModel, ids, fields, relations)
}

func (svc *EmailTemplatePreviewService) Delete(ids []int) error {
	return svc.client.delete(types.EmailTemplatePreviewModel, ids)
}
