package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductPriceListService struct {
	client *Client
}

func NewProductPriceListService(c *Client) *ProductPriceListService {
	return &ProductPriceListService{client: c}
}

func (svc *ProductPriceListService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductPriceListModel, name)
}

func (svc *ProductPriceListService) GetByIds(ids []int) (*types.ProductPriceLists, error) {
	p := &types.ProductPriceLists{}
	return p, svc.client.getByIds(types.ProductPriceListModel, ids, p)
}

func (svc *ProductPriceListService) GetByName(name string) (*types.ProductPriceLists, error) {
	p := &types.ProductPriceLists{}
	return p, svc.client.getByName(types.ProductPriceListModel, name, p)
}

func (svc *ProductPriceListService) GetAll() (*types.ProductPriceLists, error) {
	p := &types.ProductPriceLists{}
	return p, svc.client.getAll(types.ProductPriceListModel, p)
}

func (svc *ProductPriceListService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductPriceListModel, fields, relations)
}

func (svc *ProductPriceListService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductPriceListModel, ids, fields, relations)
}

func (svc *ProductPriceListService) Delete(ids []int) error {
	return svc.client.delete(types.ProductPriceListModel, ids)
}
