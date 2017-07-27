package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProcurementGroupService struct {
	client *Client
}

func NewProcurementGroupService(c *Client) *ProcurementGroupService {
	return &ProcurementGroupService{client: c}
}

func (svc *ProcurementGroupService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProcurementGroupModel, name)
}

func (svc *ProcurementGroupService) GetByIds(ids []int) (*types.ProcurementGroups, error) {
	p := &types.ProcurementGroups{}
	return p, svc.client.getByIds(types.ProcurementGroupModel, ids, p)
}

func (svc *ProcurementGroupService) GetByName(name string) (*types.ProcurementGroups, error) {
	p := &types.ProcurementGroups{}
	return p, svc.client.getByName(types.ProcurementGroupModel, name, p)
}

func (svc *ProcurementGroupService) GetAll() (*types.ProcurementGroups, error) {
	p := &types.ProcurementGroups{}
	return p, svc.client.getAll(types.ProcurementGroupModel, p)
}

func (svc *ProcurementGroupService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProcurementGroupModel, fields, relations)
}

func (svc *ProcurementGroupService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProcurementGroupModel, ids, fields, relations)
}

func (svc *ProcurementGroupService) Delete(ids []int) error {
	return svc.client.delete(types.ProcurementGroupModel, ids)
}
