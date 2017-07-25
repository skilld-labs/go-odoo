package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldMany2oneService struct {
	client *Client
}

func NewIrQwebFieldMany2oneService(c *Client) *IrQwebFieldMany2oneService {
	return &IrQwebFieldMany2oneService{client: c}
}

func (svc *IrQwebFieldMany2oneService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldMany2oneModel, name)
}

func (svc *IrQwebFieldMany2oneService) GetByIds(ids []int) (*types.IrQwebFieldMany2ones, error) {
	i := &types.IrQwebFieldMany2ones{}
	return i, svc.client.getByIds(types.IrQwebFieldMany2oneModel, ids, i)
}

func (svc *IrQwebFieldMany2oneService) GetByName(name string) (*types.IrQwebFieldMany2ones, error) {
	i := &types.IrQwebFieldMany2ones{}
	return i, svc.client.getByName(types.IrQwebFieldMany2oneModel, name, i)
}

func (svc *IrQwebFieldMany2oneService) GetAll() (*types.IrQwebFieldMany2ones, error) {
	i := &types.IrQwebFieldMany2ones{}
	return i, svc.client.getAll(types.IrQwebFieldMany2oneModel, i)
}

func (svc *IrQwebFieldMany2oneService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldMany2oneModel, fields, relations)
}

func (svc *IrQwebFieldMany2oneService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldMany2oneModel, ids, fields, relations)
}

func (svc *IrQwebFieldMany2oneService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldMany2oneModel, ids)
}
