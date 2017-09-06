package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResPartnerService struct {
	client *Client
}

func NewResPartnerService(c *Client) *ResPartnerService {
	return &ResPartnerService{client: c}
}

func (svc *ResPartnerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResPartnerModel, name)
}

func (svc *ResPartnerService) GetByIds(ids []int) (*types.ResPartners, error) {
	r := &types.ResPartners{}
	return r, svc.client.getByIds(types.ResPartnerModel, ids, r)
}

func (svc *ResPartnerService) GetByName(name string) (*types.ResPartners, error) {
	r := &types.ResPartners{}
	return r, svc.client.getByName(types.ResPartnerModel, name, r)
}

func (svc *ResPartnerService) GetByField(field string, value string) (*types.ResPartners, error) {
	r := &types.ResPartners{}
	return r, svc.client.getByField(types.ResPartnerModel, field, value, r)
}

func (svc *ResPartnerService) GetAll() (*types.ResPartners, error) {
	r := &types.ResPartners{}
	return r, svc.client.getAll(types.ResPartnerModel, r)
}

func (svc *ResPartnerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResPartnerModel, fields, relations)
}

func (svc *ResPartnerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResPartnerModel, ids, fields, relations)
}

func (svc *ResPartnerService) Delete(ids []int) error {
	return svc.client.delete(types.ResPartnerModel, ids)
}
