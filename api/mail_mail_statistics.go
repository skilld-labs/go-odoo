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

func (svc *MailMailStatisticsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailMailStatisticsModel, name)
}

func (svc *MailMailStatisticsService) GetByIds(ids []int) (*types.MailMailStatisticss, error) {
	m := &types.MailMailStatisticss{}
	return m, svc.client.getByIds(types.MailMailStatisticsModel, ids, m)
}

func (svc *MailMailStatisticsService) GetByName(name string) (*types.MailMailStatisticss, error) {
	m := &types.MailMailStatisticss{}
	return m, svc.client.getByName(types.MailMailStatisticsModel, name, m)
}

func (svc *MailMailStatisticsService) GetByField(field string, value string) (*types.MailMailStatisticss, error) {
	m := &types.MailMailStatisticss{}
	return m, svc.client.getByName(types.MailMailStatisticsModel, field, value, m)
}

func (svc *MailMailStatisticsService) GetAll() (*types.MailMailStatisticss, error) {
	m := &types.MailMailStatisticss{}
	return m, svc.client.getAll(types.MailMailStatisticsModel, m)
}

func (svc *MailMailStatisticsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailMailStatisticsModel, fields, relations)
}

func (svc *MailMailStatisticsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMailStatisticsModel, ids, fields, relations)
}

func (svc *MailMailStatisticsService) Delete(ids []int) error {
	return svc.client.delete(types.MailMailStatisticsModel, ids)
}
