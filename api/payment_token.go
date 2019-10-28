package api

import (
	"github.com/morezig/go-odoo/types"
)

type PaymentTokenService struct {
	client *Client
}

func NewPaymentTokenService(c *Client) *PaymentTokenService {
	return &PaymentTokenService{client: c}
}

func (svc *PaymentTokenService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PaymentTokenModel, name)
}

func (svc *PaymentTokenService) GetByIds(ids []int64) (*types.PaymentTokens, error) {
	p := &types.PaymentTokens{}
	return p, svc.client.getByIds(types.PaymentTokenModel, ids, p)
}

func (svc *PaymentTokenService) GetByName(name string) (*types.PaymentTokens, error) {
	p := &types.PaymentTokens{}
	return p, svc.client.getByName(types.PaymentTokenModel, name, p)
}

func (svc *PaymentTokenService) GetByField(field string, value string) (*types.PaymentTokens, error) {
	p := &types.PaymentTokens{}
	return p, svc.client.getByField(types.PaymentTokenModel, field, value, p)
}

func (svc *PaymentTokenService) GetAll() (*types.PaymentTokens, error) {
	p := &types.PaymentTokens{}
	return p, svc.client.getAll(types.PaymentTokenModel, p)
}

func (svc *PaymentTokenService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PaymentTokenModel, fields, relations)
}

func (svc *PaymentTokenService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PaymentTokenModel, ids, fields, relations)
}

func (svc *PaymentTokenService) Delete(ids []int64) error {
	return svc.client.delete(types.PaymentTokenModel, ids)
}
