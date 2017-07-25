package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductPriceHistoryService struct {
	client *Client
}

func NewProductPriceHistoryService(c *Client) *ProductPriceHistoryService {
	return &ProductPriceHistoryService{client: c}
}

func (svc *ProductPriceHistoryService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductPriceHistoryModel, name)
}

func (svc *ProductPriceHistoryService) GetByIds(ids []int) (*types.ProductPriceHistorys, error) {
	p := &types.ProductPriceHistorys{}
	return p, svc.client.getByIds(types.ProductPriceHistoryModel, ids, p)
}

func (svc *ProductPriceHistoryService) GetByName(name string) (*types.ProductPriceHistorys, error) {
	p := &types.ProductPriceHistorys{}
	return p, svc.client.getByName(types.ProductPriceHistoryModel, name, p)
}

func (svc *ProductPriceHistoryService) GetAll() (*types.ProductPriceHistorys, error) {
	p := &types.ProductPriceHistorys{}
	return p, svc.client.getAll(types.ProductPriceHistoryModel, p)
}

func (svc *ProductPriceHistoryService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductPriceHistoryModel, fields, relations)
}

func (svc *ProductPriceHistoryService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductPriceHistoryModel, ids, fields, relations)
}

func (svc *ProductPriceHistoryService) Delete(ids []int) error {
	return svc.client.delete(types.ProductPriceHistoryModel, ids)
}
