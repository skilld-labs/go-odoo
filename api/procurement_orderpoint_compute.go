package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProcurementOrderpointComputeService struct {
	client *Client
}

func NewProcurementOrderpointComputeService(c *Client) *ProcurementOrderpointComputeService {
	return &ProcurementOrderpointComputeService{client: c}
}

func (svc *ProcurementOrderpointComputeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProcurementOrderpointComputeModel, name)
}

func (svc *ProcurementOrderpointComputeService) GetByIds(ids []int) (*types.ProcurementOrderpointComputes, error) {
	p := &types.ProcurementOrderpointComputes{}
	return p, svc.client.getByIds(types.ProcurementOrderpointComputeModel, ids, p)
}

func (svc *ProcurementOrderpointComputeService) GetByName(name string) (*types.ProcurementOrderpointComputes, error) {
	p := &types.ProcurementOrderpointComputes{}
	return p, svc.client.getByName(types.ProcurementOrderpointComputeModel, name, p)
}

func (svc *ProcurementOrderpointComputeService) GetAll() (*types.ProcurementOrderpointComputes, error) {
	p := &types.ProcurementOrderpointComputes{}
	return p, svc.client.getAll(types.ProcurementOrderpointComputeModel, p)
}

func (svc *ProcurementOrderpointComputeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProcurementOrderpointComputeModel, fields, relations)
}

func (svc *ProcurementOrderpointComputeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProcurementOrderpointComputeModel, ids, fields, relations)
}

func (svc *ProcurementOrderpointComputeService) Delete(ids []int) error {
	return svc.client.delete(types.ProcurementOrderpointComputeModel, ids)
}
