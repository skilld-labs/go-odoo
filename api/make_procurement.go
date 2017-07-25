package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MakeProcurementService struct {
	client *Client
}

func NewMakeProcurementService(c *Client) *MakeProcurementService {
	return &MakeProcurementService{client: c}
}

func (svc *MakeProcurementService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MakeProcurementModel, name)
}

func (svc *MakeProcurementService) GetByIds(ids []int) (*types.MakeProcurements, error) {
	m := &types.MakeProcurements{}
	return m, svc.client.getByIds(types.MakeProcurementModel, ids, m)
}

func (svc *MakeProcurementService) GetByName(name string) (*types.MakeProcurements, error) {
	m := &types.MakeProcurements{}
	return m, svc.client.getByName(types.MakeProcurementModel, name, m)
}

func (svc *MakeProcurementService) GetAll() (*types.MakeProcurements, error) {
	m := &types.MakeProcurements{}
	return m, svc.client.getAll(types.MakeProcurementModel, m)
}

func (svc *MakeProcurementService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MakeProcurementModel, fields, relations)
}

func (svc *MakeProcurementService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MakeProcurementModel, ids, fields, relations)
}

func (svc *MakeProcurementService) Delete(ids []int) error {
	return svc.client.delete(types.MakeProcurementModel, ids)
}
