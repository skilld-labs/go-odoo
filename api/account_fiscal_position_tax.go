package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountFiscalPositionTaxService struct {
	client *Client
}

func NewAccountFiscalPositionTaxService(c *Client) *AccountFiscalPositionTaxService {
	return &AccountFiscalPositionTaxService{client: c}
}

func (svc *AccountFiscalPositionTaxService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountFiscalPositionTaxModel, name)
}

func (svc *AccountFiscalPositionTaxService) GetByIds(ids []int) (*types.AccountFiscalPositionTaxs, error) {
	a := &types.AccountFiscalPositionTaxs{}
	return a, svc.client.getByIds(types.AccountFiscalPositionTaxModel, ids, a)
}

func (svc *AccountFiscalPositionTaxService) GetByName(name string) (*types.AccountFiscalPositionTaxs, error) {
	a := &types.AccountFiscalPositionTaxs{}
	return a, svc.client.getByName(types.AccountFiscalPositionTaxModel, name, a)
}

func (svc *AccountFiscalPositionTaxService) GetAll() (*types.AccountFiscalPositionTaxs, error) {
	a := &types.AccountFiscalPositionTaxs{}
	return a, svc.client.getAll(types.AccountFiscalPositionTaxModel, a)
}

func (svc *AccountFiscalPositionTaxService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountFiscalPositionTaxModel, fields, relations)
}

func (svc *AccountFiscalPositionTaxService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFiscalPositionTaxModel, ids, fields, relations)
}

func (svc *AccountFiscalPositionTaxService) Delete(ids []int) error {
	return svc.client.delete(types.AccountFiscalPositionTaxModel, ids)
}
