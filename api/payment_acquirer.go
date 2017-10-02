package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PaymentAcquirerService struct {
	client *Client
}

func NewPaymentAcquirerService(c *Client) *PaymentAcquirerService {
	return &PaymentAcquirerService{client: c}
}

func (svc *PaymentAcquirerService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PaymentAcquirerModel, name)
}

func (svc *PaymentAcquirerService) GetByIds(ids []int64) (*types.PaymentAcquirers, error) {
	p := &types.PaymentAcquirers{}
	return p, svc.client.getByIds(types.PaymentAcquirerModel, ids, p)
}

func (svc *PaymentAcquirerService) GetByName(name string) (*types.PaymentAcquirers, error) {
	p := &types.PaymentAcquirers{}
	return p, svc.client.getByName(types.PaymentAcquirerModel, name, p)
}

func (svc *PaymentAcquirerService) GetByField(field string, value string) (*types.PaymentAcquirers, error) {
	p := &types.PaymentAcquirers{}
	return p, svc.client.getByField(types.PaymentAcquirerModel, field, value, p)
}

func (svc *PaymentAcquirerService) GetAll() (*types.PaymentAcquirers, error) {
	p := &types.PaymentAcquirers{}
	return p, svc.client.getAll(types.PaymentAcquirerModel, p)
}

func (svc *PaymentAcquirerService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PaymentAcquirerModel, fields, relations)
}

func (svc *PaymentAcquirerService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PaymentAcquirerModel, ids, fields, relations)
}

func (svc *PaymentAcquirerService) Delete(ids []int64) error {
	return svc.client.delete(types.PaymentAcquirerModel, ids)
}
