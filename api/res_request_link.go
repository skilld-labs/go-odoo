package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResRequestLinkService struct {
	client *Client
}

func NewResRequestLinkService(c *Client) *ResRequestLinkService {
	return &ResRequestLinkService{client: c}
}

func (svc *ResRequestLinkService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResRequestLinkModel, name)
}

func (svc *ResRequestLinkService) GetByIds(ids []int) (*types.ResRequestLinks, error) {
	r := &types.ResRequestLinks{}
	return r, svc.client.getByIds(types.ResRequestLinkModel, ids, r)
}

func (svc *ResRequestLinkService) GetByName(name string) (*types.ResRequestLinks, error) {
	r := &types.ResRequestLinks{}
	return r, svc.client.getByName(types.ResRequestLinkModel, name, r)
}

func (svc *ResRequestLinkService) GetByField(field string, value string) (*types.ResRequestLinks, error) {
	r := &types.ResRequestLinks{}
	return r, svc.client.getByField(types.ResRequestLinkModel, field, value, r)
}

func (svc *ResRequestLinkService) GetAll() (*types.ResRequestLinks, error) {
	r := &types.ResRequestLinks{}
	return r, svc.client.getAll(types.ResRequestLinkModel, r)
}

func (svc *ResRequestLinkService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResRequestLinkModel, fields, relations)
}

func (svc *ResRequestLinkService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResRequestLinkModel, ids, fields, relations)
}

func (svc *ResRequestLinkService) Delete(ids []int) error {
	return svc.client.delete(types.ResRequestLinkModel, ids)
}
