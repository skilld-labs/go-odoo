package api

import (
	"github.com/morezig/go-odoo/types"
)

type BusPresenceService struct {
	client *Client
}

func NewBusPresenceService(c *Client) *BusPresenceService {
	return &BusPresenceService{client: c}
}

func (svc *BusPresenceService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BusPresenceModel, name)
}

func (svc *BusPresenceService) GetByIds(ids []int64) (*types.BusPresences, error) {
	b := &types.BusPresences{}
	return b, svc.client.getByIds(types.BusPresenceModel, ids, b)
}

func (svc *BusPresenceService) GetByName(name string) (*types.BusPresences, error) {
	b := &types.BusPresences{}
	return b, svc.client.getByName(types.BusPresenceModel, name, b)
}

func (svc *BusPresenceService) GetByField(field string, value string) (*types.BusPresences, error) {
	b := &types.BusPresences{}
	return b, svc.client.getByField(types.BusPresenceModel, field, value, b)
}

func (svc *BusPresenceService) GetAll() (*types.BusPresences, error) {
	b := &types.BusPresences{}
	return b, svc.client.getAll(types.BusPresenceModel, b)
}

func (svc *BusPresenceService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BusPresenceModel, fields, relations)
}

func (svc *BusPresenceService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BusPresenceModel, ids, fields, relations)
}

func (svc *BusPresenceService) Delete(ids []int64) error {
	return svc.client.delete(types.BusPresenceModel, ids)
}
