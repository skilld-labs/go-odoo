package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResUsersLogService struct {
	client *Client
}

func NewResUsersLogService(c *Client) *ResUsersLogService {
	return &ResUsersLogService{client: c}
}

func (svc *ResUsersLogService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResUsersLogModel, name)
}

func (svc *ResUsersLogService) GetByIds(ids []int) (*types.ResUsersLogs, error) {
	r := &types.ResUsersLogs{}
	return r, svc.client.getByIds(types.ResUsersLogModel, ids, r)
}

func (svc *ResUsersLogService) GetByName(name string) (*types.ResUsersLogs, error) {
	r := &types.ResUsersLogs{}
	return r, svc.client.getByName(types.ResUsersLogModel, name, r)
}

func (svc *ResUsersLogService) GetByField(field string, value string) (*types.ResUsersLogs, error) {
	r := &types.ResUsersLogs{}
	return r, svc.client.getByField(types.ResUsersLogModel, field, value, r)
}

func (svc *ResUsersLogService) GetAll() (*types.ResUsersLogs, error) {
	r := &types.ResUsersLogs{}
	return r, svc.client.getAll(types.ResUsersLogModel, r)
}

func (svc *ResUsersLogService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResUsersLogModel, fields, relations)
}

func (svc *ResUsersLogService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResUsersLogModel, ids, fields, relations)
}

func (svc *ResUsersLogService) Delete(ids []int) error {
	return svc.client.delete(types.ResUsersLogModel, ids)
}
