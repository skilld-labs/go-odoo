package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type RatingRatingService struct {
	client *Client
}

func NewRatingRatingService(c *Client) *RatingRatingService {
	return &RatingRatingService{client: c}
}

func (svc *RatingRatingService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.RatingRatingModel, name)
}

func (svc *RatingRatingService) GetByIds(ids []int) (*types.RatingRatings, error) {
	r := &types.RatingRatings{}
	return r, svc.client.getByIds(types.RatingRatingModel, ids, r)
}

func (svc *RatingRatingService) GetByName(name string) (*types.RatingRatings, error) {
	r := &types.RatingRatings{}
	return r, svc.client.getByName(types.RatingRatingModel, name, r)
}

func (svc *RatingRatingService) GetByField(field string, value string) (*types.RatingRatings, error) {
	r := &types.RatingRatings{}
	return r, svc.client.getByName(types.RatingRatingModel, field, value, r)
}

func (svc *RatingRatingService) GetAll() (*types.RatingRatings, error) {
	r := &types.RatingRatings{}
	return r, svc.client.getAll(types.RatingRatingModel, r)
}

func (svc *RatingRatingService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.RatingRatingModel, fields, relations)
}

func (svc *RatingRatingService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.RatingRatingModel, ids, fields, relations)
}

func (svc *RatingRatingService) Delete(ids []int) error {
	return svc.client.delete(types.RatingRatingModel, ids)
}
