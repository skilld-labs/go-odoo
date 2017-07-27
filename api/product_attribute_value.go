package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductAttributeValueService struct {
	client *Client
}

func NewProductAttributeValueService(c *Client) *ProductAttributeValueService {
	return &ProductAttributeValueService{client: c}
}

func (svc *ProductAttributeValueService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductAttributeValueModel, name)
}

func (svc *ProductAttributeValueService) GetByIds(ids []int) (*types.ProductAttributeValues, error) {
	p := &types.ProductAttributeValues{}
	return p, svc.client.getByIds(types.ProductAttributeValueModel, ids, p)
}

func (svc *ProductAttributeValueService) GetByName(name string) (*types.ProductAttributeValues, error) {
	p := &types.ProductAttributeValues{}
	return p, svc.client.getByName(types.ProductAttributeValueModel, name, p)
}

func (svc *ProductAttributeValueService) GetAll() (*types.ProductAttributeValues, error) {
	p := &types.ProductAttributeValues{}
	return p, svc.client.getAll(types.ProductAttributeValueModel, p)
}

func (svc *ProductAttributeValueService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductAttributeValueModel, fields, relations)
}

func (svc *ProductAttributeValueService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductAttributeValueModel, ids, fields, relations)
}

func (svc *ProductAttributeValueService) Delete(ids []int) error {
	return svc.client.delete(types.ProductAttributeValueModel, ids)
}
