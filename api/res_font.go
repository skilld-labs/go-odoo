package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResFontService struct {
	client *Client
}

func NewResFontService(c *Client) *ResFontService {
	return &ResFontService{client: c}
}

func (svc *ResFontService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResFontModel, name)
}

func (svc *ResFontService) GetByIds(ids []int) (*types.ResFonts, error) {
	r := &types.ResFonts{}
	return r, svc.client.getByIds(types.ResFontModel, ids, r)
}

func (svc *ResFontService) GetByName(name string) (*types.ResFonts, error) {
	r := &types.ResFonts{}
	return r, svc.client.getByName(types.ResFontModel, name, r)
}

func (svc *ResFontService) GetAll() (*types.ResFonts, error) {
	r := &types.ResFonts{}
	return r, svc.client.getAll(types.ResFontModel, r)
}

func (svc *ResFontService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResFontModel, fields, relations)
}

func (svc *ResFontService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResFontModel, ids, fields, relations)
}

func (svc *ResFontService) Delete(ids []int) error {
	return svc.client.delete(types.ResFontModel, ids)
}
