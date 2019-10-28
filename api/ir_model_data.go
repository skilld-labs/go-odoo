package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrModelDataService struct {
	client *Client
}

func NewIrModelDataService(c *Client) *IrModelDataService {
	return &IrModelDataService{client: c}
}

func (svc *IrModelDataService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrModelDataModel, name)
}

func (svc *IrModelDataService) GetByIds(ids []int64) (*types.IrModelDatas, error) {
	i := &types.IrModelDatas{}
	return i, svc.client.getByIds(types.IrModelDataModel, ids, i)
}

func (svc *IrModelDataService) GetByName(name string) (*types.IrModelDatas, error) {
	i := &types.IrModelDatas{}
	return i, svc.client.getByName(types.IrModelDataModel, name, i)
}

func (svc *IrModelDataService) GetByField(field string, value string) (*types.IrModelDatas, error) {
	i := &types.IrModelDatas{}
	return i, svc.client.getByField(types.IrModelDataModel, field, value, i)
}

func (svc *IrModelDataService) GetAll() (*types.IrModelDatas, error) {
	i := &types.IrModelDatas{}
	return i, svc.client.getAll(types.IrModelDataModel, i)
}

func (svc *IrModelDataService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrModelDataModel, fields, relations)
}

func (svc *IrModelDataService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModelDataModel, ids, fields, relations)
}

func (svc *IrModelDataService) Delete(ids []int64) error {
	return svc.client.delete(types.IrModelDataModel, ids)
}
