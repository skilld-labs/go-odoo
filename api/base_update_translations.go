package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseUpdateTranslationsService struct {
	client *Client
}

func NewBaseUpdateTranslationsService(c *Client) *BaseUpdateTranslationsService {
	return &BaseUpdateTranslationsService{client: c}
}

func (svc *BaseUpdateTranslationsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseUpdateTranslationsModel, name)
}

func (svc *BaseUpdateTranslationsService) GetByIds(ids []int) (*types.BaseUpdateTranslationss, error) {
	b := &types.BaseUpdateTranslationss{}
	return b, svc.client.getByIds(types.BaseUpdateTranslationsModel, ids, b)
}

func (svc *BaseUpdateTranslationsService) GetByName(name string) (*types.BaseUpdateTranslationss, error) {
	b := &types.BaseUpdateTranslationss{}
	return b, svc.client.getByName(types.BaseUpdateTranslationsModel, name, b)
}

func (svc *BaseUpdateTranslationsService) GetAll() (*types.BaseUpdateTranslationss, error) {
	b := &types.BaseUpdateTranslationss{}
	return b, svc.client.getAll(types.BaseUpdateTranslationsModel, b)
}

func (svc *BaseUpdateTranslationsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseUpdateTranslationsModel, fields, relations)
}

func (svc *BaseUpdateTranslationsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseUpdateTranslationsModel, ids, fields, relations)
}

func (svc *BaseUpdateTranslationsService) Delete(ids []int) error {
	return svc.client.delete(types.BaseUpdateTranslationsModel, ids)
}
