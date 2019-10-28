package api

import (
	"github.com/morezig/go-odoo/types"
)

type StockChangeStandardPriceService struct {
	client *Client
}

func NewStockChangeStandardPriceService(c *Client) *StockChangeStandardPriceService {
	return &StockChangeStandardPriceService{client: c}
}

func (svc *StockChangeStandardPriceService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockChangeStandardPriceModel, name)
}

func (svc *StockChangeStandardPriceService) GetByIds(ids []int64) (*types.StockChangeStandardPrices, error) {
	s := &types.StockChangeStandardPrices{}
	return s, svc.client.getByIds(types.StockChangeStandardPriceModel, ids, s)
}

func (svc *StockChangeStandardPriceService) GetByName(name string) (*types.StockChangeStandardPrices, error) {
	s := &types.StockChangeStandardPrices{}
	return s, svc.client.getByName(types.StockChangeStandardPriceModel, name, s)
}

func (svc *StockChangeStandardPriceService) GetByField(field string, value string) (*types.StockChangeStandardPrices, error) {
	s := &types.StockChangeStandardPrices{}
	return s, svc.client.getByField(types.StockChangeStandardPriceModel, field, value, s)
}

func (svc *StockChangeStandardPriceService) GetAll() (*types.StockChangeStandardPrices, error) {
	s := &types.StockChangeStandardPrices{}
	return s, svc.client.getAll(types.StockChangeStandardPriceModel, s)
}

func (svc *StockChangeStandardPriceService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockChangeStandardPriceModel, fields, relations)
}

func (svc *StockChangeStandardPriceService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockChangeStandardPriceModel, ids, fields, relations)
}

func (svc *StockChangeStandardPriceService) Delete(ids []int64) error {
	return svc.client.delete(types.StockChangeStandardPriceModel, ids)
}
