package api

import (
	"github.com/morezig/go-odoo/types"
)

type SaleLayoutCategoryService struct {
	client *Client
}

func NewSaleLayoutCategoryService(c *Client) *SaleLayoutCategoryService {
	return &SaleLayoutCategoryService{client: c}
}

func (svc *SaleLayoutCategoryService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.SaleLayoutCategoryModel, name)
}

func (svc *SaleLayoutCategoryService) GetByIds(ids []int64) (*types.SaleLayoutCategorys, error) {
	s := &types.SaleLayoutCategorys{}
	return s, svc.client.getByIds(types.SaleLayoutCategoryModel, ids, s)
}

func (svc *SaleLayoutCategoryService) GetByName(name string) (*types.SaleLayoutCategorys, error) {
	s := &types.SaleLayoutCategorys{}
	return s, svc.client.getByName(types.SaleLayoutCategoryModel, name, s)
}

func (svc *SaleLayoutCategoryService) GetByField(field string, value string) (*types.SaleLayoutCategorys, error) {
	s := &types.SaleLayoutCategorys{}
	return s, svc.client.getByField(types.SaleLayoutCategoryModel, field, value, s)
}

func (svc *SaleLayoutCategoryService) GetAll() (*types.SaleLayoutCategorys, error) {
	s := &types.SaleLayoutCategorys{}
	return s, svc.client.getAll(types.SaleLayoutCategoryModel, s)
}

func (svc *SaleLayoutCategoryService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.SaleLayoutCategoryModel, fields, relations)
}

func (svc *SaleLayoutCategoryService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SaleLayoutCategoryModel, ids, fields, relations)
}

func (svc *SaleLayoutCategoryService) Delete(ids []int64) error {
	return svc.client.delete(types.SaleLayoutCategoryModel, ids)
}
