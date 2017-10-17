package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAllChannelsSalesService struct {
	client *Client
}

func NewReportAllChannelsSalesService(c *Client) *ReportAllChannelsSalesService {
	return &ReportAllChannelsSalesService{client: c}
}

func (svc *ReportAllChannelsSalesService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportAllChannelsSalesModel, name)
}

func (svc *ReportAllChannelsSalesService) GetByIds(ids []int) (*types.ReportAllChannelsSaless, error) {
	r := &types.ReportAllChannelsSaless{}
	return r, svc.client.getByIds(types.ReportAllChannelsSalesModel, ids, r)
}

func (svc *ReportAllChannelsSalesService) GetByName(name string) (*types.ReportAllChannelsSaless, error) {
	r := &types.ReportAllChannelsSaless{}
	return r, svc.client.getByName(types.ReportAllChannelsSalesModel, name, r)
}

func (svc *ReportAllChannelsSalesService) GetByField(field string, value string) (*types.ReportAllChannelsSaless, error) {
	r := &types.ReportAllChannelsSaless{}
	return r, svc.client.getByField(types.ReportAllChannelsSalesModel, field, value, r)
}

func (svc *ReportAllChannelsSalesService) GetAll() (*types.ReportAllChannelsSaless, error) {
	r := &types.ReportAllChannelsSaless{}
	return r, svc.client.getAll(types.ReportAllChannelsSalesModel, r)
}

func (svc *ReportAllChannelsSalesService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportAllChannelsSalesModel, fields, relations)
}

func (svc *ReportAllChannelsSalesService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAllChannelsSalesModel, ids, fields, relations)
}

func (svc *ReportAllChannelsSalesService) Delete(ids []int) error {
	return svc.client.delete(types.ReportAllChannelsSalesModel, ids)
}
