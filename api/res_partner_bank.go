package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResPartnerBankService struct {
	client *Client
}

func NewResPartnerBankService(c *Client) *ResPartnerBankService {
	return &ResPartnerBankService{client: c}
}

func (svc *ResPartnerBankService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResPartnerBankModel, name)
}

func (svc *ResPartnerBankService) GetByIds(ids []int64) (*types.ResPartnerBanks, error) {
	r := &types.ResPartnerBanks{}
	return r, svc.client.getByIds(types.ResPartnerBankModel, ids, r)
}

func (svc *ResPartnerBankService) GetByName(name string) (*types.ResPartnerBanks, error) {
	r := &types.ResPartnerBanks{}
	return r, svc.client.getByName(types.ResPartnerBankModel, name, r)
}

func (svc *ResPartnerBankService) GetByField(field string, value string) (*types.ResPartnerBanks, error) {
	r := &types.ResPartnerBanks{}
	return r, svc.client.getByField(types.ResPartnerBankModel, field, value, r)
}

func (svc *ResPartnerBankService) GetAll() (*types.ResPartnerBanks, error) {
	r := &types.ResPartnerBanks{}
	return r, svc.client.getAll(types.ResPartnerBankModel, r)
}

func (svc *ResPartnerBankService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResPartnerBankModel, fields, relations)
}

func (svc *ResPartnerBankService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResPartnerBankModel, ids, fields, relations)
}

func (svc *ResPartnerBankService) Delete(ids []int64) error {
	return svc.client.delete(types.ResPartnerBankModel, ids)
}
