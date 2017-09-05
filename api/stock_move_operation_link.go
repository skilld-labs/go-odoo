package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockMoveOperationLinkService struct {
	client *Client
}

func NewStockMoveOperationLinkService(c *Client) *StockMoveOperationLinkService {
	return &StockMoveOperationLinkService{client: c}
}

func (svc *StockMoveOperationLinkService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockMoveOperationLinkModel, name)
}

func (svc *StockMoveOperationLinkService) GetByIds(ids []int) (*types.StockMoveOperationLinks, error) {
	s := &types.StockMoveOperationLinks{}
	return s, svc.client.getByIds(types.StockMoveOperationLinkModel, ids, s)
}

func (svc *StockMoveOperationLinkService) GetByName(name string) (*types.StockMoveOperationLinks, error) {
	s := &types.StockMoveOperationLinks{}
	return s, svc.client.getByName(types.StockMoveOperationLinkModel, name, s)
}

func (svc *StockMoveOperationLinkService) GetByField(field string, value string) (*types.StockMoveOperationLinks, error) {
	s := &types.StockMoveOperationLinks{}
	return s, svc.client.getByField(types.StockMoveOperationLinkModel, field, value, s)
}

func (svc *StockMoveOperationLinkService) GetAll() (*types.StockMoveOperationLinks, error) {
	s := &types.StockMoveOperationLinks{}
	return s, svc.client.getAll(types.StockMoveOperationLinkModel, s)
}

func (svc *StockMoveOperationLinkService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockMoveOperationLinkModel, fields, relations)
}

func (svc *StockMoveOperationLinkService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockMoveOperationLinkModel, ids, fields, relations)
}

func (svc *StockMoveOperationLinkService) Delete(ids []int) error {
	return svc.client.delete(types.StockMoveOperationLinkModel, ids)
}
