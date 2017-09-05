package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldMonetaryService struct {
	client *Client
}

func NewIrQwebFieldMonetaryService(c *Client) *IrQwebFieldMonetaryService {
	return &IrQwebFieldMonetaryService{client: c}
}

func (svc *IrQwebFieldMonetaryService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldMonetaryModel, name)
}

func (svc *IrQwebFieldMonetaryService) GetByIds(ids []int) (*types.IrQwebFieldMonetarys, error) {
	i := &types.IrQwebFieldMonetarys{}
	return i, svc.client.getByIds(types.IrQwebFieldMonetaryModel, ids, i)
}

func (svc *IrQwebFieldMonetaryService) GetByName(name string) (*types.IrQwebFieldMonetarys, error) {
	i := &types.IrQwebFieldMonetarys{}
	return i, svc.client.getByName(types.IrQwebFieldMonetaryModel, name, i)
}

func (svc *IrQwebFieldMonetaryService) GetByField(field string, value string) (*types.IrQwebFieldMonetarys, error) {
	i := &types.IrQwebFieldMonetarys{}
	return i, svc.client.getByName(types.IrQwebFieldMonetaryModel, field, value, i)
}

func (svc *IrQwebFieldMonetaryService) GetAll() (*types.IrQwebFieldMonetarys, error) {
	i := &types.IrQwebFieldMonetarys{}
	return i, svc.client.getAll(types.IrQwebFieldMonetaryModel, i)
}

func (svc *IrQwebFieldMonetaryService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldMonetaryModel, fields, relations)
}

func (svc *IrQwebFieldMonetaryService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldMonetaryModel, ids, fields, relations)
}

func (svc *IrQwebFieldMonetaryService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldMonetaryModel, ids)
}
