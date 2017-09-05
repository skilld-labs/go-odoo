package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductPackagingService struct {
	client *Client
}

func NewProductPackagingService(c *Client) *ProductPackagingService {
	return &ProductPackagingService{client: c}
}

func (svc *ProductPackagingService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProductPackagingModel, name)
}

func (svc *ProductPackagingService) GetByIds(ids []int) (*types.ProductPackagings, error) {
	p := &types.ProductPackagings{}
	return p, svc.client.getByIds(types.ProductPackagingModel, ids, p)
}

func (svc *ProductPackagingService) GetByName(name string) (*types.ProductPackagings, error) {
	p := &types.ProductPackagings{}
	return p, svc.client.getByName(types.ProductPackagingModel, name, p)
}

func (svc *ProductPackagingService) GetByField(field string, value string) (*types.ProductPackagings, error) {
	p := &types.ProductPackagings{}
	return p, svc.client.getByField(types.ProductPackagingModel, field, value, p)
}

func (svc *ProductPackagingService) GetAll() (*types.ProductPackagings, error) {
	p := &types.ProductPackagings{}
	return p, svc.client.getAll(types.ProductPackagingModel, p)
}

func (svc *ProductPackagingService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProductPackagingModel, fields, relations)
}

func (svc *ProductPackagingService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductPackagingModel, ids, fields, relations)
}

func (svc *ProductPackagingService) Delete(ids []int) error {
	return svc.client.delete(types.ProductPackagingModel, ids)
}
