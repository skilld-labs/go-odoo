package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResConfigInstallerService struct {
	client *Client
}

func NewResConfigInstallerService(c *Client) *ResConfigInstallerService {
	return &ResConfigInstallerService{client: c}
}

func (svc *ResConfigInstallerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResConfigInstallerModel, name)
}

func (svc *ResConfigInstallerService) GetByIds(ids []int) (*types.ResConfigInstallers, error) {
	r := &types.ResConfigInstallers{}
	return r, svc.client.getByIds(types.ResConfigInstallerModel, ids, r)
}

func (svc *ResConfigInstallerService) GetByName(name string) (*types.ResConfigInstallers, error) {
	r := &types.ResConfigInstallers{}
	return r, svc.client.getByName(types.ResConfigInstallerModel, name, r)
}

func (svc *ResConfigInstallerService) GetByField(field string, value string) (*types.ResConfigInstallers, error) {
	r := &types.ResConfigInstallers{}
	return r, svc.client.getByField(types.ResConfigInstallerModel, field, value, r)
}

func (svc *ResConfigInstallerService) GetAll() (*types.ResConfigInstallers, error) {
	r := &types.ResConfigInstallers{}
	return r, svc.client.getAll(types.ResConfigInstallerModel, r)
}

func (svc *ResConfigInstallerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResConfigInstallerModel, fields, relations)
}

func (svc *ResConfigInstallerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResConfigInstallerModel, ids, fields, relations)
}

func (svc *ResConfigInstallerService) Delete(ids []int) error {
	return svc.client.delete(types.ResConfigInstallerModel, ids)
}
