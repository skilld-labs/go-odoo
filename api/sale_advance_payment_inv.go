package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type SaleAdvancePaymentInvService struct {
	client *Client
}

func NewSaleAdvancePaymentInvService(c *Client) *SaleAdvancePaymentInvService {
	return &SaleAdvancePaymentInvService{client: c}
}

func (svc *SaleAdvancePaymentInvService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.SaleAdvancePaymentInvModel, name)
}

func (svc *SaleAdvancePaymentInvService) GetByIds(ids []int) (*types.SaleAdvancePaymentInvs, error) {
	s := &types.SaleAdvancePaymentInvs{}
	return s, svc.client.getByIds(types.SaleAdvancePaymentInvModel, ids, s)
}

func (svc *SaleAdvancePaymentInvService) GetByName(name string) (*types.SaleAdvancePaymentInvs, error) {
	s := &types.SaleAdvancePaymentInvs{}
	return s, svc.client.getByName(types.SaleAdvancePaymentInvModel, name, s)
}

func (svc *SaleAdvancePaymentInvService) GetByField(field string, value string) (*types.SaleAdvancePaymentInvs, error) {
	s := &types.SaleAdvancePaymentInvs{}
	return s, svc.client.getByName(types.SaleAdvancePaymentInvModel, field, value, s)
}

func (svc *SaleAdvancePaymentInvService) GetAll() (*types.SaleAdvancePaymentInvs, error) {
	s := &types.SaleAdvancePaymentInvs{}
	return s, svc.client.getAll(types.SaleAdvancePaymentInvModel, s)
}

func (svc *SaleAdvancePaymentInvService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.SaleAdvancePaymentInvModel, fields, relations)
}

func (svc *SaleAdvancePaymentInvService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SaleAdvancePaymentInvModel, ids, fields, relations)
}

func (svc *SaleAdvancePaymentInvService) Delete(ids []int) error {
	return svc.client.delete(types.SaleAdvancePaymentInvModel, ids)
}
