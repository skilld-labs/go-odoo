package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrSequenceDateRangeService struct {
	client *Client
}

func NewIrSequenceDateRangeService(c *Client) *IrSequenceDateRangeService {
	return &IrSequenceDateRangeService{client: c}
}

func (svc *IrSequenceDateRangeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrSequenceDateRangeModel, name)
}

func (svc *IrSequenceDateRangeService) GetByIds(ids []int64) (*types.IrSequenceDateRanges, error) {
	i := &types.IrSequenceDateRanges{}
	return i, svc.client.getByIds(types.IrSequenceDateRangeModel, ids, i)
}

func (svc *IrSequenceDateRangeService) GetByName(name string) (*types.IrSequenceDateRanges, error) {
	i := &types.IrSequenceDateRanges{}
	return i, svc.client.getByName(types.IrSequenceDateRangeModel, name, i)
}

func (svc *IrSequenceDateRangeService) GetByField(field string, value string) (*types.IrSequenceDateRanges, error) {
	i := &types.IrSequenceDateRanges{}
	return i, svc.client.getByField(types.IrSequenceDateRangeModel, field, value, i)
}

func (svc *IrSequenceDateRangeService) GetAll() (*types.IrSequenceDateRanges, error) {
	i := &types.IrSequenceDateRanges{}
	return i, svc.client.getAll(types.IrSequenceDateRangeModel, i)
}

func (svc *IrSequenceDateRangeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrSequenceDateRangeModel, fields, relations)
}

func (svc *IrSequenceDateRangeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrSequenceDateRangeModel, ids, fields, relations)
}

func (svc *IrSequenceDateRangeService) Delete(ids []int64) error {
	return svc.client.delete(types.IrSequenceDateRangeModel, ids)
}
