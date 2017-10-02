package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type DecimalPrecisionTestService struct {
	client *Client
}

func NewDecimalPrecisionTestService(c *Client) *DecimalPrecisionTestService {
	return &DecimalPrecisionTestService{client: c}
}

func (svc *DecimalPrecisionTestService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.DecimalPrecisionTestModel, name)
}

func (svc *DecimalPrecisionTestService) GetByIds(ids []int64) (*types.DecimalPrecisionTests, error) {
	d := &types.DecimalPrecisionTests{}
	return d, svc.client.getByIds(types.DecimalPrecisionTestModel, ids, d)
}

func (svc *DecimalPrecisionTestService) GetByName(name string) (*types.DecimalPrecisionTests, error) {
	d := &types.DecimalPrecisionTests{}
	return d, svc.client.getByName(types.DecimalPrecisionTestModel, name, d)
}

func (svc *DecimalPrecisionTestService) GetByField(field string, value string) (*types.DecimalPrecisionTests, error) {
	d := &types.DecimalPrecisionTests{}
	return d, svc.client.getByField(types.DecimalPrecisionTestModel, field, value, d)
}

func (svc *DecimalPrecisionTestService) GetAll() (*types.DecimalPrecisionTests, error) {
	d := &types.DecimalPrecisionTests{}
	return d, svc.client.getAll(types.DecimalPrecisionTestModel, d)
}

func (svc *DecimalPrecisionTestService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.DecimalPrecisionTestModel, fields, relations)
}

func (svc *DecimalPrecisionTestService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.DecimalPrecisionTestModel, ids, fields, relations)
}

func (svc *DecimalPrecisionTestService) Delete(ids []int64) error {
	return svc.client.delete(types.DecimalPrecisionTestModel, ids)
}
