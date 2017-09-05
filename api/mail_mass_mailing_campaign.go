package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailMassMailingCampaignService struct {
	client *Client
}

func NewMailMassMailingCampaignService(c *Client) *MailMassMailingCampaignService {
	return &MailMassMailingCampaignService{client: c}
}

func (svc *MailMassMailingCampaignService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailMassMailingCampaignModel, name)
}

func (svc *MailMassMailingCampaignService) GetByIds(ids []int) (*types.MailMassMailingCampaigns, error) {
	m := &types.MailMassMailingCampaigns{}
	return m, svc.client.getByIds(types.MailMassMailingCampaignModel, ids, m)
}

func (svc *MailMassMailingCampaignService) GetByName(name string) (*types.MailMassMailingCampaigns, error) {
	m := &types.MailMassMailingCampaigns{}
	return m, svc.client.getByName(types.MailMassMailingCampaignModel, name, m)
}

func (svc *MailMassMailingCampaignService) GetByField(field string, value string) (*types.MailMassMailingCampaigns, error) {
	m := &types.MailMassMailingCampaigns{}
	return m, svc.client.getByField(types.MailMassMailingCampaignModel, field, value, m)
}

func (svc *MailMassMailingCampaignService) GetAll() (*types.MailMassMailingCampaigns, error) {
	m := &types.MailMassMailingCampaigns{}
	return m, svc.client.getAll(types.MailMassMailingCampaignModel, m)
}

func (svc *MailMassMailingCampaignService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailMassMailingCampaignModel, fields, relations)
}

func (svc *MailMassMailingCampaignService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMassMailingCampaignModel, ids, fields, relations)
}

func (svc *MailMassMailingCampaignService) Delete(ids []int) error {
	return svc.client.delete(types.MailMassMailingCampaignModel, ids)
}
