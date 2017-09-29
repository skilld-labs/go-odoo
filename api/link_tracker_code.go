package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type LinkTrackerCodeService struct {
	client *Client
}

func NewLinkTrackerCodeService(c *Client) *LinkTrackerCodeService {
	return &LinkTrackerCodeService{client: c}
}

func (svc *LinkTrackerCodeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.LinkTrackerCodeModel, name)
}

func (svc *LinkTrackerCodeService) GetByIds(ids []int64) (*types.LinkTrackerCodes, error) {
	l := &types.LinkTrackerCodes{}
	return l, svc.client.getByIds(types.LinkTrackerCodeModel, ids, l)
}

func (svc *LinkTrackerCodeService) GetByName(name string) (*types.LinkTrackerCodes, error) {
	l := &types.LinkTrackerCodes{}
	return l, svc.client.getByName(types.LinkTrackerCodeModel, name, l)
}

func (svc *LinkTrackerCodeService) GetByField(field string, value string) (*types.LinkTrackerCodes, error) {
	l := &types.LinkTrackerCodes{}
	return l, svc.client.getByField(types.LinkTrackerCodeModel, field, value, l)
}

func (svc *LinkTrackerCodeService) GetAll() (*types.LinkTrackerCodes, error) {
	l := &types.LinkTrackerCodes{}
	return l, svc.client.getAll(types.LinkTrackerCodeModel, l)
}

func (svc *LinkTrackerCodeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.LinkTrackerCodeModel, fields, relations)
}

func (svc *LinkTrackerCodeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.LinkTrackerCodeModel, ids, fields, relations)
}

func (svc *LinkTrackerCodeService) Delete(ids []int64) error {
	return svc.client.delete(types.LinkTrackerCodeModel, ids)
}
