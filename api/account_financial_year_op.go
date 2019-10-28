package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountFinancialYearOpService struct {
	client *Client
}

func NewAccountFinancialYearOpService(c *Client) *AccountFinancialYearOpService {
	return &AccountFinancialYearOpService{client: c}
}

func (svc *AccountFinancialYearOpService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountFinancialYearOpModel, name)
}

func (svc *AccountFinancialYearOpService) GetByIds(ids []int64) (*types.AccountFinancialYearOps, error) {
	a := &types.AccountFinancialYearOps{}
	return a, svc.client.getByIds(types.AccountFinancialYearOpModel, ids, a)
}

func (svc *AccountFinancialYearOpService) GetByName(name string) (*types.AccountFinancialYearOps, error) {
	a := &types.AccountFinancialYearOps{}
	return a, svc.client.getByName(types.AccountFinancialYearOpModel, name, a)
}

func (svc *AccountFinancialYearOpService) GetByField(field string, value string) (*types.AccountFinancialYearOps, error) {
	a := &types.AccountFinancialYearOps{}
	return a, svc.client.getByField(types.AccountFinancialYearOpModel, field, value, a)
}

func (svc *AccountFinancialYearOpService) GetAll() (*types.AccountFinancialYearOps, error) {
	a := &types.AccountFinancialYearOps{}
	return a, svc.client.getAll(types.AccountFinancialYearOpModel, a)
}

func (svc *AccountFinancialYearOpService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountFinancialYearOpModel, fields, relations)
}

func (svc *AccountFinancialYearOpService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFinancialYearOpModel, ids, fields, relations)
}

func (svc *AccountFinancialYearOpService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountFinancialYearOpModel, ids)
}
