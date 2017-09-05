package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmTeamService struct {
	client *Client
}

func NewCrmTeamService(c *Client) *CrmTeamService {
	return &CrmTeamService{client: c}
}

func (svc *CrmTeamService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CrmTeamModel, name)
}

func (svc *CrmTeamService) GetByIds(ids []int) (*types.CrmTeams, error) {
	c := &types.CrmTeams{}
	return c, svc.client.getByIds(types.CrmTeamModel, ids, c)
}

func (svc *CrmTeamService) GetByName(name string) (*types.CrmTeams, error) {
	c := &types.CrmTeams{}
	return c, svc.client.getByName(types.CrmTeamModel, name, c)
}

func (svc *CrmTeamService) GetByField(field string, value string) (*types.CrmTeams, error) {
	c := &types.CrmTeams{}
	return c, svc.client.getByField(types.CrmTeamModel, field, value, c)
}

func (svc *CrmTeamService) GetAll() (*types.CrmTeams, error) {
	c := &types.CrmTeams{}
	return c, svc.client.getAll(types.CrmTeamModel, c)
}

func (svc *CrmTeamService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CrmTeamModel, fields, relations)
}

func (svc *CrmTeamService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmTeamModel, ids, fields, relations)
}

func (svc *CrmTeamService) Delete(ids []int) error {
	return svc.client.delete(types.CrmTeamModel, ids)
}
