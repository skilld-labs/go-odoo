package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldBarcodeService struct {
	client *Client
}

func NewIrQwebFieldBarcodeService(c *Client) *IrQwebFieldBarcodeService {
	return &IrQwebFieldBarcodeService{client: c}
}

func (svc *IrQwebFieldBarcodeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldBarcodeModel, name)
}

func (svc *IrQwebFieldBarcodeService) GetByIds(ids []int) (*types.IrQwebFieldBarcodes, error) {
	i := &types.IrQwebFieldBarcodes{}
	return i, svc.client.getByIds(types.IrQwebFieldBarcodeModel, ids, i)
}

func (svc *IrQwebFieldBarcodeService) GetByName(name string) (*types.IrQwebFieldBarcodes, error) {
	i := &types.IrQwebFieldBarcodes{}
	return i, svc.client.getByName(types.IrQwebFieldBarcodeModel, name, i)
}

func (svc *IrQwebFieldBarcodeService) GetByField(field string, value string) (*types.IrQwebFieldBarcodes, error) {
	i := &types.IrQwebFieldBarcodes{}
	return i, svc.client.getByName(types.IrQwebFieldBarcodeModel, field, value, i)
}

func (svc *IrQwebFieldBarcodeService) GetAll() (*types.IrQwebFieldBarcodes, error) {
	i := &types.IrQwebFieldBarcodes{}
	return i, svc.client.getAll(types.IrQwebFieldBarcodeModel, i)
}

func (svc *IrQwebFieldBarcodeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldBarcodeModel, fields, relations)
}

func (svc *IrQwebFieldBarcodeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldBarcodeModel, ids, fields, relations)
}

func (svc *IrQwebFieldBarcodeService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldBarcodeModel, ids)
}
