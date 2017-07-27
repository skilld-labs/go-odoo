package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PublisherWarrantyContractService struct {
	client *Client
}

func NewPublisherWarrantyContractService(c *Client) *PublisherWarrantyContractService {
	return &PublisherWarrantyContractService{client: c}
}

func (svc *PublisherWarrantyContractService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.PublisherWarrantyContractModel, name)
}

func (svc *PublisherWarrantyContractService) GetByIds(ids []int) (*types.PublisherWarrantyContracts, error) {
	p := &types.PublisherWarrantyContracts{}
	return p, svc.client.getByIds(types.PublisherWarrantyContractModel, ids, p)
}

func (svc *PublisherWarrantyContractService) GetByName(name string) (*types.PublisherWarrantyContracts, error) {
	p := &types.PublisherWarrantyContracts{}
	return p, svc.client.getByName(types.PublisherWarrantyContractModel, name, p)
}

func (svc *PublisherWarrantyContractService) GetAll() (*types.PublisherWarrantyContracts, error) {
	p := &types.PublisherWarrantyContracts{}
	return p, svc.client.getAll(types.PublisherWarrantyContractModel, p)
}

func (svc *PublisherWarrantyContractService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.PublisherWarrantyContractModel, fields, relations)
}

func (svc *PublisherWarrantyContractService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PublisherWarrantyContractModel, ids, fields, relations)
}

func (svc *PublisherWarrantyContractService) Delete(ids []int) error {
	return svc.client.delete(types.PublisherWarrantyContractModel, ids)
}
