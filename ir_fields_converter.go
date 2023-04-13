package odoo

import (
	"fmt"
)

// IrFieldsConverter represents ir.fields.converter model.
type IrFieldsConverter struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// IrFieldsConverters represents array of ir.fields.converter model.
type IrFieldsConverters []IrFieldsConverter

// IrFieldsConverterModel is the odoo model name.
const IrFieldsConverterModel = "ir.fields.converter"

// Many2One convert IrFieldsConverter to *Many2One.
func (ifc *IrFieldsConverter) Many2One() *Many2One {
	return NewMany2One(ifc.Id.Get(), "")
}

// CreateIrFieldsConverter creates a new ir.fields.converter model and returns its id.
func (c *Client) CreateIrFieldsConverter(ifc *IrFieldsConverter) (int64, error) {
	ids, err := c.Create(IrFieldsConverterModel, []interface{}{ifc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrFieldsConverter creates a new ir.fields.converter model and returns its id.
func (c *Client) CreateIrFieldsConverters(ifcs []*IrFieldsConverter) ([]int64, error) {
	var vv []interface{}
	for _, v := range ifcs {
		vv = append(vv, v)
	}
	return c.Create(IrFieldsConverterModel, vv)
}

// UpdateIrFieldsConverter updates an existing ir.fields.converter record.
func (c *Client) UpdateIrFieldsConverter(ifc *IrFieldsConverter) error {
	return c.UpdateIrFieldsConverters([]int64{ifc.Id.Get()}, ifc)
}

// UpdateIrFieldsConverters updates existing ir.fields.converter records.
// All records (represented by ids) will be updated by ifc values.
func (c *Client) UpdateIrFieldsConverters(ids []int64, ifc *IrFieldsConverter) error {
	return c.Update(IrFieldsConverterModel, ids, ifc)
}

// DeleteIrFieldsConverter deletes an existing ir.fields.converter record.
func (c *Client) DeleteIrFieldsConverter(id int64) error {
	return c.DeleteIrFieldsConverters([]int64{id})
}

// DeleteIrFieldsConverters deletes existing ir.fields.converter records.
func (c *Client) DeleteIrFieldsConverters(ids []int64) error {
	return c.Delete(IrFieldsConverterModel, ids)
}

// GetIrFieldsConverter gets ir.fields.converter existing record.
func (c *Client) GetIrFieldsConverter(id int64) (*IrFieldsConverter, error) {
	ifcs, err := c.GetIrFieldsConverters([]int64{id})
	if err != nil {
		return nil, err
	}
	if ifcs != nil && len(*ifcs) > 0 {
		return &((*ifcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.fields.converter not found", id)
}

// GetIrFieldsConverters gets ir.fields.converter existing records.
func (c *Client) GetIrFieldsConverters(ids []int64) (*IrFieldsConverters, error) {
	ifcs := &IrFieldsConverters{}
	if err := c.Read(IrFieldsConverterModel, ids, nil, ifcs); err != nil {
		return nil, err
	}
	return ifcs, nil
}

// FindIrFieldsConverter finds ir.fields.converter record by querying it with criteria.
func (c *Client) FindIrFieldsConverter(criteria *Criteria) (*IrFieldsConverter, error) {
	ifcs := &IrFieldsConverters{}
	if err := c.SearchRead(IrFieldsConverterModel, criteria, NewOptions().Limit(1), ifcs); err != nil {
		return nil, err
	}
	if ifcs != nil && len(*ifcs) > 0 {
		return &((*ifcs)[0]), nil
	}
	return nil, fmt.Errorf("ir.fields.converter was not found with criteria %v", criteria)
}

// FindIrFieldsConverters finds ir.fields.converter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrFieldsConverters(criteria *Criteria, options *Options) (*IrFieldsConverters, error) {
	ifcs := &IrFieldsConverters{}
	if err := c.SearchRead(IrFieldsConverterModel, criteria, options, ifcs); err != nil {
		return nil, err
	}
	return ifcs, nil
}

// FindIrFieldsConverterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrFieldsConverterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrFieldsConverterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrFieldsConverterId finds record id by querying it with criteria.
func (c *Client) FindIrFieldsConverterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrFieldsConverterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.fields.converter was not found with criteria %v and options %v", criteria, options)
}
