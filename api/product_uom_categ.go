package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductUomCategService struct {
	client *Client
}

func NewProductUomCategService(c *Client) *ProductUomCategService {
	return &ProductUomCategService{client: c}
}

func (svc *ProductUomCategService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductUomCategModel, name)
}

func (svc *ProductUomCategService) GetByIds(ids []int) (*types.ProductUomCategs, error) {
	p := &types.ProductUomCategs{}
	return p, svc.client.getByIds(types.ProductUomCategModel, ids, p)
}

func (svc *ProductUomCategService) GetByName(name string) (*types.ProductUomCategs, error) {
	p := &types.ProductUomCategs{}
	return p, svc.client.getByName(types.ProductUomCategModel, name, p)
}

func (svc *ProductUomCategService) GetByField(field string, value string) (*types.ProductUomCategs, error) {
	p := &types.ProductUomCategs{}
	return p, svc.client.getByField(types.ProductUomCategModel, field, value, p)
}

func (svc *ProductUomCategService) GetAll() (*types.ProductUomCategs, error) {
	p := &types.ProductUomCategs{}
	return p, svc.client.getAll(types.ProductUomCategModel, p)
}

func (svc *ProductUomCategService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductUomCategModel, fields, relations)
}

func (svc *ProductUomCategService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductUomCategModel, ids, fields, relations)
}

func (svc *ProductUomCategService) Delete(ids []int) error {
	return svc.client.delete(types.ProductUomCategModel, ids)
}
