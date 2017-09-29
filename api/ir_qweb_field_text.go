package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldTextService struct {
	client *Client
}

func NewIrQwebFieldTextService(c *Client) *IrQwebFieldTextService {
	return &IrQwebFieldTextService{client: c}
}

func (svc *IrQwebFieldTextService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrQwebFieldTextModel, name)
}

func (svc *IrQwebFieldTextService) GetByIds(ids []int64) (*types.IrQwebFieldTexts, error) {
	i := &types.IrQwebFieldTexts{}
	return i, svc.client.getByIds(types.IrQwebFieldTextModel, ids, i)
}

func (svc *IrQwebFieldTextService) GetByName(name string) (*types.IrQwebFieldTexts, error) {
	i := &types.IrQwebFieldTexts{}
	return i, svc.client.getByName(types.IrQwebFieldTextModel, name, i)
}

func (svc *IrQwebFieldTextService) GetByField(field string, value string) (*types.IrQwebFieldTexts, error) {
	i := &types.IrQwebFieldTexts{}
	return i, svc.client.getByField(types.IrQwebFieldTextModel, field, value, i)
}

func (svc *IrQwebFieldTextService) GetAll() (*types.IrQwebFieldTexts, error) {
	i := &types.IrQwebFieldTexts{}
	return i, svc.client.getAll(types.IrQwebFieldTextModel, i)
}

func (svc *IrQwebFieldTextService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrQwebFieldTextModel, fields, relations)
}

func (svc *IrQwebFieldTextService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldTextModel, ids, fields, relations)
}

func (svc *IrQwebFieldTextService) Delete(ids []int64) error {
	return svc.client.delete(types.IrQwebFieldTextModel, ids)
}
