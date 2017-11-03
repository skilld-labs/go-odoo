package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PaymentIconService struct {
	client *Client
}

func NewPaymentIconService(c *Client) *PaymentIconService {
	return &PaymentIconService{client: c}
}

func (svc *PaymentIconService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PaymentIconModel, name)
}

func (svc *PaymentIconService) GetByIds(ids []int64) (*types.PaymentIcons, error) {
	p := &types.PaymentIcons{}
	return p, svc.client.getByIds(types.PaymentIconModel, ids, p)
}

func (svc *PaymentIconService) GetByName(name string) (*types.PaymentIcons, error) {
	p := &types.PaymentIcons{}
	return p, svc.client.getByName(types.PaymentIconModel, name, p)
}

func (svc *PaymentIconService) GetByField(field string, value string) (*types.PaymentIcons, error) {
	p := &types.PaymentIcons{}
	return p, svc.client.getByField(types.PaymentIconModel, field, value, p)
}

func (svc *PaymentIconService) GetAll() (*types.PaymentIcons, error) {
	p := &types.PaymentIcons{}
	return p, svc.client.getAll(types.PaymentIconModel, p)
}

func (svc *PaymentIconService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PaymentIconModel, fields, relations)
}

func (svc *PaymentIconService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PaymentIconModel, ids, fields, relations)
}

func (svc *PaymentIconService) Delete(ids []int64) error {
	return svc.client.delete(types.PaymentIconModel, ids)
}
