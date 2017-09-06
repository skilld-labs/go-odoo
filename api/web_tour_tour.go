package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WebTourTourService struct {
	client *Client
}

func NewWebTourTourService(c *Client) *WebTourTourService {
	return &WebTourTourService{client: c}
}

func (svc *WebTourTourService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WebTourTourModel, name)
}

func (svc *WebTourTourService) GetByIds(ids []int) (*types.WebTourTours, error) {
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

func (svc *WebTourTourService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WebTourTourModel, fields, relations)
}

func (svc *WebTourTourService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WebTourTourModel, ids, fields, relations)
}

func (svc *WebTourTourService) Delete(ids []int) error {
	return svc.client.delete(types.WebTourTourModel, ids)
}
