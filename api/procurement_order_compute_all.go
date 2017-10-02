package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProcurementOrderComputeAllService struct {
	client *Client
}

func NewProcurementOrderComputeAllService(c *Client) *ProcurementOrderComputeAllService {
	return &ProcurementOrderComputeAllService{client: c}
}

func (svc *ProcurementOrderComputeAllService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProcurementOrderComputeAllModel, name)
}

func (svc *ProcurementOrderComputeAllService) GetByIds(ids []int64) (*types.ProcurementOrderComputeAlls, error) {
	p := &types.ProcurementOrderComputeAlls{}
	return p, svc.client.getByIds(types.ProcurementOrderComputeAllModel, ids, p)
}

func (svc *ProcurementOrderComputeAllService) GetByName(name string) (*types.ProcurementOrderComputeAlls, error) {
	p := &types.ProcurementOrderComputeAlls{}
	return p, svc.client.getByName(types.ProcurementOrderComputeAllModel, name, p)
}

func (svc *ProcurementOrderComputeAllService) GetByField(field string, value string) (*types.ProcurementOrderComputeAlls, error) {
	p := &types.ProcurementOrderComputeAlls{}
	return p, svc.client.getByField(types.ProcurementOrderComputeAllModel, field, value, p)
}

func (svc *ProcurementOrderComputeAllService) GetAll() (*types.ProcurementOrderComputeAlls, error) {
	p := &types.ProcurementOrderComputeAlls{}
	return p, svc.client.getAll(types.ProcurementOrderComputeAllModel, p)
}

func (svc *ProcurementOrderComputeAllService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProcurementOrderComputeAllModel, fields, relations)
}

func (svc *ProcurementOrderComputeAllService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProcurementOrderComputeAllModel, ids, fields, relations)
}

func (svc *ProcurementOrderComputeAllService) Delete(ids []int64) error {
	return svc.client.delete(types.ProcurementOrderComputeAllModel, ids)
}
