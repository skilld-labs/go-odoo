package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailTestSimpleService struct {
	client *Client
}

func NewMailTestSimpleService(c *Client) *MailTestSimpleService {
	return &MailTestSimpleService{client: c}
}

func (svc *MailTestSimpleService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailTestSimpleModel, name)
}

func (svc *MailTestSimpleService) GetByIds(ids []int) (*types.MailTestSimples, error) {
	m := &types.MailTestSimples{}
	return m, svc.client.getByIds(types.MailTestSimpleModel, ids, m)
}

func (svc *MailTestSimpleService) GetByName(name string) (*types.MailTestSimples, error) {
	m := &types.MailTestSimples{}
	return m, svc.client.getByName(types.MailTestSimpleModel, name, m)
}

func (svc *MailTestSimpleService) GetByField(field string, value string) (*types.MailTestSimples, error) {
	m := &types.MailTestSimples{}
	return m, svc.client.getByField(types.MailTestSimpleModel, field, value, m)
}

func (svc *MailTestSimpleService) GetAll() (*types.MailTestSimples, error) {
	m := &types.MailTestSimples{}
	return m, svc.client.getAll(types.MailTestSimpleModel, m)
}

func (svc *MailTestSimpleService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailTestSimpleModel, fields, relations)
}

func (svc *MailTestSimpleService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailTestSimpleModel, ids, fields, relations)
}

func (svc *MailTestSimpleService) Delete(ids []int) error {
	return svc.client.delete(types.MailTestSimpleModel, ids)
}
