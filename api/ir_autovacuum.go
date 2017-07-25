package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrAutovacuumService struct {
	client *Client
}

func NewIrAutovacuumService(c *Client) *IrAutovacuumService {
	return &IrAutovacuumService{client: c}
}

func (svc *IrAutovacuumService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrAutovacuumModel, name)
}

func (svc *IrAutovacuumService) GetByIds(ids []int) (*types.IrAutovacuums, error) {
	i := &types.IrAutovacuums{}
	return i, svc.client.getByIds(types.IrAutovacuumModel, ids, i)
}

func (svc *IrAutovacuumService) GetByName(name string) (*types.IrAutovacuums, error) {
	i := &types.IrAutovacuums{}
	return i, svc.client.getByName(types.IrAutovacuumModel, name, i)
}

func (svc *IrAutovacuumService) GetAll() (*types.IrAutovacuums, error) {
	i := &types.IrAutovacuums{}
	return i, svc.client.getAll(types.IrAutovacuumModel, i)
}

func (svc *IrAutovacuumService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrAutovacuumModel, fields, relations)
}

func (svc *IrAutovacuumService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrAutovacuumModel, ids, fields, relations)
}

func (svc *IrAutovacuumService) Delete(ids []int) error {
	return svc.client.delete(types.IrAutovacuumModel, ids)
}
