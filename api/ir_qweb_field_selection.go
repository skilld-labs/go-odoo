package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldSelectionService struct {
	client *Client
}

func NewIrQwebFieldSelectionService(c *Client) *IrQwebFieldSelectionService {
	return &IrQwebFieldSelectionService{client: c}
}

func (svc *IrQwebFieldSelectionService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldSelectionModel, name)
}

func (svc *IrQwebFieldSelectionService) GetByIds(ids []int) (*types.IrQwebFieldSelections, error) {
	i := &types.IrQwebFieldSelections{}
	return i, svc.client.getByIds(types.IrQwebFieldSelectionModel, ids, i)
}

func (svc *IrQwebFieldSelectionService) GetByName(name string) (*types.IrQwebFieldSelections, error) {
	i := &types.IrQwebFieldSelections{}
	return i, svc.client.getByName(types.IrQwebFieldSelectionModel, name, i)
}

func (svc *IrQwebFieldSelectionService) GetAll() (*types.IrQwebFieldSelections, error) {
	i := &types.IrQwebFieldSelections{}
	return i, svc.client.getAll(types.IrQwebFieldSelectionModel, i)
}

func (svc *IrQwebFieldSelectionService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldSelectionModel, fields, relations)
}

func (svc *IrQwebFieldSelectionService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldSelectionModel, ids, fields, relations)
}

func (svc *IrQwebFieldSelectionService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldSelectionModel, ids)
}
