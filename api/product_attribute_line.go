package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProductAttributeLineService struct {
	client *Client
}

func NewProductAttributeLineService(c *Client) *ProductAttributeLineService {
	return &ProductAttributeLineService{client: c}
}

func (svc *ProductAttributeLineService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProductAttributeLineModel, name)
}

func (svc *ProductAttributeLineService) GetByIds(ids []int64) (*types.ProductAttributeLines, error) {
	p := &types.ProductAttributeLines{}
	return p, svc.client.getByIds(types.ProductAttributeLineModel, ids, p)
}

func (svc *ProductAttributeLineService) GetByName(name string) (*types.ProductAttributeLines, error) {
	p := &types.ProductAttributeLines{}
	return p, svc.client.getByName(types.ProductAttributeLineModel, name, p)
}

func (svc *ProductAttributeLineService) GetByField(field string, value string) (*types.ProductAttributeLines, error) {
	p := &types.ProductAttributeLines{}
	return p, svc.client.getByField(types.ProductAttributeLineModel, field, value, p)
}

func (svc *ProductAttributeLineService) GetAll() (*types.ProductAttributeLines, error) {
	p := &types.ProductAttributeLines{}
	return p, svc.client.getAll(types.ProductAttributeLineModel, p)
}

func (svc *ProductAttributeLineService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProductAttributeLineModel, fields, relations)
}

func (svc *ProductAttributeLineService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductAttributeLineModel, ids, fields, relations)
}

func (svc *ProductAttributeLineService) Delete(ids []int64) error {
	return svc.client.delete(types.ProductAttributeLineModel, ids)
}
