package api

import (
	"github.com/morezig/go-odoo/types"
)

type PublisherWarrantyContractService struct {
	client *Client
}

func NewPublisherWarrantyContractService(c *Client) *PublisherWarrantyContractService {
	return &PublisherWarrantyContractService{client: c}
}

func (svc *PublisherWarrantyContractService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PublisherWarrantyContractModel, name)
}

func (svc *PublisherWarrantyContractService) GetByIds(ids []int64) (*types.PublisherWarrantyContracts, error) {
	p := &types.PublisherWarrantyContracts{}
	return p, svc.client.getByIds(types.PublisherWarrantyContractModel, ids, p)
}

func (svc *PublisherWarrantyContractService) GetByName(name string) (*types.PublisherWarrantyContracts, error) {
	p := &types.PublisherWarrantyContracts{}
	return p, svc.client.getByName(types.PublisherWarrantyContractModel, name, p)
}

func (svc *PublisherWarrantyContractService) GetByField(field string, value string) (*types.PublisherWarrantyContracts, error) {
	p := &types.PublisherWarrantyContracts{}
	return p, svc.client.getByField(types.PublisherWarrantyContractModel, field, value, p)
}

func (svc *PublisherWarrantyContractService) GetAll() (*types.PublisherWarrantyContracts, error) {
	p := &types.PublisherWarrantyContracts{}
	return p, svc.client.getAll(types.PublisherWarrantyContractModel, p)
}

func (svc *PublisherWarrantyContractService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PublisherWarrantyContractModel, fields, relations)
}

func (svc *PublisherWarrantyContractService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PublisherWarrantyContractModel, ids, fields, relations)
}

func (svc *PublisherWarrantyContractService) Delete(ids []int64) error {
	return svc.client.delete(types.PublisherWarrantyContractModel, ids)
}
