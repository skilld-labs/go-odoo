package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResCompanyService struct {
	client *Client
}

func NewResCompanyService(c *Client) *ResCompanyService {
	return &ResCompanyService{client: c}
}

func (svc *ResCompanyService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResCompanyModel, name)
}

func (svc *ResCompanyService) GetByIds(ids []int) (*types.ResCompanys, error) {
	r := &types.ResCompanys{}
	return r, svc.client.getByIds(types.ResCompanyModel, ids, r)
}

func (svc *ResCompanyService) GetByName(name string) (*types.ResCompanys, error) {
	r := &types.ResCompanys{}
	return r, svc.client.getByName(types.ResCompanyModel, name, r)
}

func (svc *ResCompanyService) GetByField(field string, value string) (*types.ResCompanys, error) {
	r := &types.ResCompanys{}
	return r, svc.client.getByField(types.ResCompanyModel, field, value, r)
}

func (svc *ResCompanyService) GetAll() (*types.ResCompanys, error) {
	r := &types.ResCompanys{}
	return r, svc.client.getAll(types.ResCompanyModel, r)
}

func (svc *ResCompanyService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResCompanyModel, fields, relations)
}

func (svc *ResCompanyService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResCompanyModel, ids, fields, relations)
}

func (svc *ResCompanyService) Delete(ids []int) error {
	return svc.client.delete(types.ResCompanyModel, ids)
}
