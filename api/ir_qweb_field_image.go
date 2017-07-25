package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldImageService struct {
	client *Client
}

func NewIrQwebFieldImageService(c *Client) *IrQwebFieldImageService {
	return &IrQwebFieldImageService{client: c}
}

func (svc *IrQwebFieldImageService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldImageModel, name)
}

func (svc *IrQwebFieldImageService) GetByIds(ids []int) (*types.IrQwebFieldImages, error) {
	i := &types.IrQwebFieldImages{}
	return i, svc.client.getByIds(types.IrQwebFieldImageModel, ids, i)
}

func (svc *IrQwebFieldImageService) GetByName(name string) (*types.IrQwebFieldImages, error) {
	i := &types.IrQwebFieldImages{}
	return i, svc.client.getByName(types.IrQwebFieldImageModel, name, i)
}

func (svc *IrQwebFieldImageService) GetAll() (*types.IrQwebFieldImages, error) {
	i := &types.IrQwebFieldImages{}
	return i, svc.client.getAll(types.IrQwebFieldImageModel, i)
}

func (svc *IrQwebFieldImageService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldImageModel, fields, relations)
}

func (svc *IrQwebFieldImageService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldImageModel, ids, fields, relations)
}

func (svc *IrQwebFieldImageService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldImageModel, ids)
}
