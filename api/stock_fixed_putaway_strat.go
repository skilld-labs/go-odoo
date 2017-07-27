package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockFixedPutawayStratService struct {
	client *Client
}

func NewStockFixedPutawayStratService(c *Client) *StockFixedPutawayStratService {
	return &StockFixedPutawayStratService{client: c}
}

func (svc *StockFixedPutawayStratService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockFixedPutawayStratModel, name)
}

func (svc *StockFixedPutawayStratService) GetByIds(ids []int) (*types.StockFixedPutawayStrats, error) {
	s := &types.StockFixedPutawayStrats{}
	return s, svc.client.getByIds(types.StockFixedPutawayStratModel, ids, s)
}

func (svc *StockFixedPutawayStratService) GetByName(name string) (*types.StockFixedPutawayStrats, error) {
	s := &types.StockFixedPutawayStrats{}
	return s, svc.client.getByName(types.StockFixedPutawayStratModel, name, s)
}

func (svc *StockFixedPutawayStratService) GetAll() (*types.StockFixedPutawayStrats, error) {
	s := &types.StockFixedPutawayStrats{}
	return s, svc.client.getAll(types.StockFixedPutawayStratModel, s)
}

func (svc *StockFixedPutawayStratService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockFixedPutawayStratModel, fields, relations)
}

func (svc *StockFixedPutawayStratService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockFixedPutawayStratModel, ids, fields, relations)
}

func (svc *StockFixedPutawayStratService) Delete(ids []int) error {
	return svc.client.delete(types.StockFixedPutawayStratModel, ids)
}
