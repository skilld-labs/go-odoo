package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportProjectTaskUserService struct {
	client *Client
}

func NewReportProjectTaskUserService(c *Client) *ReportProjectTaskUserService {
	return &ReportProjectTaskUserService{client: c}
}

func (svc *ReportProjectTaskUserService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportProjectTaskUserModel, name)
}

func (svc *ReportProjectTaskUserService) GetByIds(ids []int) (*types.ReportProjectTaskUsers, error) {
	r := &types.ReportProjectTaskUsers{}
	return r, svc.client.getByIds(types.ReportProjectTaskUserModel, ids, r)
}

func (svc *ReportProjectTaskUserService) GetByName(name string) (*types.ReportProjectTaskUsers, error) {
	r := &types.ReportProjectTaskUsers{}
	return r, svc.client.getByName(types.ReportProjectTaskUserModel, name, r)
}

func (svc *ReportProjectTaskUserService) GetAll() (*types.ReportProjectTaskUsers, error) {
	r := &types.ReportProjectTaskUsers{}
	return r, svc.client.getAll(types.ReportProjectTaskUserModel, r)
}

func (svc *ReportProjectTaskUserService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportProjectTaskUserModel, fields, relations)
}

func (svc *ReportProjectTaskUserService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportProjectTaskUserModel, ids, fields, relations)
}

func (svc *ReportProjectTaskUserService) Delete(ids []int) error {
	return svc.client.delete(types.ReportProjectTaskUserModel, ids)
}
