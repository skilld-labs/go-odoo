package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailWizardInviteService struct {
	client *Client
}

func NewMailWizardInviteService(c *Client) *MailWizardInviteService {
	return &MailWizardInviteService{client: c}
}

func (svc *MailWizardInviteService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailWizardInviteModel, name)
}

func (svc *MailWizardInviteService) GetByIds(ids []int) (*types.MailWizardInvites, error) {
	m := &types.MailWizardInvites{}
	return m, svc.client.getByIds(types.MailWizardInviteModel, ids, m)
}

func (svc *MailWizardInviteService) GetByName(name string) (*types.MailWizardInvites, error) {
	m := &types.MailWizardInvites{}
	return m, svc.client.getByName(types.MailWizardInviteModel, name, m)
}

func (svc *MailWizardInviteService) GetAll() (*types.MailWizardInvites, error) {
	m := &types.MailWizardInvites{}
	return m, svc.client.getAll(types.MailWizardInviteModel, m)
}

func (svc *MailWizardInviteService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailWizardInviteModel, fields, relations)
}

func (svc *MailWizardInviteService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailWizardInviteModel, ids, fields, relations)
}

func (svc *MailWizardInviteService) Delete(ids []int) error {
	return svc.client.delete(types.MailWizardInviteModel, ids)
}
