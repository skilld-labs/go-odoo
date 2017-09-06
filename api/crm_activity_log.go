package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmActivityLogService struct {
	client *Client
}

func NewCrmActivityLogService(c *Client) *CrmActivityLogService {
	return &CrmActivityLogService{client: c}
}

func (svc *CrmActivityLogService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CrmActivityLogModel, name)
}

func (svc *CrmActivityLogService) GetByIds(ids []int) (*types.CrmActivityLogs, error) {
	c := &types.CrmActivityLogs{}
	return c, svc.client.getByIds(types.CrmActivityLogModel, ids, c)
}

func (svc *CrmActivityLogService) GetByName(name string) (*types.CrmActivityLogs, error) {
	c := &types.CrmActivityLogs{}
	return c, svc.client.getByName(types.CrmActivityLogModel, name, c)
}

func (svc *CrmActivityLogService) GetByField(field string, value string) (*types.CrmActivityLogs, error) {
	c := &types.CrmActivityLogs{}
	return c, svc.client.getByField(types.CrmActivityLogModel, field, value, c)
}

func (svc *CrmActivityLogService) GetAll() (*types.CrmActivityLogs, error) {
	c := &types.CrmActivityLogs{}
	return c, svc.client.getAll(types.CrmActivityLogModel, c)
}

func (svc *CrmActivityLogService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CrmActivityLogModel, fields, relations)
}

func (svc *CrmActivityLogService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmActivityLogModel, ids, fields, relations)
}

func (svc *CrmActivityLogService) Delete(ids []int) error {
	return svc.client.delete(types.CrmActivityLogModel, ids)
}
