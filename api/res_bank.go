package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResBankService struct {
	client *Client
}

func NewResBankService(c *Client) *ResBankService {
	return &ResBankService{client: c}
}

func (svc *ResBankService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResBankModel, name)
}

func (svc *ResBankService) GetByIds(ids []int64) (*types.ResBanks, error) {
	r := &types.ResBanks{}
	return r, svc.client.getByIds(types.ResBankModel, ids, r)
}

func (svc *ResBankService) GetByName(name string) (*types.ResBanks, error) {
	r := &types.ResBanks{}
	return r, svc.client.getByName(types.ResBankModel, name, r)
}

func (svc *ResBankService) GetByField(field string, value string) (*types.ResBanks, error) {
	r := &types.ResBanks{}
	return r, svc.client.getByField(types.ResBankModel, field, value, r)
}

func (svc *ResBankService) GetAll() (*types.ResBanks, error) {
	r := &types.ResBanks{}
	return r, svc.client.getAll(types.ResBankModel, r)
}

func (svc *ResBankService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResBankModel, fields, relations)
}

func (svc *ResBankService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResBankModel, ids, fields, relations)
}

func (svc *ResBankService) Delete(ids []int64) error {
	return svc.client.delete(types.ResBankModel, ids)
}
