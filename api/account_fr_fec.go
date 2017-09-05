package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountFrFecService struct {
	client *Client
}

func NewAccountFrFecService(c *Client) *AccountFrFecService {
	return &AccountFrFecService{client: c}
}

func (svc *AccountFrFecService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountFrFecModel, name)
}

func (svc *AccountFrFecService) GetByIds(ids []int) (*types.AccountFrFecs, error) {
	a := &types.AccountFrFecs{}
	return a, svc.client.getByIds(types.AccountFrFecModel, ids, a)
}

func (svc *AccountFrFecService) GetByName(name string) (*types.AccountFrFecs, error) {
	a := &types.AccountFrFecs{}
	return a, svc.client.getByName(types.AccountFrFecModel, name, a)
}

func (svc *AccountFrFecService) GetByField(field string, value string) (*types.AccountFrFecs, error) {
	a := &types.AccountFrFecs{}
	return a, svc.client.getByField(types.AccountFrFecModel, field, value, a)
}

func (svc *AccountFrFecService) GetAll() (*types.AccountFrFecs, error) {
	a := &types.AccountFrFecs{}
	return a, svc.client.getAll(types.AccountFrFecModel, a)
}

func (svc *AccountFrFecService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountFrFecModel, fields, relations)
}

func (svc *AccountFrFecService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFrFecModel, ids, fields, relations)
}

func (svc *AccountFrFecService) Delete(ids []int) error {
	return svc.client.delete(types.AccountFrFecModel, ids)
}
