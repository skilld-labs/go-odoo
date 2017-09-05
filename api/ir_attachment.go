package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrAttachmentService struct {
	client *Client
}

func NewIrAttachmentService(c *Client) *IrAttachmentService {
	return &IrAttachmentService{client: c}
}

func (svc *IrAttachmentService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrAttachmentModel, name)
}

func (svc *IrAttachmentService) GetByIds(ids []int) (*types.IrAttachments, error) {
	i := &types.IrAttachments{}
	return i, svc.client.getByIds(types.IrAttachmentModel, ids, i)
}

func (svc *IrAttachmentService) GetByName(name string) (*types.IrAttachments, error) {
	i := &types.IrAttachments{}
	return i, svc.client.getByName(types.IrAttachmentModel, name, i)
}

func (svc *IrAttachmentService) GetByField(field string, value string) (*types.IrAttachments, error) {
	i := &types.IrAttachments{}
	return i, svc.client.getByName(types.IrAttachmentModel, field, value, i)
}

func (svc *IrAttachmentService) GetAll() (*types.IrAttachments, error) {
	i := &types.IrAttachments{}
	return i, svc.client.getAll(types.IrAttachmentModel, i)
}

func (svc *IrAttachmentService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrAttachmentModel, fields, relations)
}

func (svc *IrAttachmentService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrAttachmentModel, ids, fields, relations)
}

func (svc *IrAttachmentService) Delete(ids []int) error {
	return svc.client.delete(types.IrAttachmentModel, ids)
}
