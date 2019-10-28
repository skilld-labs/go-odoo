package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProductAttributePriceService struct {
	client *Client
}

func NewProductAttributePriceService(c *Client) *ProductAttributePriceService {
	return &ProductAttributePriceService{client: c}
}

func (svc *ProductAttributePriceService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProductAttributePriceModel, name)
}

func (svc *ProductAttributePriceService) GetByIds(ids []int64) (*types.ProductAttributePrices, error) {
	p := &types.ProductAttributePrices{}
	return p, svc.client.getByIds(types.ProductAttributePriceModel, ids, p)
}

func (svc *ProductAttributePriceService) GetByName(name string) (*types.ProductAttributePrices, error) {
	p := &types.ProductAttributePrices{}
	return p, svc.client.getByName(types.ProductAttributePriceModel, name, p)
}

func (svc *ProductAttributePriceService) GetByField(field string, value string) (*types.ProductAttributePrices, error) {
	p := &types.ProductAttributePrices{}
	return p, svc.client.getByField(types.ProductAttributePriceModel, field, value, p)
}

func (svc *ProductAttributePriceService) GetAll() (*types.ProductAttributePrices, error) {
	p := &types.ProductAttributePrices{}
	return p, svc.client.getAll(types.ProductAttributePriceModel, p)
}

func (svc *ProductAttributePriceService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProductAttributePriceModel, fields, relations)
}

func (svc *ProductAttributePriceService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductAttributePriceModel, ids, fields, relations)
}

func (svc *ProductAttributePriceService) Delete(ids []int64) error {
	return svc.client.delete(types.ProductAttributePriceModel, ids)
}
