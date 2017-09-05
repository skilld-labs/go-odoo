package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAccountReportJournalService struct {
	client *Client
}

func NewReportAccountReportJournalService(c *Client) *ReportAccountReportJournalService {
	return &ReportAccountReportJournalService{client: c}
}

func (svc *ReportAccountReportJournalService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportAccountReportJournalModel, name)
}

func (svc *ReportAccountReportJournalService) GetByIds(ids []int) (*types.ReportAccountReportJournals, error) {
	r := &types.ReportAccountReportJournals{}
	return r, svc.client.getByIds(types.ReportAccountReportJournalModel, ids, r)
}

func (svc *ReportAccountReportJournalService) GetByName(name string) (*types.ReportAccountReportJournals, error) {
	r := &types.ReportAccountReportJournals{}
	return r, svc.client.getByName(types.ReportAccountReportJournalModel, name, r)
}

func (svc *ReportAccountReportJournalService) GetByField(field string, value string) (*types.ReportAccountReportJournals, error) {
	r := &types.ReportAccountReportJournals{}
	return r, svc.client.getByField(types.ReportAccountReportJournalModel, field, value, r)
}

func (svc *ReportAccountReportJournalService) GetAll() (*types.ReportAccountReportJournals, error) {
	r := &types.ReportAccountReportJournals{}
	return r, svc.client.getAll(types.ReportAccountReportJournalModel, r)
}

func (svc *ReportAccountReportJournalService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportAccountReportJournalModel, fields, relations)
}

func (svc *ReportAccountReportJournalService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportJournalModel, ids, fields, relations)
}

func (svc *ReportAccountReportJournalService) Delete(ids []int) error {
	return svc.client.delete(types.ReportAccountReportJournalModel, ids)
}
