package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductCategoryService struct {
	client *Client
}

func NewProductCategoryService(c *Client) *ProductCategoryService {
	return &ProductCategoryService{client: c}
}

func (svc *ProductCategoryService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductCategoryModel, name)
}

func (svc *ProductCategoryService) GetByIds(ids []int) (*types.ProductCategorys, error) {
	p := &types.ProductCategorys{}
	return p, svc.client.getByIds(types.ProductCategoryModel, ids, p)
}

func (svc *ProductCategoryService) GetByName(name string) (*types.ProductCategorys, error) {
	p := &types.ProductCategorys{}
	return p, svc.client.getByName(types.ProductCategoryModel, name, p)
}

func (svc *ProductCategoryService) GetAll() (*types.ProductCategorys, error) {
	p := &types.ProductCategorys{}
	return p, svc.client.getAll(types.ProductCategoryModel, p)
}

func (svc *ProductCategoryService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductCategoryModel, fields, relations)
}

func (svc *ProductCategoryService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductCategoryModel, ids, fields, relations)
}

func (svc *ProductCategoryService) Delete(ids []int) error {
	return svc.client.delete(types.ProductCategoryModel, ids)
}
