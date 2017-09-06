package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductPricelistItemService struct {
	client *Client
}

func NewProductPricelistItemService(c *Client) *ProductPricelistItemService {
	return &ProductPricelistItemService{client: c}
}

func (svc *ProductPricelistItemService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductPricelistItemModel, name)
}

func (svc *ProductPricelistItemService) GetByIds(ids []int) (*types.ProductPricelistItems, error) {
	p := &types.ProductPricelistItems{}
	return p, svc.client.getByIds(types.ProductPricelistItemModel, ids, p)
}

func (svc *ProductPricelistItemService) GetByName(name string) (*types.ProductPricelistItems, error) {
	p := &types.ProductPricelistItems{}
	return p, svc.client.getByName(types.ProductPricelistItemModel, name, p)
}

func (svc *ProductPricelistItemService) GetByField(field string, value string) (*types.ProductPricelistItems, error) {
	p := &types.ProductPricelistItems{}
	return p, svc.client.getByField(types.ProductPricelistItemModel, field, value, p)
}

func (svc *ProductPricelistItemService) GetAll() (*types.ProductPricelistItems, error) {
	p := &types.ProductPricelistItems{}
	return p, svc.client.getAll(types.ProductPricelistItemModel, p)
}

func (svc *ProductPricelistItemService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductPricelistItemModel, fields, relations)
}

func (svc *ProductPricelistItemService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductPricelistItemModel, ids, fields, relations)
}

func (svc *ProductPricelistItemService) Delete(ids []int) error {
	return svc.client.delete(types.ProductPricelistItemModel, ids)
}
