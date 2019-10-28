package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProductRemovalService struct {
	client *Client
}

func NewProductRemovalService(c *Client) *ProductRemovalService {
	return &ProductRemovalService{client: c}
}

func (svc *ProductRemovalService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProductRemovalModel, name)
}

func (svc *ProductRemovalService) GetByIds(ids []int64) (*types.ProductRemovals, error) {
	p := &types.ProductRemovals{}
	return p, svc.client.getByIds(types.ProductRemovalModel, ids, p)
}

func (svc *ProductRemovalService) GetByName(name string) (*types.ProductRemovals, error) {
	p := &types.ProductRemovals{}
	return p, svc.client.getByName(types.ProductRemovalModel, name, p)
}

func (svc *ProductRemovalService) GetByField(field string, value string) (*types.ProductRemovals, error) {
	p := &types.ProductRemovals{}
	return p, svc.client.getByField(types.ProductRemovalModel, field, value, p)
}

func (svc *ProductRemovalService) GetAll() (*types.ProductRemovals, error) {
	p := &types.ProductRemovals{}
	return p, svc.client.getAll(types.ProductRemovalModel, p)
}

func (svc *ProductRemovalService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProductRemovalModel, fields, relations)
}

func (svc *ProductRemovalService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductRemovalModel, ids, fields, relations)
}

func (svc *ProductRemovalService) Delete(ids []int64) error {
	return svc.client.delete(types.ProductRemovalModel, ids)
}
