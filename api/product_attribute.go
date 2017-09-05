package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductAttributeService struct {
	client *Client
}

func NewProductAttributeService(c *Client) *ProductAttributeService {
	return &ProductAttributeService{client: c}
}

func (svc *ProductAttributeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductAttributeModel, name)
}

func (svc *ProductAttributeService) GetByIds(ids []int) (*types.ProductAttributes, error) {
	p := &types.ProductAttributes{}
	return p, svc.client.getByIds(types.ProductAttributeModel, ids, p)
}

func (svc *ProductAttributeService) GetByName(name string) (*types.ProductAttributes, error) {
	p := &types.ProductAttributes{}
	return p, svc.client.getByName(types.ProductAttributeModel, name, p)
}

func (svc *ProductAttributeService) GetByField(field string, value string) (*types.ProductAttributes, error) {
	p := &types.ProductAttributes{}
	return p, svc.client.getByName(types.ProductAttributeModel, field, value, p)
}

func (svc *ProductAttributeService) GetAll() (*types.ProductAttributes, error) {
	p := &types.ProductAttributes{}
	return p, svc.client.getAll(types.ProductAttributeModel, p)
}

func (svc *ProductAttributeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductAttributeModel, fields, relations)
}

func (svc *ProductAttributeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductAttributeModel, ids, fields, relations)
}

func (svc *ProductAttributeService) Delete(ids []int) error {
	return svc.client.delete(types.ProductAttributeModel, ids)
}
