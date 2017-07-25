package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductUomService struct {
	client *Client
}

func NewProductUomService(c *Client) *ProductUomService {
	return &ProductUomService{client: c}
}

func (svc *ProductUomService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductUomModel, name)
}

func (svc *ProductUomService) GetByIds(ids []int) (*types.ProductUoms, error) {
	p := &types.ProductUoms{}
	return p, svc.client.getByIds(types.ProductUomModel, ids, p)
}

func (svc *ProductUomService) GetByName(name string) (*types.ProductUoms, error) {
	p := &types.ProductUoms{}
	return p, svc.client.getByName(types.ProductUomModel, name, p)
}

func (svc *ProductUomService) GetAll() (*types.ProductUoms, error) {
	p := &types.ProductUoms{}
	return p, svc.client.getAll(types.ProductUomModel, p)
}

func (svc *ProductUomService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductUomModel, fields, relations)
}

func (svc *ProductUomService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductUomModel, ids, fields, relations)
}

func (svc *ProductUomService) Delete(ids []int) error {
	return svc.client.delete(types.ProductUomModel, ids)
}
