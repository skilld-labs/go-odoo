package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldContactService struct {
	client *Client
}

func NewIrQwebFieldContactService(c *Client) *IrQwebFieldContactService {
	return &IrQwebFieldContactService{client: c}
}

func (svc *IrQwebFieldContactService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrQwebFieldContactModel, name)
}

func (svc *IrQwebFieldContactService) GetByIds(ids []int64) (*types.IrQwebFieldContacts, error) {
	i := &types.IrQwebFieldContacts{}
	return i, svc.client.getByIds(types.IrQwebFieldContactModel, ids, i)
}

func (svc *IrQwebFieldContactService) GetByName(name string) (*types.IrQwebFieldContacts, error) {
	i := &types.IrQwebFieldContacts{}
	return i, svc.client.getByName(types.IrQwebFieldContactModel, name, i)
}

func (svc *IrQwebFieldContactService) GetByField(field string, value string) (*types.IrQwebFieldContacts, error) {
	i := &types.IrQwebFieldContacts{}
	return i, svc.client.getByField(types.IrQwebFieldContactModel, field, value, i)
}

func (svc *IrQwebFieldContactService) GetAll() (*types.IrQwebFieldContacts, error) {
	i := &types.IrQwebFieldContacts{}
	return i, svc.client.getAll(types.IrQwebFieldContactModel, i)
}

func (svc *IrQwebFieldContactService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrQwebFieldContactModel, fields, relations)
}

func (svc *IrQwebFieldContactService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldContactModel, ids, fields, relations)
}

func (svc *IrQwebFieldContactService) Delete(ids []int64) error {
	return svc.client.delete(types.IrQwebFieldContactModel, ids)
}
