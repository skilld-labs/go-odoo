package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrCronService struct {
	client *Client
}

func NewIrCronService(c *Client) *IrCronService {
	return &IrCronService{client: c}
}

func (svc *IrCronService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrCronModel, name)
}

func (svc *IrCronService) GetByIds(ids []int) (*types.IrCrons, error) {
	i := &types.IrCrons{}
	return i, svc.client.getByIds(types.IrCronModel, ids, i)
}

func (svc *IrCronService) GetByName(name string) (*types.IrCrons, error) {
	i := &types.IrCrons{}
	return i, svc.client.getByName(types.IrCronModel, name, i)
}

func (svc *IrCronService) GetByField(field string, value string) (*types.IrCrons, error) {
	i := &types.IrCrons{}
	return i, svc.client.getByField(types.IrCronModel, field, value, i)
}

func (svc *IrCronService) GetAll() (*types.IrCrons, error) {
	i := &types.IrCrons{}
	return i, svc.client.getAll(types.IrCronModel, i)
}

func (svc *IrCronService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrCronModel, fields, relations)
}

func (svc *IrCronService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrCronModel, ids, fields, relations)
}

func (svc *IrCronService) Delete(ids []int) error {
	return svc.client.delete(types.IrCronModel, ids)
}
