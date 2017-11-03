package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResPartnerIndustryService struct {
	client *Client
}

func NewResPartnerIndustryService(c *Client) *ResPartnerIndustryService {
	return &ResPartnerIndustryService{client: c}
}

func (svc *ResPartnerIndustryService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResPartnerIndustryModel, name)
}

func (svc *ResPartnerIndustryService) GetByIds(ids []int64) (*types.ResPartnerIndustrys, error) {
	r := &types.ResPartnerIndustrys{}
	return r, svc.client.getByIds(types.ResPartnerIndustryModel, ids, r)
}

func (svc *ResPartnerIndustryService) GetByName(name string) (*types.ResPartnerIndustrys, error) {
	r := &types.ResPartnerIndustrys{}
	return r, svc.client.getByName(types.ResPartnerIndustryModel, name, r)
}

func (svc *ResPartnerIndustryService) GetByField(field string, value string) (*types.ResPartnerIndustrys, error) {
	r := &types.ResPartnerIndustrys{}
	return r, svc.client.getByField(types.ResPartnerIndustryModel, field, value, r)
}

func (svc *ResPartnerIndustryService) GetAll() (*types.ResPartnerIndustrys, error) {
	r := &types.ResPartnerIndustrys{}
	return r, svc.client.getAll(types.ResPartnerIndustryModel, r)
}

func (svc *ResPartnerIndustryService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResPartnerIndustryModel, fields, relations)
}

func (svc *ResPartnerIndustryService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResPartnerIndustryModel, ids, fields, relations)
}

func (svc *ResPartnerIndustryService) Delete(ids []int64) error {
	return svc.client.delete(types.ResPartnerIndustryModel, ids)
}
