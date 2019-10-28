package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountReportGeneralLedgerService struct {
	client *Client
}

func NewAccountReportGeneralLedgerService(c *Client) *AccountReportGeneralLedgerService {
	return &AccountReportGeneralLedgerService{client: c}
}

func (svc *AccountReportGeneralLedgerService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountReportGeneralLedgerModel, name)
}

func (svc *AccountReportGeneralLedgerService) GetByIds(ids []int64) (*types.AccountReportGeneralLedgers, error) {
	a := &types.AccountReportGeneralLedgers{}
	return a, svc.client.getByIds(types.AccountReportGeneralLedgerModel, ids, a)
}

func (svc *AccountReportGeneralLedgerService) GetByName(name string) (*types.AccountReportGeneralLedgers, error) {
	a := &types.AccountReportGeneralLedgers{}
	return a, svc.client.getByName(types.AccountReportGeneralLedgerModel, name, a)
}

func (svc *AccountReportGeneralLedgerService) GetByField(field string, value string) (*types.AccountReportGeneralLedgers, error) {
	a := &types.AccountReportGeneralLedgers{}
	return a, svc.client.getByField(types.AccountReportGeneralLedgerModel, field, value, a)
}

func (svc *AccountReportGeneralLedgerService) GetAll() (*types.AccountReportGeneralLedgers, error) {
	a := &types.AccountReportGeneralLedgers{}
	return a, svc.client.getAll(types.AccountReportGeneralLedgerModel, a)
}

func (svc *AccountReportGeneralLedgerService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountReportGeneralLedgerModel, fields, relations)
}

func (svc *AccountReportGeneralLedgerService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountReportGeneralLedgerModel, ids, fields, relations)
}

func (svc *AccountReportGeneralLedgerService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountReportGeneralLedgerModel, ids)
}
