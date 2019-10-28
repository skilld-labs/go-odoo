package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountAccountTagService struct {
	client *Client
}

func NewAccountAccountTagService(c *Client) *AccountAccountTagService {
	return &AccountAccountTagService{client: c}
}

func (svc *AccountAccountTagService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountAccountTagModel, name)
}

func (svc *AccountAccountTagService) GetByIds(ids []int64) (*types.AccountAccountTags, error) {
	a := &types.AccountAccountTags{}
	return a, svc.client.getByIds(types.AccountAccountTagModel, ids, a)
}

func (svc *AccountAccountTagService) GetByName(name string) (*types.AccountAccountTags, error) {
	a := &types.AccountAccountTags{}
	return a, svc.client.getByName(types.AccountAccountTagModel, name, a)
}

func (svc *AccountAccountTagService) GetByField(field string, value string) (*types.AccountAccountTags, error) {
	a := &types.AccountAccountTags{}
	return a, svc.client.getByField(types.AccountAccountTagModel, field, value, a)
}

func (svc *AccountAccountTagService) GetAll() (*types.AccountAccountTags, error) {
	a := &types.AccountAccountTags{}
	return a, svc.client.getAll(types.AccountAccountTagModel, a)
}

func (svc *AccountAccountTagService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountAccountTagModel, fields, relations)
}

func (svc *AccountAccountTagService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAccountTagModel, ids, fields, relations)
}

func (svc *AccountAccountTagService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountAccountTagModel, ids)
}
