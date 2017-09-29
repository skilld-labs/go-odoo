package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BasePartnerMergeLineService struct {
	client *Client
}

func NewBasePartnerMergeLineService(c *Client) *BasePartnerMergeLineService {
	return &BasePartnerMergeLineService{client: c}
}

func (svc *BasePartnerMergeLineService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BasePartnerMergeLineModel, name)
}

func (svc *BasePartnerMergeLineService) GetByIds(ids []int64) (*types.BasePartnerMergeLines, error) {
	b := &types.BasePartnerMergeLines{}
	return b, svc.client.getByIds(types.BasePartnerMergeLineModel, ids, b)
}

func (svc *BasePartnerMergeLineService) GetByName(name string) (*types.BasePartnerMergeLines, error) {
	b := &types.BasePartnerMergeLines{}
	return b, svc.client.getByName(types.BasePartnerMergeLineModel, name, b)
}

func (svc *BasePartnerMergeLineService) GetByField(field string, value string) (*types.BasePartnerMergeLines, error) {
	b := &types.BasePartnerMergeLines{}
	return b, svc.client.getByField(types.BasePartnerMergeLineModel, field, value, b)
}

func (svc *BasePartnerMergeLineService) GetAll() (*types.BasePartnerMergeLines, error) {
	b := &types.BasePartnerMergeLines{}
	return b, svc.client.getAll(types.BasePartnerMergeLineModel, b)
}

func (svc *BasePartnerMergeLineService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BasePartnerMergeLineModel, fields, relations)
}

func (svc *BasePartnerMergeLineService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BasePartnerMergeLineModel, ids, fields, relations)
}

func (svc *BasePartnerMergeLineService) Delete(ids []int64) error {
	return svc.client.delete(types.BasePartnerMergeLineModel, ids)
}
