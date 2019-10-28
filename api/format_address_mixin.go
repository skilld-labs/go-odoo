package api

import (
	"github.com/morezig/go-odoo/types"
)

type FormatAddressMixinService struct {
	client *Client
}

func NewFormatAddressMixinService(c *Client) *FormatAddressMixinService {
	return &FormatAddressMixinService{client: c}
}

func (svc *FormatAddressMixinService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.FormatAddressMixinModel, name)
}

func (svc *FormatAddressMixinService) GetByIds(ids []int64) (*types.FormatAddressMixins, error) {
	f := &types.FormatAddressMixins{}
	return f, svc.client.getByIds(types.FormatAddressMixinModel, ids, f)
}

func (svc *FormatAddressMixinService) GetByName(name string) (*types.FormatAddressMixins, error) {
	f := &types.FormatAddressMixins{}
	return f, svc.client.getByName(types.FormatAddressMixinModel, name, f)
}

func (svc *FormatAddressMixinService) GetByField(field string, value string) (*types.FormatAddressMixins, error) {
	f := &types.FormatAddressMixins{}
	return f, svc.client.getByField(types.FormatAddressMixinModel, field, value, f)
}

func (svc *FormatAddressMixinService) GetAll() (*types.FormatAddressMixins, error) {
	f := &types.FormatAddressMixins{}
	return f, svc.client.getAll(types.FormatAddressMixinModel, f)
}

func (svc *FormatAddressMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.FormatAddressMixinModel, fields, relations)
}

func (svc *FormatAddressMixinService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.FormatAddressMixinModel, ids, fields, relations)
}

func (svc *FormatAddressMixinService) Delete(ids []int64) error {
	return svc.client.delete(types.FormatAddressMixinModel, ids)
}
