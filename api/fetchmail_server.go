package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type FetchmailServerService struct {
	client *Client
}

func NewFetchmailServerService(c *Client) *FetchmailServerService {
	return &FetchmailServerService{client: c}
}

func (svc *FetchmailServerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.FetchmailServerModel, name)
}

func (svc *FetchmailServerService) GetByIds(ids []int) (*types.FetchmailServers, error) {
	f := &types.FetchmailServers{}
	return f, svc.client.getByIds(types.FetchmailServerModel, ids, f)
}

func (svc *FetchmailServerService) GetByName(name string) (*types.FetchmailServers, error) {
	f := &types.FetchmailServers{}
	return f, svc.client.getByName(types.FetchmailServerModel, name, f)
}

func (svc *FetchmailServerService) GetByField(field string, value string) (*types.FetchmailServers, error) {
	f := &types.FetchmailServers{}
	return f, svc.client.getByName(types.FetchmailServerModel, field, value, f)
}

func (svc *FetchmailServerService) GetAll() (*types.FetchmailServers, error) {
	f := &types.FetchmailServers{}
	return f, svc.client.getAll(types.FetchmailServerModel, f)
}

func (svc *FetchmailServerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.FetchmailServerModel, fields, relations)
}

func (svc *FetchmailServerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.FetchmailServerModel, ids, fields, relations)
}

func (svc *FetchmailServerService) Delete(ids []int) error {
	return svc.client.delete(types.FetchmailServerModel, ids)
}
