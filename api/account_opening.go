package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountOpeningService struct {
	client *Client
}

func NewAccountOpeningService(c *Client) *AccountOpeningService {
	return &AccountOpeningService{client: c}
}

func (svc *AccountOpeningService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountOpeningModel, name)
}

func (svc *AccountOpeningService) GetByIds(ids []int64) (*types.AccountOpenings, error) {
	a := &types.AccountOpenings{}
	return a, svc.client.getByIds(types.AccountOpeningModel, ids, a)
}

func (svc *AccountOpeningService) GetByName(name string) (*types.AccountOpenings, error) {
	a := &types.AccountOpenings{}
	return a, svc.client.getByName(types.AccountOpeningModel, name, a)
}

func (svc *AccountOpeningService) GetByField(field string, value string) (*types.AccountOpenings, error) {
	a := &types.AccountOpenings{}
	return a, svc.client.getByField(types.AccountOpeningModel, field, value, a)
}

func (svc *AccountOpeningService) GetAll() (*types.AccountOpenings, error) {
	a := &types.AccountOpenings{}
	return a, svc.client.getAll(types.AccountOpeningModel, a)
}

func (svc *AccountOpeningService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountOpeningModel, fields, relations)
}

func (svc *AccountOpeningService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountOpeningModel, ids, fields, relations)
}

func (svc *AccountOpeningService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountOpeningModel, ids)
}
