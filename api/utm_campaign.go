package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type UtmCampaignService struct {
	client *Client
}

func NewUtmCampaignService(c *Client) *UtmCampaignService {
	return &UtmCampaignService{client: c}
}

func (svc *UtmCampaignService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.UtmCampaignModel, name)
}

func (svc *UtmCampaignService) GetByIds(ids []int64) (*types.UtmCampaigns, error) {
	u := &types.UtmCampaigns{}
	return u, svc.client.getByIds(types.UtmCampaignModel, ids, u)
}

func (svc *UtmCampaignService) GetByName(name string) (*types.UtmCampaigns, error) {
	u := &types.UtmCampaigns{}
	return u, svc.client.getByName(types.UtmCampaignModel, name, u)
}

func (svc *UtmCampaignService) GetByField(field string, value string) (*types.UtmCampaigns, error) {
	u := &types.UtmCampaigns{}
	return u, svc.client.getByField(types.UtmCampaignModel, field, value, u)
}

func (svc *UtmCampaignService) GetAll() (*types.UtmCampaigns, error) {
	u := &types.UtmCampaigns{}
	return u, svc.client.getAll(types.UtmCampaignModel, u)
}

func (svc *UtmCampaignService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.UtmCampaignModel, fields, relations)
}

func (svc *UtmCampaignService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.UtmCampaignModel, ids, fields, relations)
}

func (svc *UtmCampaignService) Delete(ids []int64) error {
	return svc.client.delete(types.UtmCampaignModel, ids)
}
