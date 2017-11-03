package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type SmsApiService struct {
	client *Client
}

func NewSmsApiService(c *Client) *SmsApiService {
	return &SmsApiService{client: c}
}

func (svc *SmsApiService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.SmsApiModel, name)
}

func (svc *SmsApiService) GetByIds(ids []int64) (*types.SmsApis, error) {
	s := &types.SmsApis{}
	return s, svc.client.getByIds(types.SmsApiModel, ids, s)
}

func (svc *SmsApiService) GetByName(name string) (*types.SmsApis, error) {
	s := &types.SmsApis{}
	return s, svc.client.getByName(types.SmsApiModel, name, s)
}

func (svc *SmsApiService) GetByField(field string, value string) (*types.SmsApis, error) {
	s := &types.SmsApis{}
	return s, svc.client.getByField(types.SmsApiModel, field, value, s)
}

func (svc *SmsApiService) GetAll() (*types.SmsApis, error) {
	s := &types.SmsApis{}
	return s, svc.client.getAll(types.SmsApiModel, s)
}

func (svc *SmsApiService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.SmsApiModel, fields, relations)
}

func (svc *SmsApiService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SmsApiModel, ids, fields, relations)
}

func (svc *SmsApiService) Delete(ids []int64) error {
	return svc.client.delete(types.SmsApiModel, ids)
}
