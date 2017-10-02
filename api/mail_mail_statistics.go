package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailMailStatisticsService struct {
	client *Client
}

func NewMailMailStatisticsService(c *Client) *MailMailStatisticsService {
	return &MailMailStatisticsService{client: c}
}

func (svc *MailMailStatisticsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailMailStatisticsModel, name)
}

func (svc *MailMailStatisticsService) GetByIds(ids []int64) (*types.MailMailStatisticss, error) {
	m := &types.MailMailStatisticss{}
	return m, svc.client.getByIds(types.MailMailStatisticsModel, ids, m)
}

func (svc *MailMailStatisticsService) GetByName(name string) (*types.MailMailStatisticss, error) {
	m := &types.MailMailStatisticss{}
	return m, svc.client.getByName(types.MailMailStatisticsModel, name, m)
}

func (svc *MailMailStatisticsService) GetByField(field string, value string) (*types.MailMailStatisticss, error) {
	m := &types.MailMailStatisticss{}
	return m, svc.client.getByField(types.MailMailStatisticsModel, field, value, m)
}

func (svc *MailMailStatisticsService) GetAll() (*types.MailMailStatisticss, error) {
	m := &types.MailMailStatisticss{}
	return m, svc.client.getAll(types.MailMailStatisticsModel, m)
}

func (svc *MailMailStatisticsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailMailStatisticsModel, fields, relations)
}

func (svc *MailMailStatisticsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMailStatisticsModel, ids, fields, relations)
}

func (svc *MailMailStatisticsService) Delete(ids []int64) error {
	return svc.client.delete(types.MailMailStatisticsModel, ids)
}
