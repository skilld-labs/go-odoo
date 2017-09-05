package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrTranslationService struct {
	client *Client
}

func NewIrTranslationService(c *Client) *IrTranslationService {
	return &IrTranslationService{client: c}
}

func (svc *IrTranslationService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrTranslationModel, name)
}

func (svc *IrTranslationService) GetByIds(ids []int) (*types.IrTranslations, error) {
	i := &types.IrTranslations{}
	return i, svc.client.getByIds(types.IrTranslationModel, ids, i)
}

func (svc *IrTranslationService) GetByName(name string) (*types.IrTranslations, error) {
	i := &types.IrTranslations{}
	return i, svc.client.getByName(types.IrTranslationModel, name, i)
}

func (svc *IrTranslationService) GetByField(field string, value string) (*types.IrTranslations, error) {
	i := &types.IrTranslations{}
	return i, svc.client.getByName(types.IrTranslationModel, field, value, i)
}

func (svc *IrTranslationService) GetAll() (*types.IrTranslations, error) {
	i := &types.IrTranslations{}
	return i, svc.client.getAll(types.IrTranslationModel, i)
}

func (svc *IrTranslationService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrTranslationModel, fields, relations)
}

func (svc *IrTranslationService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrTranslationModel, ids, fields, relations)
}

func (svc *IrTranslationService) Delete(ids []int) error {
	return svc.client.delete(types.IrTranslationModel, ids)
}
