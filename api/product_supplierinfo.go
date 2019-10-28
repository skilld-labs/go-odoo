package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProductSupplierinfoService struct {
	client *Client
}

func NewProductSupplierinfoService(c *Client) *ProductSupplierinfoService {
	return &ProductSupplierinfoService{client: c}
}

func (svc *ProductSupplierinfoService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProductSupplierinfoModel, name)
}

func (svc *ProductSupplierinfoService) GetByIds(ids []int64) (*types.ProductSupplierinfos, error) {
	p := &types.ProductSupplierinfos{}
	return p, svc.client.getByIds(types.ProductSupplierinfoModel, ids, p)
}

func (svc *ProductSupplierinfoService) GetByName(name string) (*types.ProductSupplierinfos, error) {
	p := &types.ProductSupplierinfos{}
	return p, svc.client.getByName(types.ProductSupplierinfoModel, name, p)
}

func (svc *ProductSupplierinfoService) GetByField(field string, value string) (*types.ProductSupplierinfos, error) {
	p := &types.ProductSupplierinfos{}
	return p, svc.client.getByField(types.ProductSupplierinfoModel, field, value, p)
}

func (svc *ProductSupplierinfoService) GetAll() (*types.ProductSupplierinfos, error) {
	p := &types.ProductSupplierinfos{}
	return p, svc.client.getAll(types.ProductSupplierinfoModel, p)
}

func (svc *ProductSupplierinfoService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProductSupplierinfoModel, fields, relations)
}

func (svc *ProductSupplierinfoService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductSupplierinfoModel, ids, fields, relations)
}

func (svc *ProductSupplierinfoService) Delete(ids []int64) error {
	return svc.client.delete(types.ProductSupplierinfoModel, ids)
}
