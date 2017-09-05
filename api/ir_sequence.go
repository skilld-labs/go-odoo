package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrSequenceService struct {
	client *Client
}

func NewIrSequenceService(c *Client) *IrSequenceService {
	return &IrSequenceService{client: c}
}

func (svc *IrSequenceService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrSequenceModel, name)
}

func (svc *IrSequenceService) GetByIds(ids []int) (*types.IrSequences, error) {
	i := &types.IrSequences{}
	return i, svc.client.getByIds(types.IrSequenceModel, ids, i)
}

func (svc *IrSequenceService) GetByName(name string) (*types.IrSequences, error) {
	i := &types.IrSequences{}
	return i, svc.client.getByName(types.IrSequenceModel, name, i)
}

func (svc *IrSequenceService) GetByField(field string, value string) (*types.IrSequences, error) {
	i := &types.IrSequences{}
	return i, svc.client.getByName(types.IrSequenceModel, field, value, i)
}

func (svc *IrSequenceService) GetAll() (*types.IrSequences, error) {
	i := &types.IrSequences{}
	return i, svc.client.getAll(types.IrSequenceModel, i)
}

func (svc *IrSequenceService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrSequenceModel, fields, relations)
}

func (svc *IrSequenceService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrSequenceModel, ids, fields, relations)
}

func (svc *IrSequenceService) Delete(ids []int) error {
	return svc.client.delete(types.IrSequenceModel, ids)
}
