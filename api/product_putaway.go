package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductPutawayService struct {
	client *Client
}

func NewProductPutawayService(c *Client) *ProductPutawayService {
	return &ProductPutawayService{client: c}
}

func (svc *ProductPutawayService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProductPutawayModel, name)
}

func (svc *ProductPutawayService) GetByIds(ids []int64) (*types.ProductPutaways, error) {
	p := &types.ProductPutaways{}
	return p, svc.client.getByIds(types.ProductPutawayModel, ids, p)
}

func (svc *ProductPutawayService) GetByName(name string) (*types.ProductPutaways, error) {
	p := &types.ProductPutaways{}
	return p, svc.client.getByName(types.ProductPutawayModel, name, p)
}

func (svc *ProductPutawayService) GetByField(field string, value string) (*types.ProductPutaways, error) {
	p := &types.ProductPutaways{}
	return p, svc.client.getByField(types.ProductPutawayModel, field, value, p)
}

func (svc *ProductPutawayService) GetAll() (*types.ProductPutaways, error) {
	p := &types.ProductPutaways{}
	return p, svc.client.getAll(types.ProductPutawayModel, p)
}

func (svc *ProductPutawayService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProductPutawayModel, fields, relations)
}

func (svc *ProductPutawayService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductPutawayModel, ids, fields, relations)
}

func (svc *ProductPutawayService) Delete(ids []int64) error {
	return svc.client.delete(types.ProductPutawayModel, ids)
}
