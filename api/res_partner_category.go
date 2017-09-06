package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResPartnerCategoryService struct {
	client *Client
}

func NewResPartnerCategoryService(c *Client) *ResPartnerCategoryService {
	return &ResPartnerCategoryService{client: c}
}

func (svc *ResPartnerCategoryService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResPartnerCategoryModel, name)
}

func (svc *ResPartnerCategoryService) GetByIds(ids []int) (*types.ResPartnerCategorys, error) {
	r := &types.ResPartnerCategorys{}
	return r, svc.client.getByIds(types.ResPartnerCategoryModel, ids, r)
}

func (svc *ResPartnerCategoryService) GetByName(name string) (*types.ResPartnerCategorys, error) {
	r := &types.ResPartnerCategorys{}
	return r, svc.client.getByName(types.ResPartnerCategoryModel, name, r)
}

func (svc *ResPartnerCategoryService) GetByField(field string, value string) (*types.ResPartnerCategorys, error) {
	r := &types.ResPartnerCategorys{}
	return r, svc.client.getByField(types.ResPartnerCategoryModel, field, value, r)
}

func (svc *ResPartnerCategoryService) GetAll() (*types.ResPartnerCategorys, error) {
	r := &types.ResPartnerCategorys{}
	return r, svc.client.getAll(types.ResPartnerCategoryModel, r)
}

func (svc *ResPartnerCategoryService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResPartnerCategoryModel, fields, relations)
}

func (svc *ResPartnerCategoryService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResPartnerCategoryModel, ids, fields, relations)
}

func (svc *ResPartnerCategoryService) Delete(ids []int) error {
	return svc.client.delete(types.ResPartnerCategoryModel, ids)
}
