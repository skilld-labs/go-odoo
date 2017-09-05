package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountCashboxLineService struct {
	client *Client
}

func NewAccountCashboxLineService(c *Client) *AccountCashboxLineService {
	return &AccountCashboxLineService{client: c}
}

func (svc *AccountCashboxLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountCashboxLineModel, name)
}

func (svc *AccountCashboxLineService) GetByIds(ids []int) (*types.AccountCashboxLines, error) {
	a := &types.AccountCashboxLines{}
	return a, svc.client.getByIds(types.AccountCashboxLineModel, ids, a)
}

func (svc *AccountCashboxLineService) GetByName(name string) (*types.AccountCashboxLines, error) {
	a := &types.AccountCashboxLines{}
	return a, svc.client.getByName(types.AccountCashboxLineModel, name, a)
}

func (svc *AccountCashboxLineService) GetByField(field string, value string) (*types.AccountCashboxLines, error) {
	a := &types.AccountCashboxLines{}
	return a, svc.client.getByName(types.AccountCashboxLineModel, field, value, a)
}

func (svc *AccountCashboxLineService) GetAll() (*types.AccountCashboxLines, error) {
	a := &types.AccountCashboxLines{}
	return a, svc.client.getAll(types.AccountCashboxLineModel, a)
}

func (svc *AccountCashboxLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountCashboxLineModel, fields, relations)
}

func (svc *AccountCashboxLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountCashboxLineModel, ids, fields, relations)
}

func (svc *AccountCashboxLineService) Delete(ids []int) error {
	return svc.client.delete(types.AccountCashboxLineModel, ids)
}
