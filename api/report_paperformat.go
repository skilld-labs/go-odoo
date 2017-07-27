package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportPaperformatService struct {
	client *Client
}

func NewReportPaperformatService(c *Client) *ReportPaperformatService {
	return &ReportPaperformatService{client: c}
}

func (svc *ReportPaperformatService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportPaperformatModel, name)
}

func (svc *ReportPaperformatService) GetByIds(ids []int) (*types.ReportPaperformats, error) {
	r := &types.ReportPaperformats{}
	return r, svc.client.getByIds(types.ReportPaperformatModel, ids, r)
}

func (svc *ReportPaperformatService) GetByName(name string) (*types.ReportPaperformats, error) {
	r := &types.ReportPaperformats{}
	return r, svc.client.getByName(types.ReportPaperformatModel, name, r)
}

func (svc *ReportPaperformatService) GetAll() (*types.ReportPaperformats, error) {
	r := &types.ReportPaperformats{}
	return r, svc.client.getAll(types.ReportPaperformatModel, r)
}

func (svc *ReportPaperformatService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportPaperformatModel, fields, relations)
}

func (svc *ReportPaperformatService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportPaperformatModel, ids, fields, relations)
}

func (svc *ReportPaperformatService) Delete(ids []int) error {
	return svc.client.delete(types.ReportPaperformatModel, ids)
}
