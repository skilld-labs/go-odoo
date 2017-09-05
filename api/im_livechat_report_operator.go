package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ImLivechatReportOperatorService struct {
	client *Client
}

func NewImLivechatReportOperatorService(c *Client) *ImLivechatReportOperatorService {
	return &ImLivechatReportOperatorService{client: c}
}

func (svc *ImLivechatReportOperatorService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ImLivechatReportOperatorModel, name)
}

func (svc *ImLivechatReportOperatorService) GetByIds(ids []int) (*types.ImLivechatReportOperators, error) {
	i := &types.ImLivechatReportOperators{}
	return i, svc.client.getByIds(types.ImLivechatReportOperatorModel, ids, i)
}

func (svc *ImLivechatReportOperatorService) GetByName(name string) (*types.ImLivechatReportOperators, error) {
	i := &types.ImLivechatReportOperators{}
	return i, svc.client.getByName(types.ImLivechatReportOperatorModel, name, i)
}

func (svc *ImLivechatReportOperatorService) GetByField(field string, value string) (*types.ImLivechatReportOperators, error) {
	i := &types.ImLivechatReportOperators{}
	return i, svc.client.getByName(types.ImLivechatReportOperatorModel, field, value, i)
}

func (svc *ImLivechatReportOperatorService) GetAll() (*types.ImLivechatReportOperators, error) {
	i := &types.ImLivechatReportOperators{}
	return i, svc.client.getAll(types.ImLivechatReportOperatorModel, i)
}

func (svc *ImLivechatReportOperatorService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ImLivechatReportOperatorModel, fields, relations)
}

func (svc *ImLivechatReportOperatorService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ImLivechatReportOperatorModel, ids, fields, relations)
}

func (svc *ImLivechatReportOperatorService) Delete(ids []int) error {
	return svc.client.delete(types.ImLivechatReportOperatorModel, ids)
}
