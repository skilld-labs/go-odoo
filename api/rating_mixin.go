package api

import (
	"github.com/morezig/go-odoo/types"
)

type RatingMixinService struct {
	client *Client
}

func NewRatingMixinService(c *Client) *RatingMixinService {
	return &RatingMixinService{client: c}
}

func (svc *RatingMixinService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.RatingMixinModel, name)
}

func (svc *RatingMixinService) GetByIds(ids []int64) (*types.RatingMixins, error) {
	r := &types.RatingMixins{}
	return r, svc.client.getByIds(types.RatingMixinModel, ids, r)
}

func (svc *RatingMixinService) GetByName(name string) (*types.RatingMixins, error) {
	r := &types.RatingMixins{}
	return r, svc.client.getByName(types.RatingMixinModel, name, r)
}

func (svc *RatingMixinService) GetByField(field string, value string) (*types.RatingMixins, error) {
	r := &types.RatingMixins{}
	return r, svc.client.getByField(types.RatingMixinModel, field, value, r)
}

func (svc *RatingMixinService) GetAll() (*types.RatingMixins, error) {
	r := &types.RatingMixins{}
	return r, svc.client.getAll(types.RatingMixinModel, r)
}

func (svc *RatingMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.RatingMixinModel, fields, relations)
}

func (svc *RatingMixinService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.RatingMixinModel, ids, fields, relations)
}

func (svc *RatingMixinService) Delete(ids []int64) error {
	return svc.client.delete(types.RatingMixinModel, ids)
}
