package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IapAccountService struct {
	client *Client
}

func NewIapAccountService(c *Client) *IapAccountService {
	return &IapAccountService{client: c}
}

func (svc *IapAccountService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IapAccountModel, name)
}

func (svc *IapAccountService) GetByIds(ids []int) (*types.IapAccounts, error) {
	i := &types.IapAccounts{}
	return i, svc.client.getByIds(types.IapAccountModel, ids, i)
}

func (svc *IapAccountService) GetByName(name string) (*types.IapAccounts, error) {
	i := &types.IapAccounts{}
	return i, svc.client.getByName(types.IapAccountModel, name, i)
}

func (svc *IapAccountService) GetByField(field string, value string) (*types.IapAccounts, error) {
	i := &types.IapAccounts{}
	return i, svc.client.getByField(types.IapAccountModel, field, value, i)
}

func (svc *IapAccountService) GetAll() (*types.IapAccounts, error) {
	i := &types.IapAccounts{}
	return i, svc.client.getAll(types.IapAccountModel, i)
}

func (svc *IapAccountService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IapAccountModel, fields, relations)
}

func (svc *IapAccountService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IapAccountModel, ids, fields, relations)
}

func (svc *IapAccountService) Delete(ids []int) error {
	return svc.client.delete(types.IapAccountModel, ids)
}
