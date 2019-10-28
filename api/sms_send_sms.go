package api

import (
	"github.com/morezig/go-odoo/types"
)

type SmsSendSmsService struct {
	client *Client
}

func NewSmsSendSmsService(c *Client) *SmsSendSmsService {
	return &SmsSendSmsService{client: c}
}

func (svc *SmsSendSmsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.SmsSendSmsModel, name)
}

func (svc *SmsSendSmsService) GetByIds(ids []int64) (*types.SmsSendSmss, error) {
	s := &types.SmsSendSmss{}
	return s, svc.client.getByIds(types.SmsSendSmsModel, ids, s)
}

func (svc *SmsSendSmsService) GetByName(name string) (*types.SmsSendSmss, error) {
	s := &types.SmsSendSmss{}
	return s, svc.client.getByName(types.SmsSendSmsModel, name, s)
}

func (svc *SmsSendSmsService) GetByField(field string, value string) (*types.SmsSendSmss, error) {
	s := &types.SmsSendSmss{}
	return s, svc.client.getByField(types.SmsSendSmsModel, field, value, s)
}

func (svc *SmsSendSmsService) GetAll() (*types.SmsSendSmss, error) {
	s := &types.SmsSendSmss{}
	return s, svc.client.getAll(types.SmsSendSmsModel, s)
}

func (svc *SmsSendSmsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.SmsSendSmsModel, fields, relations)
}

func (svc *SmsSendSmsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SmsSendSmsModel, ids, fields, relations)
}

func (svc *SmsSendSmsService) Delete(ids []int64) error {
	return svc.client.delete(types.SmsSendSmsModel, ids)
}
