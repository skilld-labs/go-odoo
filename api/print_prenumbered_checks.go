package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PrintPrenumberedChecksService struct {
	client *Client
}

func NewPrintPrenumberedChecksService(c *Client) *PrintPrenumberedChecksService {
	return &PrintPrenumberedChecksService{client: c}
}

func (svc *PrintPrenumberedChecksService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PrintPrenumberedChecksModel, name)
}

func (svc *PrintPrenumberedChecksService) GetByIds(ids []int64) (*types.PrintPrenumberedCheckss, error) {
	p := &types.PrintPrenumberedCheckss{}
	return p, svc.client.getByIds(types.PrintPrenumberedChecksModel, ids, p)
}

func (svc *PrintPrenumberedChecksService) GetByName(name string) (*types.PrintPrenumberedCheckss, error) {
	p := &types.PrintPrenumberedCheckss{}
	return p, svc.client.getByName(types.PrintPrenumberedChecksModel, name, p)
}

func (svc *PrintPrenumberedChecksService) GetByField(field string, value string) (*types.PrintPrenumberedCheckss, error) {
	p := &types.PrintPrenumberedCheckss{}
	return p, svc.client.getByField(types.PrintPrenumberedChecksModel, field, value, p)
}

func (svc *PrintPrenumberedChecksService) GetAll() (*types.PrintPrenumberedCheckss, error) {
	p := &types.PrintPrenumberedCheckss{}
	return p, svc.client.getAll(types.PrintPrenumberedChecksModel, p)
}

func (svc *PrintPrenumberedChecksService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PrintPrenumberedChecksModel, fields, relations)
}

func (svc *PrintPrenumberedChecksService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PrintPrenumberedChecksModel, ids, fields, relations)
}

func (svc *PrintPrenumberedChecksService) Delete(ids []int64) error {
	return svc.client.delete(types.PrintPrenumberedChecksModel, ids)
}
