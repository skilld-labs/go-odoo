package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResCompanyLdapService struct {
	client *Client
}

func NewResCompanyLdapService(c *Client) *ResCompanyLdapService {
	return &ResCompanyLdapService{client: c}
}

func (svc *ResCompanyLdapService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResCompanyLdapModel, name)
}

func (svc *ResCompanyLdapService) GetByIds(ids []int) (*types.ResCompanyLdaps, error) {
	r := &types.ResCompanyLdaps{}
	return r, svc.client.getByIds(types.ResCompanyLdapModel, ids, r)
}

func (svc *ResCompanyLdapService) GetByName(name string) (*types.ResCompanyLdaps, error) {
	r := &types.ResCompanyLdaps{}
	return r, svc.client.getByName(types.ResCompanyLdapModel, name, r)
}

func (svc *ResCompanyLdapService) GetByField(field string, value string) (*types.ResCompanyLdaps, error) {
	r := &types.ResCompanyLdaps{}
	return r, svc.client.getByField(types.ResCompanyLdapModel, field, value, r)
}

func (svc *ResCompanyLdapService) GetAll() (*types.ResCompanyLdaps, error) {
	r := &types.ResCompanyLdaps{}
	return r, svc.client.getAll(types.ResCompanyLdapModel, r)
}

func (svc *ResCompanyLdapService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResCompanyLdapModel, fields, relations)
}

func (svc *ResCompanyLdapService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResCompanyLdapModel, ids, fields, relations)
}

func (svc *ResCompanyLdapService) Delete(ids []int) error {
	return svc.client.delete(types.ResCompanyLdapModel, ids)
}
