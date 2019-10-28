package api

import (
	"github.com/morezig/go-odoo/types"
)

type WebTourTourService struct {
	client *Client
}

func NewWebTourTourService(c *Client) *WebTourTourService {
	return &WebTourTourService{client: c}
}

func (svc *WebTourTourService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.WebTourTourModel, name)
}

func (svc *WebTourTourService) GetByIds(ids []int64) (*types.WebTourTours, error) {
	w := &types.WebTourTours{}
	return w, svc.client.getByIds(types.WebTourTourModel, ids, w)
}

func (svc *WebTourTourService) GetByName(name string) (*types.WebTourTours, error) {
	w := &types.WebTourTours{}
	return w, svc.client.getByName(types.WebTourTourModel, name, w)
}

func (svc *WebTourTourService) GetByField(field string, value string) (*types.WebTourTours, error) {
	w := &types.WebTourTours{}
	return w, svc.client.getByField(types.WebTourTourModel, field, value, w)
}

func (svc *WebTourTourService) GetAll() (*types.WebTourTours, error) {
	w := &types.WebTourTours{}
	return w, svc.client.getAll(types.WebTourTourModel, w)
}

func (svc *WebTourTourService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.WebTourTourModel, fields, relations)
}

func (svc *WebTourTourService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WebTourTourModel, ids, fields, relations)
}

func (svc *WebTourTourService) Delete(ids []int64) error {
	return svc.client.delete(types.WebTourTourModel, ids)
}
