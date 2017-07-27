package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountReportGeneralLedgerService struct {
	client *Client
}

func NewAccountReportGeneralLedgerService(c *Client) *AccountReportGeneralLedgerService {
	return &AccountReportGeneralLedgerService{client: c}
}

func (svc *AccountReportGeneralLedgerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountReportGeneralLedgerModel, name)
}

func (svc *AccountReportGeneralLedgerService) GetByIds(ids []int) (*types.AccountReportGeneralLedgers, error) {
	a := &types.AccountReportGeneralLedgers{}
	return a, svc.client.getByIds(types.AccountReportGeneralLedgerModel, ids, a)
}

func (svc *AccountReportGeneralLedgerService) GetByName(name string) (*types.AccountReportGeneralLedgers, error) {
	a := &types.AccountReportGeneralLedgers{}
	return a, svc.client.getByName(types.AccountReportGeneralLedgerModel, name, a)
}

func (svc *AccountReportGeneralLedgerService) GetAll() (*types.AccountReportGeneralLedgers, error) {
	a := &types.AccountReportGeneralLedgers{}
	return a, svc.client.getAll(types.AccountReportGeneralLedgerModel, a)
}

func (svc *AccountReportGeneralLedgerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountReportGeneralLedgerModel, fields, relations)
}

func (svc *AccountReportGeneralLedgerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountReportGeneralLedgerModel, ids, fields, relations)
}

func (svc *AccountReportGeneralLedgerService) Delete(ids []int) error {
	return svc.client.delete(types.AccountReportGeneralLedgerModel, ids)
}
