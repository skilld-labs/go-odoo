package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrNeedactionMixinService struct {
	client *Client
}

func NewIrNeedactionMixinService(c *Client) *IrNeedactionMixinService {
	return &IrNeedactionMixinService{client: c}
}

func (svc *IrNeedactionMixinService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrNeedactionMixinModel, name)
}

func (svc *IrNeedactionMixinService) GetByIds(ids []int64) (*types.IrNeedactionMixins, error) {
	i := &types.IrNeedactionMixins{}
	return i, svc.client.getByIds(types.IrNeedactionMixinModel, ids, i)
}

func (svc *IrNeedactionMixinService) GetByName(name string) (*types.IrNeedactionMixins, error) {
	i := &types.IrNeedactionMixins{}
	return i, svc.client.getByName(types.IrNeedactionMixinModel, name, i)
}

func (svc *IrNeedactionMixinService) GetByField(field string, value string) (*types.IrNeedactionMixins, error) {
	i := &types.IrNeedactionMixins{}
	return i, svc.client.getByField(types.IrNeedactionMixinModel, field, value, i)
}

func (svc *IrNeedactionMixinService) GetAll() (*types.IrNeedactionMixins, error) {
	i := &types.IrNeedactionMixins{}
	return i, svc.client.getAll(types.IrNeedactionMixinModel, i)
}

func (svc *IrNeedactionMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrNeedactionMixinModel, fields, relations)
}

func (svc *IrNeedactionMixinService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrNeedactionMixinModel, ids, fields, relations)
}

func (svc *IrNeedactionMixinService) Delete(ids []int64) error {
	return svc.client.delete(types.IrNeedactionMixinModel, ids)
}
