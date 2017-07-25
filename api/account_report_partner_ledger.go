package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountReportPartnerLedgerService struct {
	client *Client
}

func NewAccountReportPartnerLedgerService(c *Client) *AccountReportPartnerLedgerService {
	return &AccountReportPartnerLedgerService{client: c}
}

func (svc *AccountReportPartnerLedgerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountReportPartnerLedgerModel, name)
}

func (svc *AccountReportPartnerLedgerService) GetByIds(ids []int) (*types.AccountReportPartnerLedgers, error) {
	a := &types.AccountReportPartnerLedgers{}
	return a, svc.client.getByIds(types.AccountReportPartnerLedgerModel, ids, a)
}

func (svc *AccountReportPartnerLedgerService) GetByName(name string) (*types.AccountReportPartnerLedgers, error) {
	a := &types.AccountReportPartnerLedgers{}
	return a, svc.client.getByName(types.AccountReportPartnerLedgerModel, name, a)
}

func (svc *AccountReportPartnerLedgerService) GetAll() (*types.AccountReportPartnerLedgers, error) {
	a := &types.AccountReportPartnerLedgers{}
	return a, svc.client.getAll(types.AccountReportPartnerLedgerModel, a)
}

func (svc *AccountReportPartnerLedgerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountReportPartnerLedgerModel, fields, relations)
}

func (svc *AccountReportPartnerLedgerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountReportPartnerLedgerModel, ids, fields, relations)
}

func (svc *AccountReportPartnerLedgerService) Delete(ids []int) error {
	return svc.client.delete(types.AccountReportPartnerLedgerModel, ids)
}
