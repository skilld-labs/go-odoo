package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailMassMailingStageService struct {
	client *Client
}

func NewMailMassMailingStageService(c *Client) *MailMassMailingStageService {
	return &MailMassMailingStageService{client: c}
}

func (svc *MailMassMailingStageService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailMassMailingStageModel, name)
}

func (svc *MailMassMailingStageService) GetByIds(ids []int) (*types.MailMassMailingStages, error) {
	m := &types.MailMassMailingStages{}
	return m, svc.client.getByIds(types.MailMassMailingStageModel, ids, m)
}

func (svc *MailMassMailingStageService) GetByName(name string) (*types.MailMassMailingStages, error) {
	m := &types.MailMassMailingStages{}
	return m, svc.client.getByName(types.MailMassMailingStageModel, name, m)
}

func (svc *MailMassMailingStageService) GetByField(field string, value string) (*types.MailMassMailingStages, error) {
	m := &types.MailMassMailingStages{}
	return m, svc.client.getByField(types.MailMassMailingStageModel, field, value, m)
}

func (svc *MailMassMailingStageService) GetAll() (*types.MailMassMailingStages, error) {
	m := &types.MailMassMailingStages{}
	return m, svc.client.getAll(types.MailMassMailingStageModel, m)
}

func (svc *MailMassMailingStageService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailMassMailingStageModel, fields, relations)
}

func (svc *MailMassMailingStageService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMassMailingStageModel, ids, fields, relations)
}

func (svc *MailMassMailingStageService) Delete(ids []int) error {
	return svc.client.delete(types.MailMassMailingStageModel, ids)
}
