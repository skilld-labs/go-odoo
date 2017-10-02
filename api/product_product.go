package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductProductService struct {
	client *Client
}

func NewProductProductService(c *Client) *ProductProductService {
	return &ProductProductService{client: c}
}

func (svc *ProductProductService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProductProductModel, name)
}

func (svc *ProductProductService) GetByIds(ids []int64) (*types.ProductProducts, error) {
	p := &types.ProductProducts{}
	return p, svc.client.getByIds(types.ProductProductModel, ids, p)
}

func (svc *ProductProductService) GetByName(name string) (*types.ProductProducts, error) {
	p := &types.ProductProducts{}
	return p, svc.client.getByName(types.ProductProductModel, name, p)
}

func (svc *ProductProductService) GetByField(field string, value string) (*types.ProductProducts, error) {
	p := &types.ProductProducts{}
	return p, svc.client.getByField(types.ProductProductModel, field, value, p)
}

func (svc *ProductProductService) GetAll() (*types.ProductProducts, error) {
	p := &types.ProductProducts{}
	return p, svc.client.getAll(types.ProductProductModel, p)
}

func (svc *ProductProductService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProductProductModel, fields, relations)
}

func (svc *ProductProductService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductProductModel, ids, fields, relations)
}

func (svc *ProductProductService) Delete(ids []int64) error {
	return svc.client.delete(types.ProductProductModel, ids)
}
