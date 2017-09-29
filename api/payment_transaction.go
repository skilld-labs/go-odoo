package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PaymentTransactionService struct {
	client *Client
}

func NewPaymentTransactionService(c *Client) *PaymentTransactionService {
	return &PaymentTransactionService{client: c}
}

func (svc *PaymentTransactionService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PaymentTransactionModel, name)
}

func (svc *PaymentTransactionService) GetByIds(ids []int64) (*types.PaymentTransactions, error) {
	p := &types.PaymentTransactions{}
	return p, svc.client.getByIds(types.PaymentTransactionModel, ids, p)
}

func (svc *PaymentTransactionService) GetByName(name string) (*types.PaymentTransactions, error) {
	p := &types.PaymentTransactions{}
	return p, svc.client.getByName(types.PaymentTransactionModel, name, p)
}

func (svc *PaymentTransactionService) GetByField(field string, value string) (*types.PaymentTransactions, error) {
	p := &types.PaymentTransactions{}
	return p, svc.client.getByField(types.PaymentTransactionModel, field, value, p)
}

func (svc *PaymentTransactionService) GetAll() (*types.PaymentTransactions, error) {
	p := &types.PaymentTransactions{}
	return p, svc.client.getAll(types.PaymentTransactionModel, p)
}

func (svc *PaymentTransactionService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PaymentTransactionModel, fields, relations)
}

func (svc *PaymentTransactionService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PaymentTransactionModel, ids, fields, relations)
}

func (svc *PaymentTransactionService) Delete(ids []int64) error {
	return svc.client.delete(types.PaymentTransactionModel, ids)
}
