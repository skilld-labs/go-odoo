package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BusPresenceService struct {
	client *Client
}

func NewBusPresenceService(c *Client) *BusPresenceService {
	return &BusPresenceService{client: c}
}

func (svc *BusPresenceService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BusPresenceModel, name)
}

func (svc *BusPresenceService) GetByIds(ids []int) (*types.BusPresences, error) {
	b := &types.BusPresences{}
	return b, svc.client.getByIds(types.BusPresenceModel, ids, b)
}

func (svc *BusPresenceService) GetByName(name string) (*types.BusPresences, error) {
	b := &types.BusPresences{}
	return b, svc.client.getByName(types.BusPresenceModel, name, b)
}

func (svc *BusPresenceService) GetAll() (*types.BusPresences, error) {
	b := &types.BusPresences{}
	return b, svc.client.getAll(types.BusPresenceModel, b)
}

func (svc *BusPresenceService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BusPresenceModel, fields, relations)
}

func (svc *BusPresenceService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BusPresenceModel, ids, fields, relations)
}

func (svc *BusPresenceService) Delete(ids []int) error {
	return svc.client.delete(types.BusPresenceModel, ids)
}
