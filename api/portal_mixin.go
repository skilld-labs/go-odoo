package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PortalMixinService struct {
	client *Client
}

func NewPortalMixinService(c *Client) *PortalMixinService {
	return &PortalMixinService{client: c}
}

func (svc *PortalMixinService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PortalMixinModel, name)
}

func (svc *PortalMixinService) GetByIds(ids []int64) (*types.PortalMixins, error) {
	p := &types.PortalMixins{}
	return p, svc.client.getByIds(types.PortalMixinModel, ids, p)
}

func (svc *PortalMixinService) GetByName(name string) (*types.PortalMixins, error) {
	p := &types.PortalMixins{}
	return p, svc.client.getByName(types.PortalMixinModel, name, p)
}

func (svc *PortalMixinService) GetByField(field string, value string) (*types.PortalMixins, error) {
	p := &types.PortalMixins{}
	return p, svc.client.getByField(types.PortalMixinModel, field, value, p)
}

func (svc *PortalMixinService) GetAll() (*types.PortalMixins, error) {
	p := &types.PortalMixins{}
	return p, svc.client.getAll(types.PortalMixinModel, p)
}

func (svc *PortalMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PortalMixinModel, fields, relations)
}

func (svc *PortalMixinService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PortalMixinModel, ids, fields, relations)
}

func (svc *PortalMixinService) Delete(ids []int64) error {
	return svc.client.delete(types.PortalMixinModel, ids)
}
