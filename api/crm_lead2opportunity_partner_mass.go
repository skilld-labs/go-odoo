package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmLead2opportunityPartnerMassService struct {
	client *Client
}

func NewCrmLead2opportunityPartnerMassService(c *Client) *CrmLead2opportunityPartnerMassService {
	return &CrmLead2opportunityPartnerMassService{client: c}
}

func (svc *CrmLead2opportunityPartnerMassService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CrmLead2opportunityPartnerMassModel, name)
}

func (svc *CrmLead2opportunityPartnerMassService) GetByIds(ids []int64) (*types.CrmLead2opportunityPartnerMasss, error) {
	c := &types.CrmLead2opportunityPartnerMasss{}
	return c, svc.client.getByIds(types.CrmLead2opportunityPartnerMassModel, ids, c)
}

func (svc *CrmLead2opportunityPartnerMassService) GetByName(name string) (*types.CrmLead2opportunityPartnerMasss, error) {
	c := &types.CrmLead2opportunityPartnerMasss{}
	return c, svc.client.getByName(types.CrmLead2opportunityPartnerMassModel, name, c)
}

func (svc *CrmLead2opportunityPartnerMassService) GetByField(field string, value string) (*types.CrmLead2opportunityPartnerMasss, error) {
	c := &types.CrmLead2opportunityPartnerMasss{}
	return c, svc.client.getByField(types.CrmLead2opportunityPartnerMassModel, field, value, c)
}

func (svc *CrmLead2opportunityPartnerMassService) GetAll() (*types.CrmLead2opportunityPartnerMasss, error) {
	c := &types.CrmLead2opportunityPartnerMasss{}
	return c, svc.client.getAll(types.CrmLead2opportunityPartnerMassModel, c)
}

func (svc *CrmLead2opportunityPartnerMassService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CrmLead2opportunityPartnerMassModel, fields, relations)
}

func (svc *CrmLead2opportunityPartnerMassService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmLead2opportunityPartnerMassModel, ids, fields, relations)
}

func (svc *CrmLead2opportunityPartnerMassService) Delete(ids []int64) error {
	return svc.client.delete(types.CrmLead2opportunityPartnerMassModel, ids)
}
