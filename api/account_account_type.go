package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountAccountTypeService struct {
	client *Client
}

func NewAccountAccountTypeService(c *Client) *AccountAccountTypeService {
	return &AccountAccountTypeService{client: c}
}

func (svc *AccountAccountTypeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountAccountTypeModel, name)
}

func (svc *AccountAccountTypeService) GetByIds(ids []int) (*types.AccountAccountTypes, error) {
	a := &types.AccountAccountTypes{}
	return a, svc.client.getByIds(types.AccountAccountTypeModel, ids, a)
}

func (svc *AccountAccountTypeService) GetByName(name string) (*types.AccountAccountTypes, error) {
	a := &types.AccountAccountTypes{}
	return a, svc.client.getByName(types.AccountAccountTypeModel, name, a)
}

func (svc *AccountAccountTypeService) GetByField(field string, value string) (*types.AccountAccountTypes, error) {
	a := &types.AccountAccountTypes{}
	return a, svc.client.getByName(types.AccountAccountTypeModel, field, value, a)
}

func (svc *AccountAccountTypeService) GetAll() (*types.AccountAccountTypes, error) {
	a := &types.AccountAccountTypes{}
	return a, svc.client.getAll(types.AccountAccountTypeModel, a)
}

func (svc *AccountAccountTypeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountAccountTypeModel, fields, relations)
}

func (svc *AccountAccountTypeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAccountTypeModel, ids, fields, relations)
}

func (svc *AccountAccountTypeService) Delete(ids []int) error {
	return svc.client.delete(types.AccountAccountTypeModel, ids)
}
